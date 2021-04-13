package logger

import (
	"fmt"
	"io"
)

type HistoryPoint struct {
	Epoch       uint64
	NeuronError uint64
	Weights     []float64

	Target []uint8
	Out    []uint8
}

func NewHistoryPoint(epoch uint64, epsilon uint64, weights []float64, t []uint8, out []uint8) HistoryPoint {
	w := make([]float64, len(weights), cap(weights))
	copy(w, weights)
	return HistoryPoint{
		Epoch:       epoch,
		NeuronError: epsilon,
		Weights:     w,
		Target:      t,
		Out:         out,
	}
}

func (h *HistoryPoint) Log(output io.Writer) {
	line := fmt.Sprintf("Эпоха № %d. E = %d. Вектор весов: %.3f. \nT: %v.\nY: %v\n",
		h.Epoch, h.NeuronError, h.Weights, h.Target, h.Out)
	fmt.Fprintln(output, line)
}

