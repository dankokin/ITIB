package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func Round(x float64, accuracy float64) float64 {
	return math.Round(x*accuracy) / accuracy
}

func SendText(w http.ResponseWriter, statusCode int, text string) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprintln(w, text)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func SendJson(w http.ResponseWriter, statusCode int, outModel interface{}) {
	w.WriteHeader(statusCode)
	jsonModel, err := json.Marshal(outModel)
	if err != nil {
		_, err = fmt.Fprintln(w, err.Error())
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	_, err = fmt.Fprintln(w, string(jsonModel))
	if err != nil {
		fmt.Println(err.Error())
	}
}
