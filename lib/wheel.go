package lib

import (
	"fmt"
	"strings"
)
	
type WheelSlice  struct {
	Label string `json:"label"`
	Color string `json:"backgroundColor"`
}

var NegList []string
var PosList []string

// type WheelPie []WheelSlice
func ParseStringMessage(message string) {

	tasks := strings.Split(message, "\n")

	for i,task  := range tasks {
		fmt.Printf("\t%d:[%s]\n", i, task)
	}
}
