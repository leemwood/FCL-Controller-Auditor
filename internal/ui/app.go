package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/repository"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/utils"
)

type AuditorApp struct {
	App        fyne.App
	Window     fyne.Window
	RepoMgr    *repository.Manager
	CurrentPkg *utils.ParsedPackage

	// UI Components
	ControllerList *widget.List
	InfoLabel      *widget.Label
	IconImage      *canvas.Image
	Preview        *ControllerPreview
	ScreenshotCont *fyne.Container
}

func NewAuditorApp(repoRoot string) (*AuditorApp, error) {
	mgr, err := repository.NewManager(repoRoot)
	if err != nil {
		return nil, err
	}

	a := app.New()
	w := a.NewWindow("FCL Controller Auditor")
	w.Resize(fyne.NewSize(1200, 800))

	auditor := &AuditorApp{
		App:     a,
		Window:  w,
		RepoMgr: mgr,
	}

	auditor.setupUI()
	return auditor, nil
}

func (a *AuditorApp) setupUI() {
	// Left side: Controller List
	a.ControllerList = widget.NewList(
		func() int { return len(a.RepoMgr.Index) },
		func() fyne.CanvasObject { return widget.NewLabel("Template") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(a.RepoMgr.Index[id].Name)
		},
	)
	a.ControllerList.OnSelected = func(id widget.ListItemID) {
		a.loadController(a.RepoMgr.Index[id].ID)
	}

	// Right side: Details & Preview
	a.InfoLabel = widget.NewLabel("Select a controller to view details")
	a.IconImage = canvas.NewImageFromResource(nil)
	a.IconImage.FillMode = canvas.ImageFillContain
	a.IconImage.SetMinSize(fyne.NewSize(64, 64))

	a.Preview = NewControllerPreview(nil)
	a.ScreenshotCont = container.NewHBox()

	// Toolbar
	loadZipBtn := widget.NewButton("Load ZIP Package", a.showZipPicker)
	applyBtn := widget.NewButton("Apply Update", a.applyUpdate)
	toolbar := container.NewHBox(loadZipBtn, applyBtn)

	details := container.NewVBox(
		widget.NewLabelWithStyle("Controller Info", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewHBox(a.IconImage, a.InfoLabel),
		widget.NewLabelWithStyle("Preview", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.New(layout.NewMaxLayout(), a.Preview),
		widget.NewLabelWithStyle("Screenshots", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewHScroll(a.ScreenshotCont),
	)

	split := container.NewHSplit(
		container.NewVBox(widget.NewLabel("Controllers"), a.ControllerList),
		container.NewBorder(toolbar, nil, nil, nil, details),
	)
	split.Offset = 0.2

	a.Window.SetContent(split)
}

func (a *AuditorApp) loadController(id string) {
	// Load from repository
	controllerDir := filepath.Join(a.RepoMgr.RepoRoot, "repo_json", id)
	versionPath := filepath.Join(controllerDir, "version.json")
	
	// Implementation for loading existing controller for comparison...
	// For now, just clear the current package
	a.CurrentPkg = nil
	a.InfoLabel.SetText(fmt.Sprintf("Loading ID: %s", id))
}

func (a *AuditorApp) showZipPicker() {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		defer reader.Close()

		pkg, err := utils.ParseControllerZip(reader.URI().Path())
		if err != nil {
			dialog.ShowError(err, a.Window)
			return
		}

		a.displayPackage(pkg)
	}, a.Window)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".zip"}))
	fd.Show()
}

func (a *AuditorApp) displayPackage(pkg *utils.ParsedPackage) {
	a.CurrentPkg = pkg
	a.InfoLabel.SetText(fmt.Sprintf(
		"ID: %s\nName: %s\nAuthor: %s\nVersion: %s (%d)\nDescription: %s",
		pkg.ControllerID, pkg.IndexEntry.Name, pkg.Layout.Author,
		pkg.Layout.Version, pkg.VersionCode, pkg.Layout.Description,
	))

	if pkg.Layout != nil {
		a.Preview.SetLayout(pkg.Layout)
	}

	if pkg.IconPath != "" {
		a.IconImage.File = pkg.IconPath
		a.IconImage.Refresh()
	}

	// Update screenshots
	a.ScreenshotCont.Objects = []fyne.CanvasObject{}
	for _, s := range pkg.Screenshots {
		img := canvas.NewImageFromFile(s)
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(200, 112)) // 16:9 ratio
		a.ScreenshotCont.Add(img)
	}
	a.ScreenshotCont.Refresh()
}

func (a *AuditorApp) applyUpdate() {
	if a.CurrentPkg == nil {
		dialog.ShowInformation("No Package", "Please load a ZIP package first", a.Window)
		return
	}

	err := a.RepoMgr.ApplyUpdate(a.CurrentPkg)
	if err != nil {
		dialog.ShowError(err, a.Window)
		return
	}

	dialog.ShowInformation("Success", "Controller update applied successfully", a.Window)
	a.ControllerList.Refresh()
}

func (a *AuditorApp) Run() {
	a.Window.ShowAndRun()
}
