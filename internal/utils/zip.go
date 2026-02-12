package utils

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
)

type ParsedPackage struct {
	ControllerID string
	VersionCode  int
	Layout       *models.ControllerLayout
	VersionInfo  *models.RepoVersion
	IndexEntry   *models.IndexEntry
	IconPath     string
	Screenshots  []string
	TempDir      string
}

func ParseControllerZip(zipPath string) (*ParsedPackage, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	tempDir, err := os.MkdirTemp("", "fcl-auditor-*")
	if err != nil {
		return nil, err
	}

	pkg := &ParsedPackage{
		TempDir: tempDir,
	}

	var controllerID string

	// First pass: find controller ID and extract files
	for _, f := range r.File {
		parts := strings.Split(filepath.ToSlash(f.Name), "/")
		if len(parts) < 2 {
			continue
		}
		if controllerID == "" {
			controllerID = parts[0]
			pkg.ControllerID = controllerID
		} else if parts[0] != controllerID {
			return nil, fmt.Errorf("multiple controller IDs found in zip: %s and %s", controllerID, parts[0])
		}

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		destPath := filepath.Join(tempDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(destPath, 0755)
		} else {
			os.MkdirAll(filepath.Dir(destPath), 0755)
			destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				rc.Close()
				return nil, err
			}
			_, err = io.Copy(destFile, rc)
			destFile.Close()
			if err != nil {
				rc.Close()
				return nil, err
			}
		}
		rc.Close()
	}

	if controllerID == "" {
		return nil, fmt.Errorf("could not find controller ID in zip")
	}

	// Second pass: load data
	basePath := filepath.Join(tempDir, controllerID)

	// Load version.json
	vData, err := os.ReadFile(filepath.Join(basePath, "version.json"))
	if err == nil {
		var rv models.RepoVersion
		if err := json.Unmarshal(vData, &rv); err == nil {
			pkg.VersionInfo = &rv
			pkg.VersionCode = rv.Latest.VersionCode
		}
	}

	// Load index.json
	iData, err := os.ReadFile(filepath.Join(basePath, "index.json"))
	if err == nil {
		var ie models.IndexEntry
		if err := json.Unmarshal(iData, &ie); err == nil {
			pkg.IndexEntry = &ie
		}
	}

	// Load layout (controller.json in versions folder)
	if pkg.VersionCode != 0 {
		layoutPath := filepath.Join(basePath, "versions", fmt.Sprintf("%d.json", pkg.VersionCode))
		lData, err := os.ReadFile(layoutPath)
		if err == nil {
			var cl models.ControllerLayout
			if err := json.Unmarshal(lData, &cl); err == nil {
				pkg.Layout = &cl
			}
		}
	}

	// Icon
	iconPath := filepath.Join(basePath, "icon.png")
	if _, err := os.Stat(iconPath); err == nil {
		pkg.IconPath = iconPath
	}

	// Screenshots
	screenshotDir := filepath.Join(basePath, "screenshots")
	if entries, err := os.ReadDir(screenshotDir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".png") {
				pkg.Screenshots = append(pkg.Screenshots, filepath.Join(screenshotDir, entry.Name()))
			}
		}
	}

	return pkg, nil
}

func (p *ParsedPackage) Cleanup() {
	os.RemoveAll(p.TempDir)
}
