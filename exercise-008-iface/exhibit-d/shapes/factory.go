package shapes

import (
	"strings"
	"fmt"
)

func invalidParameters(shape string, count int, parameters []float64) bool {
	if (count != len(parameters)) {
		fmt.Println("Can't create", shape, "- invalid parameters:", parameters)
		return true
	}

	return false
}

func Build(shape string, parameters ...float64) (Shape) {
	const SHAPE_SQUARE = "square"
	const SHAPE_CIRCLE = "circle"
	const SHAPE_RECTANGLE = "rectangle"

	shape = strings.ToLower(shape)
	switch shape {
	case SHAPE_SQUARE:
		if (invalidParameters(shape, 1, parameters)) {
			return nil
		}

		return Square{Side: parameters[0]}
	case SHAPE_CIRCLE:
		if (invalidParameters(shape, 1, parameters)) {
			return nil
		}

		return Circle{Radius: parameters[0]}
	case SHAPE_RECTANGLE:
		if (invalidParameters(shape, 2, parameters)) {
			return nil
		}

		return Rectangle{Length: parameters[0], Width: parameters[1],}
	default:
		fmt.Println("Unsupported shape:", shape)
	}

	return nil
}