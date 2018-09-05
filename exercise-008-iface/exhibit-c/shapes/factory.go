package shapes

import (
	"fmt"
	"os"
	"strings"
)

type BuildShape func(parameters ...float64) Shape

var registeredBuildShapes = make(map[string]BuildShape)

func init() {
	registeredBuildShapes["square"] = BuildSquare
	registeredBuildShapes["circle"] = BuildCircle
	registeredBuildShapes["rectangle"] = BuildRectangle
}

func BuildSquare(parameters ...float64) Shape {
	if len(parameters) != 1 {
		fmt.Println("Square should have only 1 parameter")
		os.Exit(1)
	}

	return &Square{Side: parameters[0]}
}

func BuildCircle(parameters ...float64) Shape {
	if len(parameters) != 1 {
		fmt.Println("Circle should have only 1 parameter")
		os.Exit(1)
	}

	return &Circle{Radius: parameters[0]}
}

func BuildRectangle(parameters ...float64) Shape {
	if len(parameters) != 2 {
		fmt.Println("A rectangle should have 2 params - Length & Breadth")
		os.Exit(1)
	}

	return &Rectangle{Length: parameters[0], Breadth: parameters[1]}
}

func Build(shape string, parameters ...float64) Shape {
	shape = strings.ToLower(shape)
	buildMethod, ok := registeredBuildShapes[shape]
	if !ok {
		fmt.Println("No shape called", shape)
		os.Exit(1)
	}
	return buildMethod(parameters...)
}
