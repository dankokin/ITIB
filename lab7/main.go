package main

import (
	"fmt"
	"itib/lab7/network"
	"itib/lab7/symbol"
	"os"
)

var (
	symbol1 = symbol.StringToSymbol([]string{
		" + ",
		"++ ",
		" + ",
		" + ",
		"+++",
	}, 3, 5)
	symbol2 = symbol.StringToSymbol([]string{
		"+++",
		"  +",
		"+++",
		"+  ",
		"+++",
	}, 3, 5)
	symbol4 = symbol.StringToSymbol([]string{
		"+ +",
		"+ +",
		"+++",
		"  +",
		"  +",
	}, 3, 5)
	symbol8 = symbol.StringToSymbol([]string{
		" +++ ",
		"+   +",
		"+   +",
		" +++ ",
		"+   +",
		"+   +",
		" +++ ",
	}, 5, 7)
	symbolT = symbol.StringToSymbol([]string{
		"+++++",
		"  +  ",
		"  +  ",
		"  +  ",
		"  +  ",
		"  +  ",
		"  +  ",
	}, 5, 7)
	symbolU = symbol.StringToSymbol([]string{
		"+   +",
		"+   +",
		"+   +",
		"+   +",
		"+   +",
		"+   +",
		" +++ ",
	}, 5, 7)
)


func CatchSymbol(rnn *network.RecurrentNetwork, symbol symbol.Figure, isSync bool) {
	result := rnn.DetectFromSymbol(symbol, isSync)
	result.PrintSymbol(os.Stdout)

}
func main() {
	rnn15 := network.FromSymbols([]symbol.Figure{symbol2, symbol4}, 15, 100)

	CatchSymbol(&rnn15, symbol2, true)
	CatchSymbol(&rnn15, symbol4, true)

	fmt.Println(len(rnn15.Weights()))
	for _, row := range rnn15.Weights() {
		fmt.Print("|")
		for i, value := range row {
			if value >= 0 {
				fmt.Print(" ")
			}
			fmt.Print(value)
			if i != len(row) {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}



	fmt.Println("FIXING:")
	CatchSymbol(&rnn15, symbol.StringToSymbol([]string{
		" ++",
		"   ",
		"+++",
		"   ",
		"++ ",
	},3, 5), true)

	CatchSymbol(&rnn15, symbol.StringToSymbol([]string{
		"  +",
		"+ +",
		"+ +",
		"  +",
		"  +",
	},3, 5), true)

	fmt.Println()

	rnn35 := network.FromSymbols([]symbol.Figure{symbol8}, 35, 100)
	fmt.Println()
	for _, row := range rnn35.Weights() {
		fmt.Print("|")
		for i, value := range row {
			if value >= 0 {
				fmt.Print(" ")
			}
			fmt.Print(value)
			if i != len(row) {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}

	CatchSymbol(&rnn35, symbol.StringToSymbol([]string{
		" +++ ",
		"+   +",
		"+   +",
		"     ",
		"+   +",
		"+   +",
		" +++ ",
	},5, 7), true)
}


