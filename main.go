package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tungsten-fcl/fcl-controller-auditor/internal/ui"
)

//go:embed lib/x64/opengl32.dll lib/x64/libgallium_wgl.dll lib/x64/dxil.dll
var mesaFiles embed.FS

func init() {
	// Extract Mesa DLLs for RDP/Software rendering support on Windows
	if runtime.GOOS == "windows" {
		exePath, err := os.Executable()
		if err != nil {
			return
		}
		exeDir := filepath.Dir(exePath)
		exeName := filepath.Base(exePath)

		// 1. Create .local file to force local DLL loading
		localFile := filepath.Join(exeDir, exeName+".local")
		if _, err := os.Stat(localFile); os.IsNotExist(err) {
			_ = os.WriteFile(localFile, []byte{}, 0644)
		}

		// 2. Extract necessary DLLs
		files := []string{"opengl32.dll", "libgallium_wgl.dll", "dxil.dll"}
		for _, f := range files {
			targetPath := filepath.Join(exeDir, f)
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				fmt.Printf("Extracting %s for software rendering support...\n", f)
				data, err := mesaFiles.ReadFile("lib/x64/" + f)
				if err != nil {
					fmt.Printf("Failed to read embedded %s: %v\n", f, err)
					continue
				}
				err = os.WriteFile(targetPath, data, 0644)
				if err != nil {
					fmt.Printf("Failed to write %s: %v\n", f, err)
				}
			}
		}

		// 3. Set environment variables to hint Mesa
		os.Setenv("GALLIUM_DRIVER", "llvmpipe")
	}
}

func main() {
	// Assume we are running inside the auditor directory, so parent is the repo root
	// Or use an environment variable/flag.
	repoRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
	// If we are in FCL-Controller-Auditor, go up one level
	if filepath.Base(repoRoot) == "FCL-Controller-Auditor" {
		repoRoot = filepath.Dir(repoRoot)
	}

	app, err := ui.NewAuditorApp(repoRoot)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	app.Run()
}
