/* Go package for Plotting. */
package plot

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"sort"
)

// Plot is a struct that holds the data to be plotted.
type Plot struct {
	// The data to be plotted.
	Data []float64

	// The width of the plot.
	Width int

	// The height of the plot.
	Height int

	// The color of the plot.
	Color color.Color

	// The name of the plot.
	Name string

	// The name of the file to save the plot to.
	FileName string

	// The minimum value of the plot.
	Min float64

	// The maximum value of the plot.
	Max float64

	// The number of bins to use for the plot.
	Bins int

	// The number of bins to use for the plot.
	BinWidth float64

	// The number of bins to use for the plot.
	BinHeight float64

	// The number of bins to use for the plot.
	BinCounts []int

	// The number of bins to use for the plot.
	BinCenters []float64
}

// NewPlot returns a new plot.
func NewPlot(data []float64, width int, height int, color color.Color, name string, fileName string) *Plot {
	return &Plot{
		Data:      data,
		Width:     width,
		Height:    height,
		Color:     color,
		Name:      name,
		FileName:  fileName,
		Min:       math.Inf(1),
		Max:       math.Inf(-1),
		Bins:      100,
		BinWidth:  float64(width) / 100,
		BinHeight: float64(height) / 100,
	}
}

// SetBins sets the number of bins to use for the plot.
func (p *Plot) SetBins(bins int) {
	p.Bins = bins
	p.BinWidth = float64(p.Width) / float64(p.Bins)
	p.BinHeight = float64(p.Height) / float64(p.Bins)
}

// SetMin sets the minimum value of the plot.
func (p *Plot) SetMin(min float64) {
	p.Min = min
}

// SetMax sets the maximum value of the plot.
func (p *Plot) SetMax(max float64) {
	p.Max = max
}

// SetMinMax sets the minimum and maximum values of the plot.
func (p *Plot) SetMinMax(min float64, max float64) {
	p.Min = min
	p.Max = max
}

// GetMinMax returns the minimum and maximum values of the plot.
func (p *Plot) GetMinMax() (float64, float64) {
	return p.Min, p.Max
}

// GetBinCenters returns the centers of the bins.
func (p *Plot) GetBinCenters() []float64 {
	return p.BinCenters
}

// GetBinCounts returns the counts of the bins.
func (p *Plot) GetBinCounts() []int {
	return p.BinCounts
}

// GetBinWidth returns the width of the bins.
func (p *Plot) GetBinWidth() float64 {
	return p.BinWidth
}

// GetBinHeight returns the height of the bins.
func (p *Plot) GetBinHeight() float64 {
	return p.BinHeight
}

// GetBins returns the number of bins.
func (p *Plot) GetBins() int {
	return p.Bins
}

// GetData returns the data.
func (p *Plot) GetData() []float64 {
	return p.Data
}

// GetWidth returns the width.
func (p *Plot) GetWidth() int {
	return p.Width
}

// GetHeight returns the height.
func (p *Plot) GetHeight() int {
	return p.Height
}

// GetColor returns the color.
func (p *Plot) GetColor() color.Color {
	return p.Color
}

// GetName returns the name.
func (p *Plot) GetName() string {
	return p.Name
}

// GetFileName returns the file name.
func (p *Plot) GetFileName() string {
	return p.FileName
}

// SetData sets the data.
func (p *Plot) SetData(data []float64) {
	p.Data = data
}

// SetWidth sets the width.
func (p *Plot) SetWidth(width int) {
	p.Width = width
}

// SetHeight sets the height.
func (p *Plot) SetHeight(height int) {
	p.Height = height
}

// SetColor sets the color.
func (p *Plot) SetColor(color color.Color) {
	p.Color = color
}

// SetName sets the name.
func (p *Plot) SetName(name string) {
	p.Name = name
}

// SetFileName sets the file name.
func (p *Plot) SetFileName(fileName string) {
	p.FileName = fileName
}

// Histogram plots the histogram of the data.
func (p *Plot) Histogram() {
	// Get the min and max values.
	for _, v := range p.Data {
		if v < p.Min {
			p.Min = v
		}
		if v > p.Max {
			p.Max = v
		}
	}

	// Get the bin centers.
	p.BinCenters = make([]float64, p.Bins)
	for i := 0; i < p.Bins; i++ {
		p.BinCenters[i] = p.Min + (p.Max-p.Min)*float64(i)/float64(p.Bins)
	}

	// Get the bin counts.
	p.BinCounts = make([]int, p.Bins)
	for _, v := range p.Data {
		i := int(math.Floor((v - p.Min) / (p.Max - p.Min) * float64(p.Bins)))
		if i == p.Bins {
			i--
		}
		p.BinCounts[i]++
	}

	// Get the maximum bin count.
	maxBinCount := 0
	for _, v := range p.BinCounts {
		if v > maxBinCount {
			maxBinCount = v
		}
	}

	// Create the image.
	img := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Draw the histogram.
	for i := 0; i < p.Bins; i++ {
		x := int(float64(i)*p.BinWidth + p.BinWidth/2)
		y := int(float64(p.BinCounts[i]) / float64(maxBinCount) * float64(p.Height))
		for j := 0; j < y; j++ {
			img.Set(x, p.Height-j, p.Color)
		}
	}

	// Save the image.
	f, err := os.Create(p.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, img)
}

// Scatter plots the scatter plot of the data.
func (p *Plot) Scatter() {
	// Get the min and max values.
	for _, v := range p.Data {
		if v < p.Min {
			p.Min = v
		}
		if v > p.Max {
			p.Max = v
		}
	}

	// Create the image.
	img := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Draw the scatter plot.
	for _, v := range p.Data {
		x := int(float64(p.Width)*(v-p.Min)/(p.Max-p.Min)) + 1
		for j := 0; j < p.Height; j++ {
			img.Set(x, p.Height-j, p.Color)
		}
	}

	// Save the image.
	f, err := os.Create(p.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, img)
}

// Line plots the line plot of the data.
func (p *Plot) Line() {
	// Get the min and max values.
	for _, v := range p.Data {
		if v < p.Min {
			p.Min = v
		}
		if v > p.Max {
			p.Max = v
		}
	}

	// Create the image.
	img := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Draw the line plot.
	for i, v := range p.Data {
		x := int(float64(i)*float64(p.Width)/float64(len(p.Data))) + 1
		y := int(float64(p.Height)*(v-p.Min)/(p.Max-p.Min)) + 1
		for j := 0; j < y; j++ {
			img.Set(x, p.Height-j, p.Color)
		}
	}

	// Save the image.
	f, err := os.Create(p.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, img)
}

// Box plots the box plot of the data.
func (p *Plot) Box() {
	// Get the min and max values.
	for _, v := range p.Data {
		if v < p.Min {
			p.Min = v
		}
		if v > p.Max {
			p.Max = v
		}
	}

	// Get the median.
	sort.Float64s(p.Data)
	median := p.Data[len(p.Data)/2]

	// Get the first and third quartiles.
	firstQuartile := p.Data[len(p.Data)/4]
	thirdQuartile := p.Data[3*len(p.Data)/4]

	// Get the interquartile range.
	interquartileRange := thirdQuartile - firstQuartile

	// Get the lower and upper fences.
	lowerFence := firstQuartile - 1.5*interquartileRange
	upperFence := thirdQuartile + 1.5*interquartileRange

	// Create the image.
	img := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Draw the box plot.
	for i := 0; i < p.Width; i++ {
		img.Set(i, int(float64(p.Height)*(lowerFence-p.Min)/(p.Max-p.Min)), p.Color)
		img.Set(i, int(float64(p.Height)*(upperFence-p.Min)/(p.Max-p.Min)), p.Color)
	}
	for i := 0; i < p.Height; i++ {
		img.Set(int(float64(p.Width)*(firstQuartile-p.Min)/(p.Max-p.Min)), i, p.Color)
		img.Set(int(float64(p.Width)*(thirdQuartile-p.Min)/(p.Max-p.Min)), i, p.Color)
		img.Set(int(float64(p.Width)*(median-p.Min)/(p.Max-p.Min)), i, p.Color)
	}

	// Save the image.
	f, err := os.Create(p.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	png.Encode(f, img)
}
