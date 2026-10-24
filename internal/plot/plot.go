package plot

import (
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func init() {
	plot.DefaultFont = font.Font{
		Typeface: "Liberation",
		Variant:  "Sans",
	}
}

func savePNG(p *plot.Plot, filename string) error {
	width := 5 * vg.Inch
	height := 5 * vg.Inch

	canvas := vgimg.NewWith(
		vgimg.UseWH(width, height),
		vgimg.UseDPI(300),
	)

	dc := draw.New(canvas)
	p.Draw(draw.Crop(
		dc,
		vg.Points(5),  // left
		-vg.Points(5), // right
		vg.Points(5),  // bottom
		-vg.Points(5), // top
	))

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = (vgimg.PngCanvas{Canvas: canvas}).WriteTo(f)
	return err
}

func Scatter(fname string, points [][]float64, clusters []int, title string) error {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "PC0"
	p.Y.Label.Text = "PC1"

	clusterMap := make(map[int]plotter.XYs)

	for i, point := range points {
		cluster := clusters[i]
		clusterMap[cluster] = append(clusterMap[cluster], plotter.XY{X: point[0], Y: point[1]})
	}

	for cluster, xys := range clusterMap {
		scatter, err := plotter.NewScatter(xys)
		if err != nil {
			panic(err)
		}
		scatter.GlyphStyle = draw.GlyphStyle{
			Color:  plotutil.Color(cluster),
			Radius: vg.Points(2),
			Shape:  draw.CircleGlyph{},
		}

		p.Add(scatter)
		p.Legend.Add("Cluster "+strconv.Itoa(cluster), scatter)
	}

	p.X.Padding = 0
	p.Y.Padding = 0
	if err := savePNG(p, fname); err != nil {
		return err
	}
	return nil
}
