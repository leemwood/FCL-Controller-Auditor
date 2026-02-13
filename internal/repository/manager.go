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
	os.RemoveAll(screenshotDest)
	os.MkdirAll(screenshotDest, 0755)
	for _, src := range pkg.Screenshots {
		if err := copyFile(src, filepath.Join(screenshotDest, filepath.Base(src))); err != nil {
			return err
		}
	}

	// Save version.json
	if pkg.VersionInfo != nil {
		versionPath := filepath.Join(destDir, "version.json")
		var finalVersion models.RepoVersion
		
		// Try to load existing version.json
		if existingData, err := os.ReadFile(versionPath); err == nil {
			var existingVersion models.RepoVersion
			if err := json.Unmarshal(existingData, &existingVersion); err == nil {
				finalVersion = existingVersion
				
				// If the new version is different from current latest, move current latest to history
				if pkg.VersionInfo.Latest.VersionCode != existingVersion.Latest.VersionCode {
					// Check if already in history
					existsInHistory := false
					for _, h := range existingVersion.History {
						if h.VersionCode == existingVersion.Latest.VersionCode {
							existsInHistory = true
							break
						}
					}
					if !existsInHistory && existingVersion.Latest.VersionCode != 0 {
						finalVersion.History = append(finalVersion.History, existingVersion.Latest)
					}
				}
			}
		}

		// Update with new info
		finalVersion.Latest = pkg.VersionInfo.Latest
		finalVersion.Author = pkg.VersionInfo.Author
		finalVersion.Description = pkg.VersionInfo.Description
		finalVersion.Screenshot = len(pkg.Screenshots)
		if pkg.VersionInfo.Screenshot > finalVersion.Screenshot {
			finalVersion.Screenshot = pkg.VersionInfo.Screenshot
		}

		// Merge history from package if any (though usually ZIP only has new version)
		for _, h := range pkg.VersionInfo.History {
			exists := false
			for _, eh := range finalVersion.History {
				if eh.VersionCode == h.VersionCode {
					exists = true
					break
				}
			}
			if !exists && h.VersionCode != finalVersion.Latest.VersionCode {
				finalVersion.History = append(finalVersion.History, h)
			}
		}

		vData, _ := json.MarshalIndent(finalVersion, "", "  ")
		os.WriteFile(versionPath, vData, 0644)
	}

	// Save layout
	if pkg.Layout != nil {
		layoutDest := filepath.Join(destDir, "versions")
		os.MkdirAll(layoutDest, 0755)
		lData, _ := json.MarshalIndent(pkg.Layout, "", "  ")
		fileName := fmt.Sprintf("%d.json", pkg.VersionCode)
		if pkg.VersionCode == 0 && pkg.Layout.VersionCode != 0 {
			fileName = fmt.Sprintf("%d.json", pkg.Layout.VersionCode)
		}
		os.WriteFile(filepath.Join(layoutDest, fileName), lData, 0644)
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
