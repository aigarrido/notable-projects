package main


import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"sync"
)

// Filter defining
var sharpenBlur = [3][3]float64{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}

var ridgeBlur = [3][3]float64{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

var kernelBoxBlur = [3][3]float64{
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
}

// Kernel mapping
var kernels = map[int][3][3]float64{
	0: sharpenBlur,
	1: ridgeBlur,
	2: kernelBoxBlur,

}

func main() {
	// Instructions by flags
	kernelChoice := flag.Int("k", 0, "Kernel")
	inputFile := flag.String("i", "input.jpg", "Input")
	outputFile := flag.String("o", "output.jpg", "Output")
	numThreads := flag.Int("t", 4, "N_threads")

	// Flag parse
	flag.Parse()

	// Open image
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	imagen, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Image bounds
	bounds := imagen.Bounds()

	// Result image 
	convolutedImage := image.NewRGBA(bounds)

	// Kernel error management
	kernel, exists := kernels[*kernelChoice]
	if !exists {
		fmt.Println("Invalid kernel choice")
		return
	}

	// Gorutine waitgroup
	var wg sync.WaitGroup

	// Flag element to int
	numSections := *numThreads

	// Section separation based on number of threads
	sectionHeight := bounds.Max.Y / numSections

	for i := 0; i < numSections; i++ {
		startY := i * sectionHeight
		endY := (i + 1) * sectionHeight

		if i == numSections-1 {
			endY = bounds.Max.Y
		}

		wg.Add(1)

		// Gorutines
		go func(startY, endY int) {
			defer wg.Done()
			applyConvolution(imagen, convolutedImage, startY, endY, kernel)
		}(startY, endY)
	}

	// Wait for all gorutines to finish
	wg.Wait()


	// Save image
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	jpeg.Encode(outFile, convolutedImage, nil)
}

// Applies the convolution to the image
func applyConvolution(imagen image.Image, convolutedImage *image.RGBA, startY, endY int, kernel [3][3]float64) {
	bounds := imagen.Bounds()
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := startY; y < endY && y < bounds.Max.Y-1; y++ {
			var rSum, gSum, bSum float64
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					r, g, b, _ := imagen.At(x+i, y+j).RGBA()
					rSum += float64(r) * kernel[i+1][j+1]
					gSum += float64(g) * kernel[i+1][j+1]
					bSum += float64(b) * kernel[i+1][j+1]
				}
			}

			// Pixel clamp
			rSum = clamp(rSum, 0, 65535)
			gSum = clamp(gSum, 0, 65535)
			bSum = clamp(bSum, 0, 65535)
			convolutedImage.Set(x, y, color.RGBA64{uint16(rSum), uint16(gSum), uint16(bSum), 65535})
		}
	}
}

// Adjusts min and max values of the pixels
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
