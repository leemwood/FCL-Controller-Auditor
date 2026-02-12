package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/utils"
)

type Manager struct {
	RepoRoot   string
	Index      []models.IndexEntry
	Categories []models.Category
}

func NewManager(repoRoot string) (*Manager, error) {
	m := &Manager{RepoRoot: repoRoot}
	if err := m.Load(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Manager) Load() error {
	// Load index.json
	indexPath := filepath.Join(m.RepoRoot, "index.json")
	iData, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(iData, &m.Index); err != nil {
		return err
	}

	// Load category.json
	catPath := filepath.Join(m.RepoRoot, "category.json")
	cData, err := os.ReadFile(catPath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(cData, &m.Categories); err != nil {
		return err
	}

	return nil
}

func (m *Manager) Save() error {
	// Save index.json
	indexPath := filepath.Join(m.RepoRoot, "index.json")
	iData, err := json.MarshalIndent(m.Index, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(indexPath, iData, 0644); err != nil {
		return err
	}

	return nil
}

func (m *Manager) ApplyUpdate(pkg *utils.ParsedPackage) error {
	destDir := filepath.Join(m.RepoRoot, "repo_json", pkg.ControllerID)
	
	// Ensure directory exists
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// Copy Icon
	if pkg.IconPath != "" {
		if err := copyFile(pkg.IconPath, filepath.Join(destDir, "icon.png")); err != nil {
			return err
		}
	}

	// Copy Screenshots
	screenshotDest := filepath.Join(destDir, "screenshots")
	os.MkdirAll(screenshotDest, 0755)
	for _, src := range pkg.Screenshots {
		if err := copyFile(src, filepath.Join(screenshotDest, filepath.Base(src))); err != nil {
			return err
		}
	}

	// Save version.json
	if pkg.VersionInfo != nil {
		vData, _ := json.MarshalIndent(pkg.VersionInfo, "", "  ")
		os.WriteFile(filepath.Join(destDir, "version.json"), vData, 0644)
	}

	// Save layout
	if pkg.Layout != nil {
		layoutDest := filepath.Join(destDir, "versions")
		os.MkdirAll(layoutDest, 0755)
		lData, _ := json.MarshalIndent(pkg.Layout, "", "  ")
		os.WriteFile(filepath.Join(layoutDest, fmt.Sprintf("%d.json", pkg.VersionCode)), lData, 0644)
	}

	// Update index.json in memory
	if pkg.IndexEntry != nil {
		found := false
		for i, entry := range m.Index {
			if entry.ID == pkg.IndexEntry.ID {
				m.Index[i] = *pkg.IndexEntry
				found = true
				break
			}
		}
		if !found {
			m.Index = append(m.Index, *pkg.IndexEntry)
		}
	}

	return m.Save()
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
