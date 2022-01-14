package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var image []uint16

var IMaxX int
var IMaxY int

var xmin float64
var xmax float64
var ymin float64
var ymax float64

func main() {

	IMaxX = 10000
	IMaxY = 10000
	xmin = -math.Pi
	xmax = math.Pi
	ymin = -math.Pi
	ymax = math.Pi

	var escape float64
	var iterations uint16
	var k uint16

	iterations = 65535
	escape = 10000

	rseed := time.Now().UnixNano()
	randomSource := rand.NewSource(rseed)
	rnd := rand.New(randomSource)

	image = make([]uint16, IMaxX*IMaxY)

	/*
	   cr = (rand() % 10000) / 500.0 - 10; // -10 -> 10;
	   ci = (rand() % 10000) / 500.0 - 10;
	*/

	cr := rnd.Intn(10000)/500 - 10 // range +- 10
	ci := rnd.Intn(10000)/500 - 10
	fmt.Printf("cr:%d ci:%d\n", cr, ci)

	t := 0
	tt := 10000000

	for x := 0; x < IMaxX; x++ {
		zr := xmin + float64(x)*(xmax-xmin)/float64(IMaxX)

		for y := 0; y < IMaxY; y++ {
			zi := ymin + float64(y)*(ymax-ymin)/float64(IMaxY)

			ir := zr
			ii := zi

			for k = 0; k < iterations; k++ {
				a := ir
				b := ii
				ir = a/math.Cos(b) + float64(cr)
				ii = b/math.Sin(a) + float64(ci)
				if ir*ir+ii*ii > escape {
					break
				}
			}
			// density[j*NX+i] = k;
			//fmt.Printf("%04X\n", k)
			image[y*IMaxX+x] = k
			if int(k) > t {
				t = int(k)
			}

			if int(k) < tt {
				tt = int(k)
			}

		}
	}
	fmt.Println(tt, t)
	file, err := os.Create("test.pgm")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

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
}
