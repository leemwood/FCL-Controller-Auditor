package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/repository"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	manager *repository.Manager
	pkg     *utils.ParsedPackage
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SelectRepoRoot opens a directory dialog to select the repository root
func (a *App) SelectRepoRoot() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select FCL Controller Repository Root",
	})
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", nil
	}

	mgr, err := repository.NewManager(dir)
	if err != nil {
		return "", fmt.Errorf("invalid repository: %v", err)
	}
	a.manager = mgr
	return dir, nil
}

// SelectZip opens a file dialog to select a controller ZIP package
func (a *App) SelectZip() (*utils.ParsedPackage, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Controller ZIP",
		Filters: []runtime.FileFilter{
			{DisplayName: "ZIP Files (*.zip)", Pattern: "*.zip"},
		},
	})
	if err != nil {
		return nil, err
	}
	if file == "" {
		return nil, nil
	}

	pkg, err := utils.ParseControllerZip(file)
	if err != nil {
		return nil, err
	}

	// Check if this is an update
	if a.manager != nil {
		for _, entry := range a.manager.Index {
			if entry.ID == pkg.ControllerID {
				pkg.IsUpdate = true
				copyEntry := entry
				pkg.CurrentIndex = &copyEntry
				break
			}
		}
	}

	a.pkg = pkg
	return pkg, nil
}

// LoadController loads an existing controller from the repository
func (a *App) LoadController(id string) (*utils.ParsedPackage, error) {
	if a.manager == nil {
		return nil, fmt.Errorf("repository not selected")
	}

	version, layout, err := a.manager.LoadControllerDetails(id)
	if err != nil {
		return nil, err
	}

	pkg := &utils.ParsedPackage{
		ControllerID: id,
		VersionInfo:  version,
		IsUpdate:     true,
	}

	if layout != nil {
		lData, _ := json.Marshal(layout)
		var cl models.ControllerLayout
		json.Unmarshal(lData, &cl)
		pkg.Layout = &cl
		pkg.VersionCode = cl.VersionCode
	}

	if version != nil && pkg.VersionCode == 0 {
		pkg.VersionCode = version.Latest.VersionCode
	}

	// Find in index
	for _, entry := range a.manager.Index {
		if entry.ID == id {
			copyEntry := entry
			pkg.IndexEntry = &copyEntry
			pkg.CurrentIndex = &copyEntry
			break
		}
	}

	basePath := filepath.Join(a.manager.RepoRoot, "repo_json", id)
	iconPath := filepath.Join(basePath, "icon.png")
	if _, err := os.Stat(iconPath); err == nil {
		pkg.IconPath = iconPath
	}

	screenshotDir := filepath.Join(basePath, "screenshots")
	if files, err := os.ReadDir(screenshotDir); err == nil {
		for _, f := range files {
			if !f.IsDir() && (strings.HasSuffix(f.Name(), ".png") || strings.HasSuffix(f.Name(), ".jpg")) {
				pkg.Screenshots = append(pkg.Screenshots, filepath.Join(screenshotDir, f.Name()))
			}
		}
	}

	a.pkg = pkg
	return pkg, nil
}

// GetImageBase64 returns the base64 string of an image file
func (a *App) GetImageBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// ApplyUpdate applies the current package update to the repository
func (a *App) ApplyUpdate(selectedCategories []int, author, description, name, intro string) error {
	if a.manager == nil || a.pkg == nil {
		return fmt.Errorf("repo or package not selected")
	}
	if a.pkg.IndexEntry != nil {
		a.pkg.IndexEntry.Categories = selectedCategories
		if name != "" {
			a.pkg.IndexEntry.Name = name
		}
		if intro != "" {
			a.pkg.IndexEntry.Introduction = intro
		}
	}
	if a.pkg.VersionInfo != nil {
		if author != "" {
			a.pkg.VersionInfo.Author = author
		}
		if description != "" {
			a.pkg.VersionInfo.Description = description
		}
	}
	return a.manager.ApplyUpdate(a.pkg)
}

// GetCategories returns the available categories from category.json
func (a *App) GetCategories() []models.Category {
	if a.manager == nil {
		return nil
	}
	return a.manager.Categories
}

// GetRepoIndex returns the current repository index
func (a *App) GetRepoIndex() []models.IndexEntry {
	if a.manager == nil {
		return nil
	}
	return a.manager.Index
}
