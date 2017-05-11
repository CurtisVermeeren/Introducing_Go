package main
import (
	"fmt"
	"Introducing_Go/Chapter_8_Packages/createPackages/math"
)

func main(){
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
