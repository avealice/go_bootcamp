package tests

import (
	"math"
	"stats/metrics"
	"testing"
)

func TestMetrics(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		data := metrics.ReadNumbersFromFile("test1.txt")

		mean := data.GetMean()
		trueMean := 1.50
		if mean != trueMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, trueMean)
		}
		median := data.GetMedian()
		trueMedian := 1.50
		if median != trueMedian {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", median, trueMedian)
		}
		mode := data.GetMode()
		trueMode := 0
		if mode != trueMode {
			t.Errorf("[Expected] %d != %d [Real]\n", mode, trueMode)
		}
		SD := data.GetSD()
		trueSD := 1.118033988749895
		if math.Abs(SD-trueSD) > 0.00001 {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", SD, trueSD)
		}
	})

	t.Run("Test2", func(t *testing.T) {
		data := metrics.ReadNumbersFromFile("test2.txt")

		mean := data.GetMean()
		trueMean := 4.0
		if mean != trueMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, trueMean)
		}
		median := data.GetMedian()
		trueMedian := 3.5
		if median != trueMedian {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", median, trueMedian)
		}
		mode := data.GetMode()
		trueMode := 0
		if mode != trueMode {
			t.Errorf("[Expected] %d != %d [Real]\n", mode, trueMode)
		}
		SD := data.GetSD()
		trueSD := 4.041451884327381
		if math.Abs(SD-trueSD) > 0.00001 {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", SD, trueSD)
		}
	})

	t.Run("Test3", func(t *testing.T) {
		data := metrics.ReadNumbersFromFile("test3.txt")

		mean := data.GetMean()
		trueMean := 4.0
		if mean != trueMean {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", mean, trueMean)
		}
		median := data.GetMedian()
		trueMedian := 4.0
		if median != trueMedian {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", median, trueMedian)
		}
		mode := data.GetMode()
		trueMode := 2
		if mode != trueMode {
			t.Errorf("[Expected] %d != %d [Real]\n", mode, trueMode)
		}
		SD := data.GetSD()
		trueSD := 1.5811388300841898
		if math.Abs(SD-trueSD) > 0.00001 {
			t.Errorf("[Expected] %.2f != %.2f [Real]\n", SD, trueSD)
		}
	})

}
