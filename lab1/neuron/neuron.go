package neuron

import (
	"fmt"
	"io"

	"itib/lab1/combinations"
	"itib/lab1/utils"
)

type Neuron struct {
	activationFunction ActivationFunction
	derivative Derivative
	teachingRate float64

	weights []float64
	target []uint8

	sets [][]uint8
	variablesQuantity uint8

	writer io.Writer
}

func CreateNeuron(function ActivationFunction, weights []float64,
	teachRate float64, target []uint8, varsQuantity uint8, derivative Derivative, writer io.Writer) *Neuron {
	sets := utils.MakeAllSets(varsQuantity)
	return &Neuron{
		activationFunction: function,
		weights:            weights,
		teachingRate:       teachRate,
		target:             target,
		sets: sets,
		variablesQuantity: varsQuantity,
		writer: writer,
		derivative: derivative,
	}
}

func (n *Neuron) SetOutput(writer io.Writer) {
	n.writer = writer
}

func (n *Neuron) getActivationFunction(set []uint8) uint8 {
	net := n.weights[0]
	for i, weight := range n.weights[1:] {
		net += weight * float64(set[i])
	}
	return n.activationFunction.Result(net)
}

func (n * Neuron) calculateFunctionVector() []uint8 {
	vector := make([]uint8, 0, len(n.target))
	for i := 0; i < 1 << n.variablesQuantity; i++ {
		value := n.getActivationFunction(n.sets[i])
		vector = append(vector, value)
	}
	return vector
}

func (n *Neuron) PrintInfo(epoch uint16, err uint8, out []uint8) {
	info := fmt.Sprintf("Эпоха № %d. Выходной вектор: %v. Вектор весов: %.3f. Суммарная ошибка: %d",
		epoch, out, n.weights, err)
	fmt.Fprintln(n.writer, info)
}

func (n *Neuron) Train(epochs uint16, isPartly bool, graphicName string, sets ...[]uint8) bool {
	xPoints := make([]float64, 0, epochs)
	yPoints := make([]float64, 0, len(n.target))

	for epoch := uint16(0); epoch < epochs; epoch++ {
		vector := n.calculateFunctionVector()
		err := utils.HammingDistance(n.target, vector)

		xPoints = append(xPoints, float64(epoch))
		yPoints = append(yPoints, float64(err))

		if !isPartly {
			n.PrintInfo(epoch, err, vector)
		}
		if err == 0 {
			if isPartly {
				fmt.Fprintf(n.writer, "Минимальный набор из %d векторов: %v\n", len(sets), sets)
			}
			p := utils.CreatePlotter()
			p.DrawGraph(xPoints, yPoints, len(xPoints), graphicName, 100, 10, 255)
			return true
		}
		var teachSet [][]uint8
		if len(sets) == 0 {
			teachSet = n.sets
		} else {
			teachSet = sets
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < len(teachSet); j++ {
				if i == 0 {
					n.weights[i] += n.teachingRate * (float64(n.target[j]) - float64(vector[j])) *
						n.derivative.Result(n.weights, teachSet[j])
				} else {
					n.weights[i] += n.teachingRate * (float64(n.target[j]) - float64(vector[j])) *
						float64(teachSet[j][i-1]) * n.derivative.Result(n.weights, teachSet[j])
				}
			}
		}
	}
	return false
}

func (n *Neuron) TrainPartly(epochs uint16, graphicName string) bool {
	for i := 2; i < 16; i++ {
		setCombinations := combinations.Combinations(n.sets, i)
		for _, setCombination := range setCombinations {
			n.weights = []float64{0.0, 0.0, 0.0, 0.0, 0.0}
			result := n.Train(100, true, graphicName, setCombination...)
			if result {
				n.weights = []float64{0.0, 0.0, 0.0, 0.0, 0.0}
				n.Train(100, false, graphicName, setCombination...)
				return true
			}
		}
	}
	return false
}

type ActivationFunction interface {
	Result(float64) uint8
}

type Derivative interface {
	Result([]float64, []uint8) float64
}
