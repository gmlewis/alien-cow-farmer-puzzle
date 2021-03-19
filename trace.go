// trace converts the p1 and p2 panels into an SVG representation.
// The alien cow farmer puzzle is found on YouTube:
// https://www.youtube.com/watch?v=y6gioQ_1FYU
//
// Units are in mm.
package main

import (
	"log"
	"os"

	svg "github.com/ajstarks/svgo/float"
)

const (
	pinR    = 2
	padding = 2
)

func main() {
	var maxx, maxy int

	for k := range p2 {
		if k.x > maxx {
			maxx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
	}

	c := svg.New(os.Stdout)
	lr := key{maxx, maxy}
	vw, vh := lr.rightX(padding), lr.downY(padding)
	c.Startview(vw, vh, 0, 0, vw, vh)
	c.Gstyle("fill:none;stroke:black;stroke-width:0.1")
	renderPanel(c, p1)
	renderPanel(c, p2)
	c.Gend()
	c.End()

	log.Printf("Done.")
}

func renderPanel(c *svg.SVG, p panel) {
	for k := range p {
		render(c, k, p)
	}

	var initialized bool
	var minx, miny, maxx, maxy int
	for k := range p {
		if !initialized {
			minx, maxx = k.x, k.x
			miny, maxy = k.y, k.y
			initialized = true
			continue
		}
		if k.x < minx {
			minx = k.x
		}
		if k.y < miny {
			miny = k.y
		}
		if k.x > maxx {
			maxx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
	}

	ul := key{minx, miny}
	ur := key{maxx, miny}
	ll := key{minx, maxy}
	lr := key{maxx, maxy}

	c.Arc(ul.leftX(-padding), ul.upY(pinR-padding), pinR, pinR, pinR, false, true, ul.leftX(pinR-padding), ul.upY(-padding))
	c.Line(ul.leftX(pinR-padding), ul.upY(-padding), ur.rightX(padding-pinR), ur.upY(-padding))
	c.Arc(ur.rightX(padding-pinR), ur.upY(-padding), pinR, pinR, pinR, false, true, ur.rightX(padding), ur.upY(pinR-padding))
	c.Line(ur.rightX(padding), ur.upY(pinR-padding), lr.rightX(padding), lr.downY(padding-pinR))
	c.Arc(lr.rightX(padding), lr.downY(padding-pinR), pinR, pinR, pinR, false, true, lr.rightX(padding-pinR), lr.downY(padding))
	c.Line(lr.rightX(padding-pinR), lr.downY(padding), ll.leftX(pinR-padding), ll.downY(padding))
	c.Arc(ll.leftX(pinR-padding), ll.downY(padding), pinR, pinR, pinR, false, true, ll.leftX(-padding), ll.downY(padding-pinR))
	c.Line(ll.leftX(-padding), ll.downY(padding-pinR), ul.leftX(-padding), ul.upY(pinR-padding))
}

func render(c *svg.SVG, k key, p panel) {
	up := p[key{k.x, k.y - 1}]
	down := p[key{k.x, k.y + 1}]
	left := p[key{k.x - 1, k.y}]
	right := p[key{k.x + 1, k.y}]

	// First, render the padded connections.
	if up {
		if left {
			c.Line(k.leftX(0), k.upY(0), k.leftX(0), k.upY(-0.5*padding))
		} else {
			c.Line(k.leftX(0), k.upY(pinR), k.leftX(0), k.upY(-0.5*padding))
		}
		if right {
			c.Line(k.rightX(0), k.upY(0), k.rightX(0), k.upY(-0.5*padding))
		} else {
			c.Line(k.rightX(0), k.upY(pinR), k.rightX(0), k.upY(-0.5*padding))
		}
	}
	if down {
		if left {
			c.Line(k.leftX(0), k.downY(0), k.leftX(0), k.downY(0.5*padding))
		} else {
			c.Line(k.leftX(0), k.downY(-pinR), k.leftX(0), k.downY(0.5*padding))
		}
		if right {
			c.Line(k.rightX(0), k.downY(0), k.rightX(0), k.downY(0.5*padding))
		} else {
			c.Line(k.rightX(0), k.downY(-pinR), k.rightX(0), k.downY(0.5*padding))
		}
	}
	if left {
		if up {
			c.Line(k.leftX(0), k.upY(0), k.leftX(-0.5*padding), k.upY(0))
		} else {
			c.Line(k.leftX(pinR), k.upY(0), k.leftX(-0.5*padding), k.upY(0))
		}
		if down {
			c.Line(k.leftX(0), k.downY(0), k.leftX(-0.5*padding), k.downY(0))
		} else {
			c.Line(k.leftX(pinR), k.downY(0), k.leftX(-0.5*padding), k.downY(0))
		}
	}
	if right {
		if up {
			c.Line(k.rightX(0), k.upY(0), k.rightX(0.5*padding), k.upY(0))
		} else {
			c.Line(k.rightX(-pinR), k.upY(0), k.rightX(0.5*padding), k.upY(0))
		}
		if down {
			c.Line(k.rightX(0), k.downY(0), k.rightX(0.5*padding), k.downY(0))
		} else {
			c.Line(k.rightX(-pinR), k.downY(0), k.rightX(0.5*padding), k.downY(0))
		}
	}

	// Now render the arcs.
	if !up && !left {
		c.Arc(k.leftX(0), k.upY(pinR), pinR, pinR, pinR, false, true, k.leftX(pinR), k.upY(0))
	}
	if !up && !right {
		c.Arc(k.rightX(-pinR), k.upY(0), pinR, pinR, pinR, false, true, k.rightX(0), k.upY(pinR))
	}
	if !down && !left {
		c.Arc(k.leftX(pinR), k.downY(0), pinR, pinR, pinR, false, true, k.leftX(0), k.downY(-pinR))
	}
	if !down && !right {
		c.Arc(k.rightX(0), k.downY(-pinR), pinR, pinR, pinR, false, true, k.rightX(-pinR), k.downY(0))
	}
}

func (k key) leftX(dx float64) float64 {
	return pinR + padding + dx + float64(k.x*(padding+2*pinR)) - pinR
}

func (k key) rightX(dx float64) float64 {
	return pinR + padding + dx + float64(k.x*(padding+2*pinR)) + pinR
}

func (k key) upY(dy float64) float64 {
	return pinR + padding + dy + float64(k.y*(padding+2*pinR)) - pinR
}

func (k key) downY(dy float64) float64 {
	return pinR + padding + dy + float64(k.y*(padding+2*pinR)) + pinR
}

type key struct {
	x int
	y int
}

type panel map[key]bool

var p1 = panel{
	{1, 0}: true,
	{5, 0}: true,
	{0, 1}: true,
	{1, 1}: true,
	{2, 1}: true,
	{4, 1}: true,
	{5, 1}: true,
	{6, 1}: true,
	{0, 2}: true,
	{2, 2}: true,
	{3, 2}: true,
	{4, 2}: true,
	{6, 2}: true,
	{0, 3}: true,
	{3, 3}: true,
	{6, 3}: true,
	{0, 4}: true,
	{6, 4}: true,
	{0, 5}: true,
	{1, 5}: true,
	{5, 5}: true,
	{6, 5}: true,
	{1, 6}: true,
	{5, 6}: true,

	{0, 8}:  true,
	{5, 8}:  true,
	{0, 9}:  true,
	{4, 9}:  true,
	{5, 9}:  true,
	{0, 10}: true,
	{1, 10}: true,
	{3, 10}: true,
	{4, 10}: true,
	{1, 11}: true,
	{2, 11}: true,
	{3, 11}: true,
	{3, 12}: true,
	{3, 13}: true,
	{4, 13}: true,
	{6, 13}: true,
	{4, 14}: true,
	{5, 14}: true,
	{6, 14}: true,
}

var p2 = panel{
	{9, 0}:  true,
	{13, 0}: true,
	{8, 1}:  true,
	{9, 1}:  true,
	{12, 1}: true,
	{13, 1}: true,
	{14, 1}: true,
	{8, 2}:  true,
	{11, 2}: true,
	{12, 2}: true,
	{14, 2}: true,
	{8, 3}:  true,
	{11, 3}: true,
	{14, 3}: true,
	{8, 4}:  true,
	{10, 4}: true,
	{11, 4}: true,
	{14, 4}: true,
	{8, 5}:  true,
	{9, 5}:  true,
	{10, 5}: true,
	{13, 5}: true,
	{14, 5}: true,
	{13, 6}: true,

	{9, 8}:   true,
	{13, 8}:  true,
	{8, 9}:   true,
	{9, 9}:   true,
	{10, 9}:  true,
	{13, 9}:  true,
	{14, 9}:  true,
	{8, 10}:  true,
	{10, 10}: true,
	{11, 10}: true,
	{14, 10}: true,
	{8, 11}:  true,
	{11, 11}: true,
	{14, 11}: true,
	{8, 12}:  true,
	{11, 12}: true,
	{14, 12}: true,
	{8, 13}:  true,
	{9, 13}:  true,
	{11, 13}: true,
	{12, 13}: true,
	{14, 13}: true,
	{9, 14}:  true,
	{12, 14}: true,
	{13, 14}: true,
	{14, 14}: true,
}
