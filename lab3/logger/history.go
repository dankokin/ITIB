package logger

import (
	"fmt"
	"io"
)

type HistoryPoint struct {
	epoch uint64
	epsilon float64
	weights []float64
}

func NewHistoryPoint(epoch uint64, epsilon float64, weights []float64) HistoryPoint {
	return HistoryPoint{
		epoch:   epoch,
		epsilon: epsilon,
		weights: weights,
	}
}

func (h *HistoryPoint) Log(output io.Writer) {
	line := fmt.Sprintf("Эпоха № %d. E = %f.6. Вектор весов: %.3f", h.epoch, h.epsilon, h.weights)
	fmt.Fprintln(output, line)
}
