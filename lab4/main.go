package main

import (
	"flag"
	"fmt"
	"io"
	"itib/lab4/combinations"
	"itib/lab4/utils"
	"math"
	"os"

	"itib/lab4/neuron"
)

type thresholdActivationFunction struct {
}

func (t *thresholdActivationFunction) Derivative(net float64) float64 {
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

func (s *sigmoidActivationFunction) Derivative(net float64) float64 {
	return 1 / (2 * math.Cosh(net) * math.Cosh(net))
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

	fmt.Fprintln(output, "Задание 1. Обучение с пороговой функцией активации\n")
	var thresholdActivationFunction thresholdActivationFunction
	neuronWithThresholdActivationFunctions := neuron.CreateNeuron(&thresholdActivationFunction,
		0.3,
		[]uint8{0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
		4,
		)
	neuronWithThresholdActivationFunctions.Train(200)
	neuronWithThresholdActivationFunctions.Log(output)

	history := neuronWithThresholdActivationFunctions.GetHistoryPoints()
	xPoints := make([]float64, len(history))
	yPoints := make([]float64, len(history))
	for i := range history {
		xPoints[i] = float64(history[i].Epoch)
		yPoints[i] = float64(history[i].NeuronError)
	}

	plotter := utils.CreatePlotter("Зависимость ошибок нейросети от количества эпох",
		utils.CoordinateAxes{
			XLabel: "K",
			YLabel: "E(K)",
		}, 8, 6)

	realFunctionLine := utils.Line{
		XPoints:    xPoints,
		YPoints:    yPoints,
		ColorModel: utils.ColorModel{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		},
		LineName:   "E(K)",
	}

	err := plotter.DrawGraph([]utils.Line{realFunctionLine}, "results/threshold.png")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(output, "Задание 2. Обучение с сигмоидальной функцией\n")
	var sigmoid sigmoidActivationFunction
	neuronWithSigmoid := neuron.CreateNeuron(&sigmoid,
		0.3,
		[]uint8{0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
		4,
	)
	neuronWithSigmoid.Train(200)
	neuronWithSigmoid.Log(output)

	sigmoidHistory := neuronWithSigmoid.GetHistoryPoints()
	sigmoidXPoints := make([]float64, len(history))
	sigmoidYPoints := make([]float64, len(history))
	for i := range history {
		sigmoidXPoints[i] = float64(sigmoidHistory[i].Epoch)
		sigmoidYPoints[i] = float64(sigmoidHistory[i].NeuronError)
	}

	sigmoidPlotter := utils.CreatePlotter("Зависимость ошибок нейросети от количества эпох",
		utils.CoordinateAxes{
			XLabel: "K",
			YLabel: "E(K)",
		}, 8, 6)

	sigmoidRealFunctionLine := utils.Line{
		XPoints:    sigmoidXPoints,
		YPoints:    sigmoidYPoints,
		ColorModel: utils.ColorModel{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		},
		LineName:   "E(K)",
	}

	err = sigmoidPlotter.DrawGraph([]utils.Line{sigmoidRealFunctionLine}, "results/sigmoid.png")
	if err != nil {
		fmt.Println(err)
	}


	fmt.Fprintln(output, "Задание 3. Нахождение минимальной обучающей выборки\n")

	for i := 1; i < 16; i++ {
		allCombinations := combinations.Combinations(utils.MakeAllSets(4), i)
		for _, combinationSet := range allCombinations {
			n := neuron.CreateNeuron(&thresholdActivationFunction,
			0.3,
			[]uint8{0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
			4,
			)
			ok := n.Train(200, combinationSet...)
			if ok {
				fmt.Println(combinationSet)
				n.Log(output)
				shortestHistory := n.GetHistoryPoints()
				shortestXPoints := make([]float64, len(shortestHistory))
				shortestYPoints := make([]float64, len(shortestHistory))
				for i := range shortestHistory {
					shortestXPoints[i] = float64(shortestHistory[i].Epoch)
					shortestYPoints[i] = float64(shortestHistory[i].NeuronError)
				}

				shortestPlotter := utils.CreatePlotter("Зависимость ошибок нейросети от количества эпох",
					utils.CoordinateAxes{
						XLabel: "K",
						YLabel: "E(K)",
					}, 8, 6)

				shortestRealFunctionLine := utils.Line{
					XPoints:    shortestXPoints,
					YPoints:    shortestYPoints,
					ColorModel: utils.ColorModel{
						R: 255,
						G: 255,
						B: 0,
						A: 0,
					},
					LineName:   "E(K)",
				}

				err = shortestPlotter.DrawGraph([]utils.Line{shortestRealFunctionLine}, "results/shortest.png")
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}
	}





	//neuronWithMinSets := neuron.CreateNeuron(&thresholdActivationFunction,
	//	0.3,
	//	[]uint8{0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	//	4,
	//)

	//neuronWithMinSets.Train(200, [][]uint8{{0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 0, 1}, {0, 1, 1, 0}, {1, 0, 0, 0}}...)
	//neuronWithMinSets.Log(output)
}
