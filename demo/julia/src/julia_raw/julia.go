package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"

	"github.com/pkg/profile"
)

const (
	output    = "julia_raw.png"
	size      = 512
	limit     = 200
	colorized = true
)

var flagfill bool
var flagfillgopixel bool
var flagfillgorow bool
var flagPCPU bool
var flagPTrace bool
var flagVerbose bool

func init() {
	flag.BoolVar(&flagfill, "fill", false, "no go routine")
	flag.BoolVar(&flagfillgopixel, "fillgopixel", false, "one go routine per pixel")
	flag.BoolVar(&flagfillgorow, "fillgorow", false, "one go routine per row")
	flag.BoolVar(&flagPCPU, "pcpu", false, "CPU profiling")
	flag.BoolVar(&flagPTrace, "ptrace", false, "Trace profiling")
	flag.BoolVar(&flagVerbose, "verbose", false, "verbose (display duration)")
}
func main() {
	flag.Parse()
	name := ""
	if flagfillgopixel {
		name = "pixel"
	}
	if flagfillgorow {
		name = "row"
	}
	if flagPCPU {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NameProfile(name)).Stop()
	}
	if flagPTrace {
		defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NameProfile(name)).Stop()
	}

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
	start := time.Now()
	if flagfill {
		fillImage(img, c)
	} else if flagfillgopixel {
		fillImagePixel(img, c)
	} else if flagfillgorow {
		fillImageCol(img, c)
	} else {
		panic("flag fill missing")
	}
	t := time.Now()
	elapsed := t.Sub(start)
	if flagVerbose {
		fmt.Println("Duration", elapsed)
	}
	return img
}

func fillImage(img *image.RGBA, c complex128) {

	mapColors := constructColorMap(limit, true)

	for x := float64(0); x < size; x++ {
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

func fillImagePixel(img *image.RGBA, c complex128) {

	mapColors := constructColorMap(limit, true)
	var w sync.WaitGroup
	w.Add(size * size)
	for x := float64(0); x < size; x++ {
		// Our go routine (we have to pass x as a value otherwise its value will change overtime)
		// Check for our column
		for y := float64(0); y < size; y++ {

			go func(i, j float64) {
				_, gap := InJulia(complex(3*i/size-1.5, 3*j/size-1.5), c, limit)
				r, g, b := mapColors(gap)
				// Set the color of our pixel
				img.Set(int(i), int(j), color.RGBA{r, g, b, 255})
				w.Done()
			}(x, y)
		}
	}
	w.Wait()
}

func fillImageCol(img *image.RGBA, c complex128) {

	mapColors := constructColorMap(limit, true)

	var w sync.WaitGroup
	w.Add(size)
	for x := float64(0); x < size; x++ {
		go func(i float64) {
			for y := float64(0); y < size; y++ {
				_, gap := InJulia(complex(3*i/size-1.5, 3*y/size-1.5), c, limit)
				r, g, b := mapColors(gap)
				// Set the color of our pixel
				img.Set(int(i), int(y), color.RGBA{r, g, b, 255})
			}
			w.Done()
		}(x)
	}
	w.Wait()
}
