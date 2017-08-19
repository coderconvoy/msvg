package msvg

import (
	"bytes"
	"fmt"

	svg "github.com/ajstarks/svgo"
)

const (
	A4W = 2480
	A4H = 3508
)

type CardFunc func(int, int, int, *svg.SVG)

func PageA4(n, nw int, cf CardFunc) *bytes.Buffer {
	return Page(n, nw, A4W, A4H, cf)
}

func Page(n, nw, pw, ph int, cf CardFunc) *bytes.Buffer {
	MW := pw / 20 //Margins
	MH := ph / 20
	nh := (n + nw - 1) / nw
	CW := (pw - 2*MW) / nw //Card widths
	CH := (ph - 2*MH) / nh

	res := &bytes.Buffer{}
	g := svg.New(res)
	g.Start(pw, ph)
	for i := 0; i < n; i++ {
		x := i % nw
		y := i / nw
		g.Gtransform(fmt.Sprintf("translate(%d,%d)", MW+x*CW, MH+y*CH))
		cf(i, CW, CH, g)
		g.Gend()
	}
	g.End()

	return res
}

func ImageHolder(path string, marginX, marginY int) CardFunc {
	return func(n, cw, ch int, g *svg.SVG) {
		g.Image(marginX, marginY, cw-marginX*2, ch-marginY*2, path)
	}
}
