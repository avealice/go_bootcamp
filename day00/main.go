package main

import (
	"stats/metrics"
)

func main() {
	data := metrics.ReadNumbers()
	metrics.PrintMetrics(data)
}
