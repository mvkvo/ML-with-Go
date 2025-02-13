package graphs

import (
	"fmt"
	"image"
	"image/color"
	"net/http"

	"github.com/labstack/echo/v4"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func CreateGraphFromData(title string, label_x string, label_y string, graph_type string, points plotter.XYs) *plot.Plot {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = label_x
	p.Y.Label.Text = label_y

	switch graph_type {
	case "linear":
		line, err := plotter.NewLine(points)
		if err != nil {
			panic(err)
		}
		line.Color = plotter.DefaultLineStyle.Color
		line.LineStyle = draw.LineStyle{Width: vg.Points(2)}
		p.Add(line)
	case "scatter":
		scatter, err := plotter.NewScatter(points)
		if err != nil {
			panic(err)
		}
		scatter.GlyphStyle.Color = color.RGBA{R: 255, G: 255, B: 0}
		p.Add(scatter)
	}
	return p
}

func ToPlotter(x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, len(x))
	for i := range x {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

func SaveGraph(p *plot.Plot) {
	name := fmt.Sprint(p.Title.Text, ".png")
	if err := p.Save(10*vg.Inch, 7*vg.Inch, fmt.Sprint("./static/imgs/", name)); err != nil {
	}
	fmt.Printf("\n\nPlot has been saved as %s", name)
}

func ImageFromPlot(p *plot.Plot, width, height int) (*image.RGBA, error) {
	// Tworzymy płótno dla wykresu
	img := vgimg.New(vg.Points(float64(width)), vg.Points(float64(height)))
	dc := draw.New(img)

	// Rysujemy wykres na płótnie
	p.Draw(dc)

	// Pobieramy końcowy obraz i konwertujemy na *image.RGBA
	rgbaImg, ok := img.Image().(*image.RGBA)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to convert image")
	}
	return rgbaImg, nil
}
