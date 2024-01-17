package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

/* Wrong example given in book. 0xRR is not a number literal:

var palette = []color.Color{color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, color.Black}

*/

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	//rand.Seed(time.Now().UTC().UnixNano()) //rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with a random value.
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	if err := gif.EncodeAll(out, &anim); err != nil {
		panic("encoding process returned as error")
	}

}
