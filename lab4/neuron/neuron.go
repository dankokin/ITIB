package neuron

import (
	"io"
	"itib/lab4/logger"
	"itib/lab4/utils"
	"math"
)

type ActivationFunction interface {
	Activate(float64) uint8
	Derivative(float64) float64
}

// (1, 3, 5, 6, 8)

type Neuron struct {
	activationFunction ActivationFunction
	teachingRate       float64
	weights            []float64

	target            []uint8
	sets              [][]uint8
	cSet              [][]uint8
	variablesQuantity uint8

	history []logger.HistoryPoint
}

func CreateNeuron(function ActivationFunction, teachRate float64, target []uint8, varsQuantity uint8) *Neuron {
	sets := utils.MakeAllSets(varsQuantity)

	cSetIndexes := utils.GetCSet(target)
	cSet := make([][]uint8, 0, len(cSetIndexes))
	for _, index := range cSetIndexes {
		cSet = append(cSet, sets[index])
	}

	return &Neuron{
		activationFunction: function,
		weights:            make([]float64, len(cSet) + 1),
		teachingRate:       teachRate,
		target:             target,
		sets:               sets,
		variablesQuantity:  varsQuantity,
		cSet:               cSet,
	}
}


func (n *Neuron) getPhi(variablesSet []uint8, cSet []uint8) float64 {
	var phi float64
	for i := range variablesSet {
		phi += math.Pow(float64(variablesSet[i]) - float64(cSet[i]), 2)
	}
	return math.Exp(-phi)
}

func (n *Neuron) getNet(variablesSet []uint8) float64 {
	net := n.weights[len(n.cSet)]
	for j := range n.cSet {
		net += n.weights[j] * n.getPhi(variablesSet, n.cSet[j])
	}
	return net
}

func (n *Neuron) getSetIndex(set []uint8) uint64 {
	var index float64
	for i, digit := range set {
		index += float64(digit) * math.Pow(2, float64(len(set) - i - 1))
	}
	return uint64(index)
}

func (n *Neuron) getDelta(target uint8, y uint8) float64 {
	return float64(target) - float64(y)
}

func (n *Neuron) getActivationFunction(net float64) uint8 {
	return n.activationFunction.Activate(net)
}

func (n *Neuron) calculateFunctionVector() []uint8 {
	vector := make([]uint8, 0, len(n.target))
	for i := 0; i < len(n.target); i++ {
		net := n.getNet(n.sets[i])
		value := n.getActivationFunction(net)
		vector = append(vector, value)
	}
	return vector
}

func (n *Neuron) Train(maxIterations uint64, sets ...[]uint8) bool {
	trainSets := make ([][]uint8, 0, len(n.sets))

	if sets == nil {
		trainSets = n.sets
	} else {
		trainSets = sets
	}
	n.history = make([]logger.HistoryPoint, 0, maxIterations)

	for i := uint64(0); i < maxIterations; i++ {
		yVector := n.calculateFunctionVector()
		distance := utils.HammingDistance(yVector, n.target)
		n.history = append(n.history, logger.NewHistoryPoint(i, distance, n.weights, n.target, yVector))
		if distance == 0 {
			return true
		}

		for _, trainSet := range trainSets {
			net := n.getNet(trainSet)
			y := n.getActivationFunction(net)
			index := n.getSetIndex(trainSet)

			t := n.target[index]
			delta := n.getDelta(t, y)
			n.weights[len(n.cSet)] += n.teachingRate * delta * n.activationFunction.Derivative(net)

			for j, cSet := range n.cSet {
				n.weights[j] += n.teachingRate *
					delta * n.getPhi(trainSet, cSet) * n.activationFunction.Derivative(net)
			}
		}
	}

	return false
}

func (n *Neuron) Log(writer io.Writer) {
	for _, point := range n.history {
		point.Log(writer)
	}
}

func (n *Neuron) GetHistoryPoints() []logger.HistoryPoint {
	return n.history
}

