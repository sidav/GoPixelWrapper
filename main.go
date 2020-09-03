package main

import (
	gpw "GoPixelWrapper/pixel_wrapper"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

func loop() {
	gpw.Init("gpw WRAPPER TEST", 800, 600, 0, 0)
	for !gpw.IsClosed() {
		gpw.Clear()

		r := uint8(rand.Int31n(256))
		g := uint8(rand.Int31n(256))
		b := uint8(rand.Int31n(256))

		gpw.SetColor(r, g, b)

		x := 400.0 // rand.Int31n(800)
		y := 300.0 // rand.Int31n(600)
		rad := float64(rand.Int31n(500)) + 11
		prec := time.Now()
		// gpw.FillPreciseCircle(x, y, rad-10)
		timePrec := time.Since(prec)
		app := time.Now()
		// gpw.DrawApproxCircle(x, y, rad, 1000)
		gpw.FillCircle(x, y, rad)
		timeapp := time.Since(app)
		gpw.PutString(0, 0, fmt.Sprintf("PREC %d, APP %d, diff %d", timePrec, timeapp, timePrec - timeapp))
		fmt.Printf("PREC %d, APP %d, diff %d", timePrec, timeapp, timePrec - timeapp)

		//x1 := int32(100)// rand.Int31n(500)
		//x2 := int32(200) // rand.Int31n(500)
		//x3 := int32(150) // rand.Int31n(500)
		//y1 := int32(100) // rand.Int31n(500)
		//y2 := int32(150) // rand.Int31n(500)
		//y3 := int32(400) // rand.Int31n(500)
		//gpw.FillTriangle(x1, y1, x2, y2, x3, y3)
		//fmt.Printf("%d %d %d %d %d %d \n", x1, y1, x2, y2, x3, y3)

		gpw.Flush()
	}
}

func main() {
	total := 0
	pixelgl.Run(loop)

	fmt.Print(total)

	gpw.PutString(0, 0, "WOLOLO")

	time.Sleep(1000 * time.Millisecond)
}
