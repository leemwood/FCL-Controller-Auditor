package main

import (
	"context"
	"fmt"
	"os"
	"encoding/base64"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/repository"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/utils"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
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

// GetImageBase64 returns the base64 string of an image file
func (a *App) GetImageBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// ApplyUpdate applies the current package update to the repository
func (a *App) ApplyUpdate() error {
	if a.manager == nil || a.pkg == nil {
		return fmt.Errorf("repo or package not selected")
	}
	return a.manager.ApplyUpdate(a.pkg)
}

// GetRepoIndex returns the current repository index
func (a *App) GetRepoIndex() []models.IndexEntry {
	if a.manager == nil {
		return nil
	}
	return a.manager.Index
}
