package squarify

import (
	"image"
	"image/color"
	"image/draw"
)

// Image returns a squared version of the source image without cropping.
// The added space is filled with a background image that must be at least max * max in size,
// where max is Max(width, height).
// If the background image is nil, white will be used to fill in the space.
func Image(src image.Image, background image.Image) image.Image {
	var (
		rect image.Rectangle
		m    = src.Bounds().Max
		max  = m.X
	)
	if max < m.Y {
		// portrait
		max = m.Y
		rect = image.Rect((max-m.X)/2, 0, (max+m.X)/2, max)
	} else {
		// landscape
		rect = image.Rect(0, (max-m.Y)/2, max, (max+m.Y)/2)
	}
	dst := image.NewRGBA(image.Rect(0, 0, max, max))
	if background == nil {
		background = &image.Uniform{color.White}
	}
	draw.Draw(dst, dst.Bounds(), background, image.ZP, draw.Src)
	draw.Draw(dst, rect, src, image.ZP, draw.Src)
	return dst
}
