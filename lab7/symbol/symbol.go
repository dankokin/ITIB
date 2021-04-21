package symbol

import (
	"fmt"
	"io"
)

const (
	Value = '+'
	Void = ' '
)


type Figure struct {
	width  uint64
	height uint64
	data   []int
}


func CreateFigure(w uint64, h uint64, data []int) Figure {
	return Figure{
		width:  w,
		height: h,
		data:   data,
	}
}


func StringToSymbol(rows []string, w uint64, h uint64) Figure {
	symbolData := make([]int, w * h)

	for i := uint64(0); i < h; i++ {
		for j := uint64(0); j < w; j++ {
			if rows[i][j] == Value {
				symbolData[h * j + i] = 1
			} else {
				symbolData[h * j + i] = -1
			}
		}
	}
	return Figure{
		width:  w,
		height: h,
		data:   symbolData,
	}
}


func (s *Figure) Height() uint64 {
	return s.height
}


func (s *Figure) Width() uint64 {
	return s.width
}


func (s *Figure) Data() []int {
	return s.data
}
func (s *Figure) ToString() []string {
	stringSymbol := make([]string, s.Height())
	for j := uint64(0); j < s.Height(); j++ {
		for i := uint64(0); i < s.Width(); i++ {
			if s.data[s.height*i+j] > 0 {
				stringSymbol[j] += string(Value)
			} else {
				stringSymbol[j] += string(Void)
			}
		}
	}
	return stringSymbol
}


func (s *Figure) PrintSymbol(writer io.Writer) {
	for j := uint64(0); j < s.Height(); j++ {
		for i := uint64(0); i < s.Width(); i++ {
			if s.data[s.Height() * i + j] > 0 {
				fmt.Fprint(writer, string(Value))
			} else {
				fmt.Fprint(writer, string(Void))
			}
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer)
}
