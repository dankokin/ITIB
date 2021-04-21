package network

import (
	"itib/lab7/symbol"
)


type RecurrentNetwork struct {
	weights [][]int
	k uint64
	maxIterations uint64
}


func CreateRecurrentNetwork(x [][]int, k uint64, maxEpochs uint64) RecurrentNetwork {
	weights := make([][]int, k)
	for i := uint64(0); i < k; i++ {
		weights[i] = make([]int, k)
		for j := uint64(0); j < k; j++ {
			if i == j {
				break
			}

			for l := range x {
				weights[i][j] += x[l][i] * x[l][j]
			}

			weights[j][i] = weights[i][j]
		}
	}

	return RecurrentNetwork{
		weights:       weights,
		k:             k,
		maxIterations: maxEpochs,
	}
}


func FromSymbols(symbols []symbol.Figure, k uint64, maxEpochs uint64) RecurrentNetwork {
	x := make([][]int, len(symbols))

	for l := range symbols {
		x[l] = symbols[l].Data()
	}
	return CreateRecurrentNetwork(x, k, maxEpochs)
}


func (rnn *RecurrentNetwork) Weights() [][]int {
	return rnn.weights
}


func (rnn *RecurrentNetwork) GetOut(x []int, isSync bool) []int {
	out := make([]int, rnn.k)
	for i := range x {
		out[i] = x[i]
	}

	var prevOut []int
	var checkOut []int
	if isSync {
		prevOut = make([]int, rnn.k)
		checkOut = prevOut
	} else {
		prevOut = out
		checkOut = make([]int, rnn.k)
	}

	for epoch := uint64(0); epoch < rnn.maxIterations; epoch++ {
		if isSync {
			out, prevOut = prevOut, out
			checkOut = prevOut
		} else {
			copy(checkOut, prevOut)
		}

		for k := uint64(0); k < rnn.k; k++ {
			net := 0
			for j := uint64(0); j < rnn.k; j++ {
				if j == k {
					continue
				}
				net += rnn.weights[j][k] * prevOut[j]
			}

			if net < 0 {
				out[k] = -1
			} else if net > 0 {
				out[k] = 1
			} else {
				out[k] = prevOut[k]
			}
		}

		isDiff := false
		for i := uint64(0); i < rnn.k; i++ {
			if out[i] != checkOut[i] {
				isDiff = true
			}
		}

		if !isDiff {
			break
		}
	}

	return out
}


func (rnn *RecurrentNetwork) DetectFromSymbol(s symbol.Figure, isSync bool) symbol.Figure {
	return symbol.CreateFigure(s.Width(), s.Height(), rnn.GetOut(s.Data(), isSync))
}

