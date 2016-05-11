package main

import "fmt"
import "math"

func main() {

	var entropy float64 = 0.0
	var proba int = 0
	var probas = [...]float64{104.0, 101.0, 108.0, 108.0, 111.0, 27.0, 
		119.0, 110.0, 114.0, 108.0, 100.0}

	for proba < len(probas) {

		entropy -= probas[proba] * math.Log2(probas[proba])
		proba++
	}

	fmt.Printf("Hello World Entropique: %f\n", entropy)
}
