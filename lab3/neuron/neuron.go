package neuron

import (
	"fmt"
	"io"
	"itib/lab3/logger"
	"math"
)

type Neuron struct {
	weights      []float64
	neuronValues []float64
	realValues   []float64

	windowLength   uint64
	teachingRate   float64
	pointsQuantity uint64

	history []logger.HistoryPoint
}

func CreateNeuron(windowLength uint64, funcValues []float64, n float64, points uint64) *Neuron {
	neuronValues := make([]float64, points, len(funcValues))
	for i := uint64(0); i < windowLength; i++ {
		neuronValues[i] = funcValues[i]
	}

	return &Neuron{
		weights:        make([]float64, windowLength+1),
		neuronValues:   neuronValues,
		realValues:     funcValues,
		windowLength:   windowLength,
		history:        make([]logger.HistoryPoint, 0, points),
		teachingRate:   n,
		pointsQuantity: points,
	}
}

func (n *Neuron) getDelta(realValue float64, neuronValue float64) float64 {
	return realValue - neuronValue
}

func (n *Neuron) getEpsilon(lastIndex uint64) float64 {
	var epsilon float64
	for i := n.windowLength; i < lastIndex; i++ {
		delta := n.getDelta(n.realValues[i], n.neuronValues[i])
		epsilon += delta * delta
	}
	return math.Sqrt(epsilon)
}

func (n *Neuron) Train(maxIterations uint64) {
	for finishEpoch := uint64(0); finishEpoch < maxIterations; finishEpoch++ {

		for i := n.windowLength; i < n.pointsQuantity; i++ {
			net := n.weights[n.windowLength]
			for j := uint64(0); j < n.windowLength; j++ {
				net += n.weights[j] * n.realValues[i-n.windowLength+j]
			}
			n.neuronValues[i] = net
			for k := uint64(0); k < n.windowLength; k++ {
				n.weights[k] += n.teachingRate * n.getDelta(n.realValues[i], n.neuronValues[i]) *
					n.realValues[i-n.windowLength+k]
			}
			//n.weights[n.windowLength] += n.teachingRate * n.getDelta(n.realValues[i], n.neuronValues[i])
			if n.getEpsilon(n.pointsQuantity) < 0.0001 {
				fmt.Println(n.getEpsilon(n.pointsQuantity))
				break
			}
		}
		n.history = append(n.history, logger.NewHistoryPoint(finishEpoch, n.getEpsilon(n.pointsQuantity), n.weights))
	}
}

func (n *Neuron) Predict(quantityToPredict uint64) {
	for i := n.pointsQuantity; i < n.pointsQuantity+quantityToPredict; i++ {
		net := n.weights[n.windowLength]
		for j := uint64(0); j < n.windowLength; j++ {
			net += n.weights[j] * n.neuronValues[i-n.windowLength+j]
		}
		n.neuronValues = append(n.neuronValues, net)
	}
}

func (n *Neuron) Log(output io.Writer) {
	for _, value := range n.history {
		value.Log(output)
	}
}

func (n *Neuron) GetNeuronValues() []float64 {
	return n.neuronValues
}

func (n *Neuron) GetWeights() []float64 {
	return n.weights
}
