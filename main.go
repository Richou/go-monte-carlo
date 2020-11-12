package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var numberOccurrence int
	var verbose bool
	circle := 0
	pi := 0.0

	flag.IntVar(&numberOccurrence, "occurrence", 100, "Number of occurence")
	flag.BoolVar(&verbose, "v", false, "Verbose display")

	flag.Parse()

	for i := 1; i <= numberOccurrence; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if d <= 1 {
			circle += 1
		}
		pi = float64(circle) / float64(i) * 4
		if verbose {
			fmt.Printf("Occurrence number: %d, pi estimation : %f\n", i, pi)
		}
	}
	fmt.Println(pi)
}
