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

/* Wrong example given in the book. 0xRR is not a number literal:

var palette = []color.Color{color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, color.Black}

*/

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 255, 0, 255}, color.RGBA{255, 111, 255, 255}}

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
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex uint8 = blackIndex
		if i%3 == 0 {
			colorIndex = uint8(rand.Uint32()%5 + 1)
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors

}
