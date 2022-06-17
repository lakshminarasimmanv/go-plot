 # plot

 Go package for Plotting.

 ## Installation

 ```
 go get github.com/lakshminarasimmanv/go-plot
 ```

 ## Usage

 ```
 import "github.com/lakshminarasimmanv/go-plot
 ```

 ### Example

 ```
 import (
         "github.com/lakshminarasimmanv/go-plot"
         "image/color"
 )

 func main() {

         data := []float64{
                 0.0,
                 1.0,
                 2.0,
                 3.0,
                 4.0,
                 5.0,
                 6.0,
                 7.0,
                 8.0,
                 9.0,
         }

         plot := plot.NewPlot(data, 800, 600, color.Black, "Example", "example.png")

         plot.Histogram()
         plot.Scatter()
         plot.Line()
         plot.Box()
 }
 ```

## How the program works, a step-by-step guide:

 1. Create a new plot.

    ``p := plot.NewPlot(data, width, height, color, name, fileName)``

 2. Set the number of bins to use for the plot.

    ``p.SetBins(bins)``

 3. Set the minimum value of the plot.

    ``p.SetMin(min)``

 4. Set the maximum value of the plot.

    ``p.SetMax(max)``

 5. Set the minimum and maximum values of the plot.

    ``p.SetMinMax(min, max)``

 6. Get the minimum and maximum values of the plot.

    ``min, max := p.GetMinMax()``

 7. Get the centers of the bins.

    ``binCenters := p.GetBinCenters()``

 8. Get the counts of the bins.

    ``binCounts := p.GetBinCounts()``

 9. Get the width of the bins.

    ``binWidth := p.GetBinWidth()``

 10. Get the height of the bins.

    ``binHeight := p.GetBinHeight()``

 11. Get the number of bins.

    ``bins := p.GetBins()``

 12. Get the data.

    ``data := p.GetData()``

 13. Get the width.

    ``width := p.GetWidth()``

 14. Get the height.

    ``height := p.GetHeight()``

 15. Get the color.

    ``color := p.GetColor()``

 16. Get the name.

    ``name := p.GetName()``

 17. Get the file name.

    ``fileName := p.GetFileName()``

 18. Set the data.

    ``p.SetData(data)``

 19. Set the width.

    ``p.SetWidth(width)``

 20. Set the height.

    ``p.SetHeight(height)``

 21. Set the color.

    ``p.SetColor(color)``

 22. Set the name.

    ``p.SetName(name)``

 23. Set the file name.

    ``p.SetFileName(fileName)``

 24. Plot the histogram of the data.

    ``p.Histogram()``

 25. Plot the scatter plot of the data.

    ``p.Scatter()``

 26. Plot the line plot of the data.

    ``p.Line()``

 27. Plot the box plot of the data.

    ``p.Box()``