package models

type ControllerLayout struct {
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	Version           string           `json:"version"`
	VersionCode       int              `json:"versionCode"`
	Author            string           `json:"author"`
	Description       string           `json:"description"`
	ControllerVersion int              `json:"controllerVersion"`
	ButtonStyles      []ButtonStyle    `json:"buttonStyles"`
	DirectionStyles   []DirectionStyle `json:"directionStyles"`
	ViewGroups        []ViewGroup      `json:"viewGroups"`
}

type ButtonStyle struct {
	Name                string `json:"name"`
	TextColor           int    `json:"textColor"`
	TextSize            int    `json:"textSize"`
	StrokeColor         int    `json:"strokeColor"`
	StrokeWidth         int    `json:"strokeWidth"`
	CornerRadius        int    `json:"cornerRadius"`
	FillColor           int    `json:"fillColor"`
	TextColorPressed    int    `json:"textColorPressed"`
	TextSizePressed     int    `json:"textSizePressed"`
	StrokeColorPressed  int    `json:"strokeColorPressed"`
	StrokeWidthPressed  int    `json:"strokeWidthPressed"`
	CornerRadiusPressed int    `json:"cornerRadiusPressed"`
	FillColorPressed    int    `json:"fillColorPressed"`
}

type DirectionStyle struct {
	Name        string      `json:"name"`
	StyleType   string      `json:"styleType"`
	ButtonStyle ButtonStyle `json:"buttonStyle"`
	RockerStyle RockerStyle `json:"rockerStyle"`
}

type RockerStyle struct {
	RockerSize         int `json:"rockerSize"`
	BgCornerRadius     int `json:"bgCornerRadius"`
	BgStrokeWidth      int `json:"bgStrokeWidth"`
	BgStrokeColor      int `json:"bgStrokeColor"`
	BgFillColor        int `json:"bgFillColor"`
	RockerCornerRadius int `json:"rockerCornerRadius"`
	RockerStrokeWidth  int `json:"rockerStrokeWidth"`
	RockerStrokeColor  int `json:"rockerStrokeColor"`
	RockerFillColor    int `json:"rockerFillColor"`
}

type ViewGroup struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Visibility string   `json:"visibility"`
	ViewData   ViewData `json:"viewData"`
}

type ViewData struct {
	ButtonList    []Button    `json:"buttonList"`
	DirectionList []Direction `json:"directionList"`
}

type Button struct {
	ID       string   `json:"id"`
	Text     string   `json:"text"`
	Style    string   `json:"style"`
	BaseInfo BaseInfo `json:"baseInfo"`
	Event    Event    `json:"event"`
}

type Direction struct {
	ID       string   `json:"id"`
	Style    string   `json:"style"`
	BaseInfo BaseInfo `json:"baseInfo"`
	// Direction specific fields if any
}

type BaseInfo struct {
	VisibilityType   string     `json:"visibilityType"`
	XPosition        int        `json:"xPosition"`
	YPosition        int        `json:"yPosition"`
	SizeType         string     `json:"sizeType"`
	AbsoluteWidth    int        `json:"absoluteWidth"`
	AbsoluteHeight   int        `json:"absoluteHeight"`
	PercentageWidth  Percentage `json:"percentageWidth"`
	PercentageHeight Percentage `json:"percentageHeight"`
}

type Percentage struct {
	Reference string `json:"reference"`
	Size      int    `json:"size"`
}

type Event struct {
	PointerFollow bool       `json:"pointerFollow"`
	Movable       bool       `json:"Movable"`
	PressEvent    PressEvent `json:"pressEvent"`
}

type PressEvent struct {
	AutoKeep        bool     `json:"autoKeep"`
	AutoClick       bool     `json:"autoClick"`
	OpenMenu        bool     `json:"openMenu"`
	SwitchTouchMode bool     `json:"switchTouchMode"`
	Input           bool     `json:"input"`
	QuickInput      bool     `json:"quickInput"`
	OutputText      string   `json:"outputText"`
	OutputKeycodes  []int    `json:"outputKeycodes"`
	BindViewGroup   []string `json:"bindViewGroup"`
}
