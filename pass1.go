package main

import (
	"math"
)

func Pass1(cr, ci float64) (uint16, uint16) {

	var k uint16
	var t, tt uint16
	t = 0
	tt = 65535

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

			//rimage[y*IMaxX+x] = k

			if k > t {
				t = k
			}

			if k < tt {
				tt = k
			}

		}
	}

	return tt, t
}
