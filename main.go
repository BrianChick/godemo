package main

import (
	"fmt"

	"github.com/BrianChick/godemo/math"
	"github.com/BrianChick/godemo/util"
)

func main() {
	x := util.GetX()
	y := util.GetY()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("blah")
	sum := math.Add(x, y)
	sum2 := math.SubtractNums(x, y)
	uuid := math.GetUUID()

	fmt.Println(sum)
	fmt.Println(sum2)
	fmt.Println(uuid)
}
