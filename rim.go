// rim.go
// read as "r img o" as in "random image output", "rimgo"
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func Gen(n int) uint8 {
	return uint8(rand.Intn(n))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	img := image.NewRGBA(image.Rect(0,0,200,200))
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			img.Set(x, y, color.RGBA{Gen(0xff),Gen(0xff),Gen(0xff),Gen(0xff)})
		}
	}

	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}