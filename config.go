package ipxsgn

import "errors"

type ResizeMode string
type Gravity string
type Extension string

const (
	// Available resize mode
	ResizeAuto ResizeMode = "auto"
	ResizeFit  ResizeMode = "fit"
	ResizeFill ResizeMode = "fill"

	// Available gravity type
	GravityCenter    Gravity = "ce"
	GravitySmart     Gravity = "sm"
	GravityNorth     Gravity = "no"
	GravitySouth     Gravity = "so"
	GravityWest      Gravity = "we"
	GravityEast      Gravity = "ea"
	GravityNorthWest Gravity = "nowe"
	GravityNorthEast Gravity = "noea"
	GravitySouthWest Gravity = "sowe"
	GravitySouthEast Gravity = "soea"

	ExtPNG  Extension = "png"
	ExtJPG  Extension = "jpg"
	ExtWEBP Extension = "webp"
	ExtAVIF Extension = "avif"
	ExtGIF  Extension = "gif"
	ExtICO  Extension = "ico"
	ExtSVG  Extension = "svg"
	ExtHEIC Extension = "heic"
	ExtBMP  Extension = "bmp"
	ExtTIFF Extension = "tiff"
	ExtPDF  Extension = "pdf"
	ExtMP4  Extension = "mp4"
)

type Config struct {
	resize    ResizeMode
	width     uint
	height    uint
	gravity   Gravity
	enlarge   uint8
	extension Extension
}

func NewConfig(resizeMode ResizeMode, gravityType Gravity, outputExtention Extension, w, h uint, bEnlarge bool) (*Config, error) {
	var iEnlarge uint8 = 1
	if !bEnlarge {
		iEnlarge = 0
	}

	if w < 1 || h < 1 {
		return nil, errors.New("config width or height must greater than zero")
	}

	return &Config{
		resize:    resizeMode,
		width:     w,
		height:    h,
		gravity:   gravityType,
		enlarge:   iEnlarge,
		extension: outputExtention,
	}, nil
}

func NewDefaultConfig() (*Config, error) {
	return NewConfig(ResizeAuto, GravitySmart, "", 0, 0, true)
}
