package functions

import (
	"math"
)

type Function struct {
	function func(float64) float64
}

func NewFunction(f func(float64) float64) Function {
	return Function{
		function: f,
	}
}

func (f *Function) SetFunction(function func(float64) float64) {
	f.function = function
}

func (f *Function) CalculateFunctionInterval(begin float64, end float64, pointsQuantity uint64) ([]float64, []float64) {
	if math.Abs(end - begin) < 1e-8 || pointsQuantity == 0 {
		return nil, nil
	}

	step := (end - begin) / (float64(pointsQuantity) - 1)

	functionResults := make([]float64, 0, pointsQuantity)
	points := make([]float64, 0, pointsQuantity)
	for i := uint64(0); i < pointsQuantity; i++ {
		points = append(points, begin + step * float64(i))
		functionResults = append(functionResults, f.function(begin + step * float64(i)))
	}
	//for x := begin; math.Abs(x - end) < 1e-8; x += step {
	//	points = append(points, x)
	//	functionResults = append(functionResults, f.function(x))
	//}

	return points, functionResults
}
