package neuron

import (
	"fmt"
	"math"
)

type Neuron struct {
	input  []float64
	target []float64

	n uint8
	j uint8
	m uint8

	learningRate float64
	hiddenLayerWeights [][]float64
	outputLayerWeights [][]float64
}

func fillSlice(vector *[]float64, value float64) {
	for i := range *vector {
		(*vector)[i] = value
	}
}
func NewNeuron(input []float64, t []float64, n uint8, j uint8, m uint8, rate float64) Neuron {
	hiddenLayerWeights := make([][]float64, 0, j)
	for i := uint8(0); i < j; i++ {
		//weightsVector := []float64{0.5, 0.5, -0.4}
		weightsVector := make([]float64, n + 1)
		//fillSlice(&weightsVector, N)
		hiddenLayerWeights = append(hiddenLayerWeights, weightsVector)
	}

	outputLayerWeights := make([][]float64, 0, m)
	for i := uint8(0); i < m; i++ {
		weightsVector := make([]float64, j+1)
		//fillSlice(&weightsVector, N)
		outputLayerWeights = append(outputLayerWeights, weightsVector)
	}

	return Neuron{
		input:              input,
		target:             t,
		n:                  n,
		j:                  j,
		m:                  m,
		learningRate:       rate,
		hiddenLayerWeights: hiddenLayerWeights,
		outputLayerWeights: outputLayerWeights,
	}
}

func (n *Neuron) netInHiddenLayer(index int) float64 {
	net := n.hiddenLayerWeights[index][0]
	for i := uint8(0); i < n.n; i++ {
		net += n.hiddenLayerWeights[index][i + 1] * n.input[i]
	}

	return net
}

func (n *Neuron) netInOutputLayer(index int, inputInOutputLayer []float64) float64 {
	net := n.outputLayerWeights[index][0]
	for i := uint8(0); i < n.n; i++ {
		net += n.outputLayerWeights[index][i + 1] * inputInOutputLayer[i]
	}

	return net
}

func (n *Neuron) f(net float64) float64 {
	return (1 - math.Exp((-1) * net)) / (1 + math.Exp((-1) * net))
}

func (n *Neuron) derivative(net float64) float64 {
	return 0.5 * (1 - math.Pow(n.f(net), 2))
}

func (n *Neuron) getEpsilon(neuronValues []float64) float64 {
	var epsilon float64
	for i, target := range n.target {
		epsilon += math.Pow(target - neuronValues[i],2)
	}
	return math.Sqrt(epsilon)
}

func (n *Neuron) summarize(index int, deltaList []float64) float64 {
	var result float64
	for i, delta := range deltaList {
		result += n.outputLayerWeights[i][index] * delta
	}
	return result
}

func (n *Neuron) Train() {
	var epsilon = 100500.42
	epoch := 0
	for ; epsilon > 0.0001; {
		allNetsInHiddenLayer := make([]float64, 0, n.j)
		allNetsInOutputLayer := make([]float64, 0, n.m)
		inputInOutputLayer := make([]float64, n.j)
		out := make([]float64, 0, n.m)

		for j := 0; j < int(n.j); j++ {
			net := n.netInHiddenLayer(j)
			allNetsInHiddenLayer = append(allNetsInHiddenLayer, net)
			inputInOutputLayer[j] = n.f(net)
		}

		for m := 0; m < int(n.m); m++ {
			net := n.netInOutputLayer(m, inputInOutputLayer)
			allNetsInOutputLayer = append(allNetsInOutputLayer, net)
			out = append(out, n.f(net))
		}

		allOutErrors := make([]float64, 0, len(allNetsInOutputLayer))
		allHiddenErrors := make([]float64, 0, len(allNetsInHiddenLayer))

		for m, net := range allNetsInOutputLayer {
			derivative := n.derivative(net)
			delta := derivative * (n.target[m] - out[m])
			allOutErrors = append(allOutErrors, delta)
		}

		for j, net := range allNetsInHiddenLayer {
			derivative := n.derivative(net)
			delta := derivative * n.summarize(j, allOutErrors)
			allHiddenErrors = append(allHiddenErrors, delta)
		}

		for j := 0; j < int(n.j); j++ {
			n.hiddenLayerWeights[j][0] += n.learningRate * allHiddenErrors[j]
			for k := 0; k < int(n.n); k++ {
				n.hiddenLayerWeights[j][k + 1] += n.learningRate * n.input[k] * allHiddenErrors[j]
			}
		}

		for m := 0; m < int(n.m); m++ {
			n.outputLayerWeights[m][0] += n.learningRate * allOutErrors[m]
			for j := 0; j < int(n.j); j++ {
				n.outputLayerWeights[m][j + 1] += n.learningRate * inputInOutputLayer[j] * allOutErrors[m]
			}
		}

		epoch++
		epsilon = n.getEpsilon(out)
		fmt.Printf("%d. out: %.3f. Epsilon= %f \n", epoch, out, epsilon)
	}
}
