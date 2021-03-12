package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"itib/lab1/neuron"
)

type thresholdActivationFunction struct {
}

func (t *thresholdActivationFunction) Derivative(w []float64, set []uint8) float64 {
	return 1
}
func (t *thresholdActivationFunction) Activate(net float64) uint8 {
	if net >= 0 {
		return 1
	} else {
		return 0
	}
}

type sigmoidActivationFunction struct {
}

func (s *sigmoidActivationFunction) Activate(net float64) uint8 {
	if (1 + math.Tanh(net)) / 2 >= 0.5 {
		return 1
	} else {
		return 0
	}
}

func (s *sigmoidActivationFunction) Derivative(w []float64, set []uint8) float64 {
	net := w[0]
	for i, weight := range w[1:] {
		net += weight * float64(set[i])
	}
	return (1 - math.Pow(math.Tanh(net), 2)) / 2
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
	neuronWithThresholdActivationFunctions := neuron.CreateNeuron(&thresholdActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.3,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		output)
	neuronWithThresholdActivationFunctions.Train(100, false, "results/neuronWithThresholdActivationFunctions.png")

	// -0.4449454703530218, 0.5924632297479917, 0.14400496434762708, 0.2962520287466416, 0.43264427686926255
	// -0.8999999999999999, 1.0078181144620333, 0.18878066418306055, 0.24495723896961066, 0.7696645608241754
	fmt.Fprintln(output, "\nЗадание 2. Обучение с сигмоидальной функцией активации")
	var sigmoidActivationFunction sigmoidActivationFunction
	neuronWithSigmoidFunctions := neuron.CreateNeuron(&sigmoidActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.4,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		output)
	neuronWithSigmoidFunctions.Train(100, false, "results/neuronWithSigmoidFunctions.png")


	fmt.Fprintln(output, "\nЗадание 3. Обучение с сигмоидальной функцией активации и неполной выборкой")
	neuronWithPartlyLearning := neuron.CreateNeuron(&sigmoidActivationFunction,
		[]float64{0.0, 0.0, 0.0, 0.0, 0.0},
		0.3,
		[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		4,
		output)
	neuronWithPartlyLearning.TrainPartly(100, "results/neuronWithPartlyLearning.png")

}
