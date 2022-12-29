package components

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/rumourscape/temporary-mailer/assets"
)

var res = fyne.NewStaticResource("intro.jpg", assets.ImageContent)

func BackgroundImage() *canvas.Image {
	image := canvas.NewImageFromResource(res)
	image.FillMode = canvas.ImageFillStretch
	image.ScaleMode = canvas.ImageScaleFastest

	return image
}
