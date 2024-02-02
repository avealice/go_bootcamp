package metrics

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type Flag struct {
	printMean   bool
	printMedian bool
	printMode   bool
	printSD     bool
}

type Data struct {
	Numbers []int
	lenght  int
}

func NewData(numbers []int) *Data {
	return &Data{Numbers: numbers, lenght: len(numbers)}
}

func ReadNumbers() *Data {
	array := Input(os.Stdin)

	return NewData(array)
}

func ReadNumbersFromFile(fileName string) *Data {
	file, _ := os.Open(fileName)
	defer file.Close()
	array := Input(file)
	return NewData(array)
}

func (d *Data) GetMean() float64 {
	sum := 0.0
	for _, num := range d.Numbers {
		sum += float64(num)
	}
	return float64(sum) / float64(d.lenght)
}

func (d *Data) GetMedian() float64 {
	if d.lenght == 0 {
		return math.NaN()
	}
	mid1 := d.lenght / 2
	if d.lenght%2 == 1 {
		return float64(d.Numbers[mid1])
	}
	mid2 := d.lenght/2 - 1
	return float64((d.Numbers[mid1] + d.Numbers[mid2])) / 2.0
}

func (d *Data) GetMode() int {
	counts := make(map[int]int)

	for _, num := range d.Numbers {
		counts[num]++
	}

	maxCount := 0
	var modeValue int

	for num, count := range counts {
		if count > maxCount || (count == maxCount && num < modeValue) {
			maxCount = count
			modeValue = num
		}
	}

	return modeValue
}

func (d *Data) GetSD() float64 {
	mean := d.GetMean()

	var squaredDifferences float64

	for _, num := range d.Numbers {
		squaredDifferences += math.Pow(float64(num)-mean, 2)
	}

	sd := math.Sqrt(squaredDifferences / float64(d.lenght))

	return sd
}

func PrintMetrics(data *Data) {
	if data.lenght == 0 {
		return
	}

	var f Flag

	flag.BoolVar(&f.printMean, "mean", false, "mean output")
	flag.BoolVar(&f.printMedian, "median", false, "median output")
	flag.BoolVar(&f.printMode, "mode", false, "mode output")
	flag.BoolVar(&f.printSD, "SD", false, "SD output")

	flag.Parse()

	if !f.printMean && !f.printMedian && !f.printMode && !f.printSD {
		f.printMean, f.printMedian, f.printMode, f.printSD = true, true, true, true
	}

	if f.printMean {
		mean := data.GetMean()
		fmt.Printf("Mean: %.2f\n", mean)
	}

	if f.printMedian {
		median := data.GetMedian()
		fmt.Printf("Median: %.2f\n", median)
	}

	if f.printMode {
		mode := data.GetMode()
		fmt.Printf("Mode: %d\n", mode)
	}

	if f.printSD {
		SD := data.GetSD()
		fmt.Printf("SD: %.2f\n", SD)
	}
}
