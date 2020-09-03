package pixel_wrapper

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

var (
	window   *pixelgl.Window
	surface  *imdraw.IMDraw
)

// system funcs

func Init(title string, windowW, windowH, renderW, renderH int32) {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		Resizable:true,
		VSync:  true,
	}
	var err error
	window, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	surface = imdraw.New(nil)
	Clear()
}

func push(x, y float64) {
	surface.Push(pixel.Vec{x, y})
}

func IsClosed() bool {
	return window.Closed()
}

///
///
/////
///////
/////////

// clear/flush funcs

func Flush() {
	surface.Draw(window)
	window.Update()
}

func Clear() {
	window.Clear(color.RGBA{255,255, 255, 255})
}

// draw funcs

func SetColor(r, g, b uint8) {
	surface.Color = color.RGBA{r, g, b, 255}
}

func DrawLine(x, y, x1, y1 float64) {
	surface.Push(pixel.Vec{x, y}, pixel.Vec{x1, y1})
}

//func WaitKey() rune {
//}

func PutString(x, y float64, str string) {
}

func DrawCircle(x0, y0, radius float64) { // midpoint circle algorithm. Calculates each point of the circle.
	push(x0, y0)
	surface.Circle(radius, 1)
}

func FillCircle(x0, y0, radius float64) { // midpoint circle algorithm. Calculates each point of the circle.
	push(x0, y0)
	surface.Circle(radius, 0)
}

func fillTriangle(x0, y0, radius float64) {

}
