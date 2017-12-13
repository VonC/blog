package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/cmplx"
	"os"

	"github.com/pkg/profile"
)

const (
	output    = "julia_raw.png"
	size      = 512
	limit     = 200
	colorized = true
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	// uncomment these lines to generate traces into stdout.
	// trace.Start(os.Stdout)
	// defer trace.Stop()

	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img := createImage(size, limit, complex(0.312, 0.5))

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func createImage(size float64, limit float64, c complex128) *image.RGBA {
	// Create our image
	img := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))
	// initialize image
	background := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	fillImage(img, c)
	return img
}

func fillImage(img *image.RGBA, c complex128) {

	mapColors := constructColorMap(limit, true)

	for x := float64(0); x < size; x++ {
		// Our go routine (we have to pass x as a value otherwise its value will change overtime)
		// Check for our column
		for y := float64(0); y < size; y++ {
			_, gap := InJulia(complex(3*x/size-1.5, 3*y/size-1.5), c, limit)
			r, g, b := mapColors(gap)
			// Set the color of our pixel
			img.Set(int(x), int(y), color.RGBA{r, g, b, 255})
		}
	}
}

// InJulia checks if complex number c is in Julia z0 set
func InJulia(z0, c complex128, n float64) (bool, float64) {
	z := z0
	for i := float64(0); i < n; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return false, i
		}
	}

	return true, n
}
