package app

import (
	"fmt"
	"strconv"
)

func ParseToFloat32(strValue string) float32 {
	floatValue, err := strconv.ParseFloat(strValue, 32)
	if err != nil {
		fmt.Println("Error in the parse value")
	}
	return float32(floatValue)
}
