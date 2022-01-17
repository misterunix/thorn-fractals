package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

// NewValue = (((OldValue - OldMin) * (NewMax - NewMin)) / (OldMax - OldMin)) + NewMin
//var rimage []uint16

var IMaxX int
var IMaxY int

var xmin float64
var xmax float64
var ymin float64
var ymax float64

var iterations uint16
var escape float64

var imagecount float64

func main() {

	IMaxX = 1000
	IMaxY = 1000
	xmin = -math.Pi
	xmax = math.Pi
	ymin = -math.Pi
	ymax = math.Pi

	p1 := image.Point{
		X: 0,
		Y: 0,
	}
	p2 := image.Point{
		X: IMaxX,
		Y: IMaxY,
	}
	rec := image.Rectangle{Min: p1, Max: p2}
	Img := image.NewRGBA(rec)
	//color1 := color.RGBA{0xff, 0xff, 0xff, 0xff}
	bgcolor := image.NewUniform(color.Black)

	//draw.Draw(Img, Img.Bounds(), bgcolor, image.Point{}, draw.Src)

	var k uint16
	//var avraw uint64

	iterations = 65535
	escape = 10000

	rseed := time.Now().UnixNano()
	randomSource := rand.NewSource(rseed)
	rnd := rand.New(randomSource)

	//rimage = make([]uint16, IMaxX*IMaxY)
	buckets := make(map[int]int)

	/*
	   cr = (rand() % 10000) / 500.0 - 10; // -10 -> 10;
	   ci = (rand() % 10000) / 500.0 - 10;
	*/

	//cr := float64(rnd.Intn(IMaxX))/(float64(IMaxX)/2.0) - 10.0 // range +- 10
	//ci := float64(rnd.Intn(IMaxY))/(float64(IMaxY)/2.0) - 10.0
	cr := float64(rnd.Intn(IMaxX))/(float64(IMaxX)/2.0) - 10.0 // range +- 10
	ci := float64(rnd.Intn(IMaxY))/(float64(IMaxY)/2.0) - 10.0
	t, tt := Pass1(cr, ci)
	cro := cr
	fmt.Println(t, tt)
	fmt.Printf("cr:%f ci:%f\n", cr, ci)
	var fnci int
	fnci = 0
	for imagecount = 0; imagecount < 360; imagecount += 0.1 {
		ra := (imagecount / 360.0) * math.Pi * 2
		cr = cro + math.Sin(ra)
		//ci = ci + math.Cos(ra)
		draw.Draw(Img, Img.Bounds(), bgcolor, image.Point{}, draw.Src)
		for x := 0; x < IMaxX; x++ {
			zr := xmin + float64(x)*(xmax-xmin)/float64(IMaxX)

			for y := 0; y < IMaxY; y++ {
				zi := ymin + float64(y)*(ymax-ymin)/float64(IMaxY)

				ir := zr
				ii := zi

				for k = 0; k < iterations; k++ {
					a := ir
					b := ii
					ir = a/math.Cos(b) + cr
					ii = b/math.Sin(a) + ci
					if ir*ir+ii*ii > escape {
						break
					}
				}
				// density[j*NX+i] = k;
				//fmt.Printf("%04X\n", k)
				if k == 0xFFFF {
					k = 0
				}

				buckets[int(k)]++

				//rimage[y*IMaxX+x] = k

				//tv := (((k - t) * (255 - 1)) / (tt - t)) + 1

				tv := (((k - t) * (255 - 64)) / (tt - t)) + 64

				//fmt.Printf("%d = (((%d - %d) * (255 - 1)) / (%d - %d)) + 1\n", tv, k, t, tt, t)

				//fmt.Println(tv)
				r := uint8(tv)
				g := r
				b := r
				al := uint8(255)
				cc := color.RGBA{r, g, b, al}
				//fmt.Printf("%+v\n", cc)

				Img.Set(x, y, cc)

				//avraw += uint64(k)
				/*
					if k > t {
						t = k
					}

					if k < tt {
						tt = k
					}
				*/
			}
		}
		fnci++
		sfn := fmt.Sprintf("images/%07d.png", fnci)
		f, err := os.Create(sfn)
		if err != nil {
			// Handle error
			log.Fatalln(err)
		}
		err = png.Encode(f, Img)
		if err != nil {
			log.Fatalln(err)
			// Handle error
		}
		f.Close()

	}

	fmt.Println(tt, t)
	//for e1, e2 := range buckets {
	//		fmt.Printf("%d,%d\n", e1, e2)
	//}

	/*
		fmt.Println(tt, t)
		file, err := os.Create("test.pgm")
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		av := float64(avraw) / (float64(IMaxX) * float64(IMaxY))
		fmt.Printf("Average value: %f\n", av)

		header := fmt.Sprintf("P5 %d %d 65535\n", IMaxX, IMaxY)

		_, err = file.Write([]byte(header))
		if err != nil {
			log.Fatal(err)
		}

		var binbuf bytes.Buffer
		binary.Write(&binbuf, binary.LittleEndian, image)
		_, err = file.Write(binbuf.Bytes())
		if err != nil {
			log.Fatal(err)
		}

	*/
}
