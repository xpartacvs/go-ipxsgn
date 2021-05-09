package ipxsgn

const (
	ResizeAuto = "auto"
	ResizeFit  = "fit"
	ResizeFill = "fill"

	GravityCenter    = "ce"
	GravitySmart     = "sm"
	GravityNorth     = "no"
	GravitySouth     = "so"
	GravityWest      = "we"
	GravityEast      = "ea"
	GravityNorthWest = "nowe"
	GravityNorthEast = "noea"
	GravitySouthWest = "sowe"
	GravitySouthEast = "soea"

	ExtPNG  = "png"
	ExtJPG  = "jpg"
	ExtWEBP = "webp"
	ExtAVIF = "avif"
	ExtGIF  = "gif"
	ExtICO  = "ico"
	ExtSVG  = "svg"
	ExtHEIC = "heic"
	ExtBMP  = "bmp"
	ExtTIFF = "tiff"
	ExtPDF  = "pdf"
	ExtMP4  = "mp4"
)

type Config struct {
	Resize    string `validate:"oneof=auto fit fill"`
	Width     uint   `validate:"gte=0"`
	Height    uint   `validate:"gte=0"`
	Gravity   string `validate:"oneof=ce sm no so we ea nowe noea sowe soea"`
	Enlarge   uint8  `validate:"gte=0,max=1"`
	Extension string `validate:"omitempty,oneof=png jpg webp avif gif ico svg heic bmp tiff pdf mp4"`
}

func NewConfig(strResize, strGravity, strExtension string, iWidth, iHeight uint, bEnlarge bool) *Config {
	var iEnlarge uint8 = 1
	if !bEnlarge {
		iEnlarge = 0
	}
	return &Config{
		Resize:    strResize,
		Width:     iWidth,
		Height:    iHeight,
		Gravity:   strGravity,
		Enlarge:   iEnlarge,
		Extension: strExtension,
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Resize:    ResizeAuto,
		Width:     0,
		Height:    0,
		Gravity:   GravitySmart,
		Enlarge:   1,
		Extension: "",
	}
}
