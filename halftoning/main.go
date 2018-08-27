package main

import (
	"flag"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/codemodus/vitals"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func loadImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func toGrayscale(i image.Image) *image.Gray {
	bs := i.Bounds()
	gi := image.NewGray(bs)

	for x := 0; x < bs.Max.X; x++ {
		for y := 0; y < bs.Max.Y; y++ {
			c := i.At(x, y)
			gi.Set(x, y, c)
		}
	}

	return gi
}

func blackWhite(g color.Gray) color.Gray {
	if g.Y < 123 {
		return color.Gray{0}
	}

	return color.Gray{255}
}

func whiteImage(bounds image.Rectangle) *image.Gray {
	i := image.NewGray(bounds)

	for k := range i.Pix {
		i.Pix[k] = 255
	}

	return i
}

func thresholdDither(i *image.Gray) *image.Gray {
	bs := i.Bounds()
	di := image.NewGray(bs)
	ix := bs.Dx()
	iy := bs.Dy()

	for x := 0; x < ix; x++ {
		for y := 0; y < iy; y++ {
			c := blackWhite(i.GrayAt(x, y))
			di.SetGray(x, y, c)
		}
	}

	return di
}

func avgIntensity(i *image.Gray) float64 {
	sum := 0.0
	for _, v := range i.Pix {
		sum += float64(v)
	}

	return sum / float64(len(i.Pix)*256)
}

func randInt(min, max int, rn *rand.Rand) int {
	mmm := max - min
	if mmm <= 0 {
		mmm = 1
	}
	return rn.Intn(mmm) + min
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func gridDither(i *image.Gray, k int, gamma float64) *image.Gray {
	bs := i.Bounds()
	di := whiteImage(bs)
	ix := bs.Dx()
	iy := bs.Dy()

	for x := 0; x < ix; x += 5 {
		if x > 3000 {
			break
		}
		for y := 0; y < iy; y += 5 {
			cell := toGrayscale(i.SubImage(image.Rect(x, y, x+5, y+5))) // 5 = k
			mu := avgIntensity(cell)
			n := math.Pow((1-mu)*8, 2) / 3 // 8 = beta

			if n < 3 { // 3 = alpha
				n = 0
			}

			for j := 0; j < int(n); j++ {
				dix := randInt(x, min(x+5, ix), rng) // 5 = k
				diy := randInt(y, min(y+5, iy), rng) // 5 = k

				di.SetGray(dix, diy, color.Gray{0})
			}
		}
	}

	return di
}

func main() {
	file := ""

	flag.StringVar(&file, "file", file, "file to process")
	flag.Parse()

	if file == "" {
		log.Fatalln("must set file")
	}

	StopCPUProfile, err := vitals.StartCPUProfile("./p.prof")
	if err != nil {
		log.Fatalln(err)
	}
	defer StopCPUProfile()

	in, err := loadImage(file)
	if err != nil {
		log.Fatalln(err)
	}

	gray := toGrayscale(in)
	out := gridDither(gray, 5, 8)

	f, err := os.Create("out.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = f.Close()
	}()

	if err := png.Encode(f, out); err != nil {
		log.Fatalln(err)
	}
}
