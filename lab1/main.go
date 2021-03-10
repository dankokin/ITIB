package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"itib/lab1/neuron"
)


type thresholdDerivativeFunction struct {
}

func (td *thresholdDerivativeFunction) Result (w []float64, set []uint8) float64 {
	return 1
}

type sigmoidDerivativeFunction struct {
}

type thresholdActivationFunction struct {
}

func (t *thresholdActivationFunction) Result(net float64) uint8 {
	if net >= 0 {
		return 1
	} else {
		return 0
	}
}

type sigmoidActivationFunction struct {
}

func (s *sigmoidActivationFunction) Result(net float64) uint8 {
	if net >= 0.5 {
		return 1
	} else {
		return 0
	}
}

func (sd *sigmoidDerivativeFunction) Result (w []float64, set []uint8) float64 {
	net := w[0]
	for i := range w[1:] {
		net += w[i] * float64(set[i])
	}
	return 1 / (2 * math.Pow(math.Cos(net), 2))
}

func main() {
	outputToFile := flag.Bool("f", false, "True if need to write to file. False if need to write to Stdout")
	flag.Parse()

	var output io.Writer
	if *outputToFile {
		file, err := os.Create("results/report.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		output = file
	} else {
		output = os.Stdout
	}

	fmt.Fprintln(output, "Задание 1. Обучение с пороговой функцией активации")
	var thresholdActivationFunction thresholdActivationFunction
	var thresholdDerivativeFunction thresholdDerivativeFunction
	neuronWithThresholdActivationFunctions := neuron.CreateNeuron(&thresholdActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.3,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		&thresholdDerivativeFunction,
		output)
	neuronWithThresholdActivationFunctions.Train(100, false, "results/neuronWithThresholdActivationFunctions.png")

	fmt.Fprintln(output, "\nЗадание 2. Обучение с сигмоидальной функцией активации")
	var sigmoidActivationFunction sigmoidActivationFunction
	var sigmoidDerivativeFunction sigmoidDerivativeFunction
	neuronWithSigmoidFunctions := neuron.CreateNeuron(&sigmoidActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.3,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		&sigmoidDerivativeFunction,
		output)
	neuronWithSigmoidFunctions.Train(100, false, "results/neuronWithSigmoidFunctions.png")

	fmt.Fprintln(output, "\nЗадание 3. Обучение с сигмоидальной функцией активации и неполной выборкой")
	neuronWithPartlyLearning := neuron.CreateNeuron(&sigmoidActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.3,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		&sigmoidDerivativeFunction,
		output)
	neuronWithPartlyLearning.TrainPartly(100, "results/neuronWithPartlyLearning.png")
}
