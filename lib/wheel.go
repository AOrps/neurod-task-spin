package lib

import (
	"fmt"
	"strings"
	"github.com/muesli/gamut"
)
	
type WheelSlice  struct {
	Label string `json:"label"`
	Color string `json:"backgroundColor"`
}

var NegList []string
var PosList []string

// type WheelPie []WheelSlice
func ParseStringMessage(message string) []WheelSlice {

	var wheelPie []WheelSlice
	
	tasks := strings.Split(message, "\n")

	tasksCount := len(tasks)
	hexColors := GenColors(tasksCount, gamut.PastelGenerator{})
	
	for i  := range tasks {
		wheelPie = append(wheelPie, WheelSlice{
			Label : tasks[i],
			Color : hexColors[i],
		})
	}

	for _,val := range wheelPie {
		fmt.Printf("<%s>:[%s]\n", val.Label, val.Color)
	}

	// for _,hc := range hexColors {
	// 	fmt.Println(hc)
	// }

	return wheelPie
}


// generates hexcolors via generator
func GenColors(n int, generator gamut.ColorGenerator) []string {

	var retList []string
	
	colors, err := gamut.Generate(n, generator)

	if err != nil {
		fmt.Println(err)
	}

	for _,color := range colors {
		retList = append(retList, gamut.ToHex(color))
	}

	return retList
}
