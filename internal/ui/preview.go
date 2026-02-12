package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tungsten-fcl/fcl-controller-auditor/internal/models"
)

type ControllerPreview struct {
	widget.BaseWidget
	Layout *models.ControllerLayout
}

func NewControllerPreview(layout *models.ControllerLayout) *ControllerPreview {
	p := &ControllerPreview{Layout: layout}
	p.ExtendBaseWidget(p)
	return p
}

func (p *ControllerPreview) CreateRenderer() fyne.WidgetRenderer {
	return &controllerPreviewRenderer{
		preview: p,
		objects: []fyne.CanvasObject{},
	}
}

type controllerPreviewRenderer struct {
	preview *ControllerPreview
	objects []fyne.CanvasObject
}

func (r *controllerPreviewRenderer) Destroy() {}

func (r *controllerPreviewRenderer) Layout(size fyne.Size) {
	r.objects = []fyne.CanvasObject{}
	if r.preview.Layout == nil {
		return
	}

	// Calculate scale to fit in the given size while maintaining a 16:9 aspect ratio (common for mobile)
	// Or just use the full available size.
	screenWidth := size.Width
	screenHeight := size.Height

	// Map styles for quick lookup
	styles := make(map[string]models.ButtonStyle)
	for _, s := range r.preview.Layout.ButtonStyles {
		styles[s.Name] = s
	}

	for _, group := range r.preview.Layout.ViewGroups {
		if group.Visibility != "VISIBLE" {
			continue
		}

		for _, btn := range group.ViewData.ButtonList {
			style, ok := styles[btn.Style]
			if !ok && len(r.preview.Layout.ButtonStyles) > 0 {
				style = r.preview.Layout.ButtonStyles[0]
			}

			w, h := calculateSize(btn.BaseInfo, screenWidth, screenHeight)
			x, y := float32(btn.BaseInfo.XPosition)*screenWidth/1000, float32(btn.BaseInfo.YPosition)*screenHeight/1000

			// Adjust for center positioning if needed? FCL usually uses top-left for (x,y)
			// but size might be centered. Let's assume (x,y) is top-left for now.

			rect := canvas.NewRectangle(intToColor(style.FillColor))
			rect.StrokeColor = intToColor(style.StrokeColor)
			rect.StrokeWidth = float32(style.StrokeWidth) / 10
			rect.Resize(fyne.NewSize(w, h))
			rect.Move(fyne.NewPos(x, y))

			text := canvas.NewText(btn.Text, intToColor(style.TextColor))
			text.Alignment = fyne.TextAlignCenter
			text.TextSize = float32(style.TextSize)
			text.Resize(fyne.NewSize(w, h))
			text.Move(fyne.NewPos(x, y))

			r.objects = append(r.objects, rect, text)
		}
	}
}

func (r *controllerPreviewRenderer) MinSize() fyne.Size {
	return fyne.NewSize(320, 180) // 16:9 minimum
}

func (r *controllerPreviewRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *controllerPreviewRenderer) Refresh() {
	r.Layout(r.preview.Size())
	canvas.Refresh(r.preview)
}

func calculateSize(info models.BaseInfo, sw, sh float32) (float32, float32) {
	if info.SizeType == "ABSOLUTE" {
		return float32(info.AbsoluteWidth), float32(info.AbsoluteHeight)
	}

	var w, h float32
	if info.PercentageWidth.Reference == "SCREEN_WIDTH" {
		w = float32(info.PercentageWidth.Size) * sw / 1000
	} else {
		w = float32(info.PercentageWidth.Size) * sh / 1000
	}

	if info.PercentageHeight.Reference == "SCREEN_HEIGHT" {
		h = float32(info.PercentageHeight.Size) * sh / 1000
	} else {
		h = float32(info.PercentageHeight.Size) * sw / 1000
	}

	return w, h
}

func intToColor(val int) color.Color {
	// FCL uses ARGB int32
	a := uint8((val >> 24) & 0xff)
	r := uint8((val >> 16) & 0xff)
	g := uint8((val >> 8) & 0xff)
	b := uint8(val & 0xff)
	return color.RGBA{r, g, b, a}
}

func (p *ControllerPreview) SetLayout(l *models.ControllerLayout) {
	p.Layout = l
	p.Refresh()
}
