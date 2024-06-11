package debug

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/nart4hire/fingerprints/lib/types"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{25, 215, 0, 255}
	cyan  = color.RGBA{20, 200, 200, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

// DrawFeatures draws the original image with all the features that are
// detected drawn on top of it. Useful for understanding what data
// we are gathering, and visualise it. Helpful for detecting issues with
// the algorithms or potential next steps.
func DrawFeatures(original image.Image, result *types.DetectionResult) {
	dst := original.(draw.Image)

	for _, minutiae := range result.Minutia {
		switch minutiae.Type {
		case types.Bifurcation:
			drawSquare(dst, minutiae.X, minutiae.Y, red)
		case types.Termination:
			drawSquare(dst, minutiae.X, minutiae.Y, blue)
		case types.Pore:
			drawSquare(dst, minutiae.X, minutiae.Y, green)
		}
	}

	drawFrame(dst, result.Frame.Horizontal, cyan)
	drawFrame(dst, result.Frame.Vertical, cyan)
	drawDiagonalFrame(dst, result.Frame.Diagonal, cyan)

	drawHalfPoint(dst, result.Frame.Diagonal, cyan)
	drawHalfPoint(dst, result.Frame.Horizontal, cyan)
	drawHalfPoint(dst, result.Frame.Vertical, cyan)

}

func drawFrame(dst draw.Image, r image.Rectangle, c color.Color) {
	drawCross(dst, r.Bounds().Min.X, r.Bounds().Min.Y, c)
	drawCross(dst, r.Bounds().Max.X, r.Bounds().Max.Y, c)
}

func drawDiagonalFrame(dst draw.Image, r image.Rectangle, c color.Color) {
	drawEdgeTopLeft(dst, r.Bounds().Min.X, r.Bounds().Min.Y, c)
	drawEdgeBottomRight(dst, r.Bounds().Max.X, r.Bounds().Max.Y, c)
}

func drawHalfPoint(dst draw.Image, r image.Rectangle, c color.Color) {
	halfX, halfY := halfPoint(r)
	drawX(dst, halfX, halfY, c)
}

func drawCircle(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y-1, c)
	dst.Set(x+1, y, c)
	dst.Set(x, y+1, c)
	dst.Set(x-1, y, c)
}

func drawSquare(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y-1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x+1, y-1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x-1, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x-1, y+1, c)
}

func drawCross(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y, c)
	dst.Set(x, y-1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x-1, y, c)
}

func drawX(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x-1, y+1, c)
	dst.Set(x+1, y-1, c)
}

func drawEdgeTopLeft(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x-1, y+1, c)
	dst.Set(x-1, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x, y-1, c)
	dst.Set(x+1, y-1, c)
}

func drawEdgeBottomRight(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x-1, y+1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x+1, y-1, c)
}

func halfPoint(r image.Rectangle) (int, int) {
	return (r.Max.X + r.Min.X) / 2, (r.Max.Y + r.Min.Y) / 2
}
