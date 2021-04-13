package utils

import (
	"errors"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type ColorModel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c ColorModel) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

type CoordinateAxes struct {
	XLabel string
	YLabel string
}

type Plotter struct {
	plot   *plot.Plot
	weight int
	height int
}

type Line struct {
	XPoints    []float64
	YPoints    []float64
	ColorModel ColorModel
	LineName   string
}

func CreatePlotter(label string, coordinates CoordinateAxes, w int, h int) *Plotter {
	p := plot.New()
	p.Title.Text = label
	p.X.Label.Text = coordinates.XLabel
	p.Y.Label.Text = coordinates.YLabel

	return &Plotter{
		plot:   p,
		weight: w,
		height: h,
	}
}

func (p *Plotter) DrawGraph(lines []Line, filename string) error {
	for _, line := range lines {
		if len(line.XPoints) != len(line.YPoints) {
			return errors.New("length of point arrays don't match")
		}

		l, err := plotter.NewLine(funcPoints(line.XPoints, line.YPoints))
		if err != nil {
			return err
		}

		l.LineStyle.Color = line.ColorModel
		p.plot.Legend.Add(line.LineName, l)
		p.plot.Add(l)
	}

	err := p.plot.Save(vg.Length(p.weight)*vg.Inch, vg.Length(p.height)*vg.Inch, filename)
	return err
}

func funcPoints(xPoints []float64, yPoints []float64) plotter.XYs {
	pts := make(plotter.XYs, len(xPoints))
	for i := range pts {
		pts[i].X = xPoints[i]
		pts[i].Y = yPoints[i]
	}
	return pts
}
