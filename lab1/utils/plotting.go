package utils

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Plotter struct {
	plot *plot.Plot
}

func CreatePlotter() *Plotter {
	return new(Plotter)
}

func (p *Plotter) DrawGraph(xPoints []float64, yPoints []float64, k int, filename string,
	r uint8, g uint8, b uint8) {
	p.plot = plot.New()
	p.plot.Title.Text = "Зависимость суммарной ошибки от количества эпох"
	p.plot.X.Label.Text = "Количество эпох"
	p.plot.Y.Label.Text = "Суммарная ошибка"

	l, err := plotter.NewLine(funcPoints(xPoints, yPoints, k))
	if err != nil {
		fmt.Println(err)
		return
	}
	l.LineStyle.Color = color.RGBA{R: r, G: g, B: b, A: 255}
	p.plot.Add(l)
	if err := p.plot.Save(10 * vg.Inch, 8 * vg.Inch, filename); err != nil {
		fmt.Println(err)
		return
	}
}

func funcPoints(xPoints []float64, yPoints []float64, k int) plotter.XYs {
	pts := make(plotter.XYs, k)
	for i := range pts {
		pts[i].X = xPoints[i]
		pts[i].Y = yPoints[i]
	}
	return pts
}
