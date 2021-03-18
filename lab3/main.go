package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"itib/lab3/functions"
	"itib/lab3/neuron"
	"itib/lab3/utils"
)

func main() {
	outputToFile := flag.Bool("f", true, "True if need to write to file. False if need to write to Stdout")
	flag.Parse()

	var output io.Writer
	if *outputToFile {
		file, err := os.Create("results/result.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		output = file
	} else {
		output = os.Stdout
	}


	f := functions.NewFunction(func(x float64) float64 {
		return math.Exp((-0.1) * x * x)
	})

	trainX, trainY := f.CalculateFunctionInterval(-5, 5, 20)
	predictX, predictY := f.CalculateFunctionInterval(5, 15, 20)
	trainY = append(trainY, predictY...)
	trainX = append(trainX, predictX...)

	n := neuron.CreateNeuron(8, trainY, 0.3, 20)
	n.Train(4000)
	n.Predict(20)
	n.Log(output)

	neuronY := n.GetNeuronValues()

	plotter := utils.CreatePlotter("График функции и ее прогноз",
		utils.CoordinateAxes{
		XLabel: "x",
		YLabel: "y",
	}, 8, 6)

	realFunctionLine := utils.Line{
		XPoints:    trainX,
		YPoints:    trainY,
		ColorModel: utils.ColorModel{
			R: 255,
			G: 0,
			B: 0,
			A: 0,
		},
		LineName:   "Целевая функция",
	}

	predictFunctionLine := utils.Line{
		XPoints:    trainX,
		YPoints:    neuronY,
		ColorModel: utils.ColorModel{
			R: 0,
			G: 255,
			B: 0,
			A: 0,
		},
		LineName:   "Прогноз функции",
	}

	err := plotter.DrawGraph([]utils.Line{realFunctionLine, predictFunctionLine}, "results/graph1.png")
	if err != nil {
		fmt.Println(err)
	}
}
