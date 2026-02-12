package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tungsten-fcl/fcl-controller-auditor/internal/ui"
)

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
