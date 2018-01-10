package main

import (
	"fmt"
	"log"
	"net/http"

	chart "github.com/kharland/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Drawing")
	graph := chart.Heatmap{
		Width:  2000,
		Height: 1000,
		ColLabels: []string{
			"alpha", "bravo", "charlie", "delta", "eve", "fox",
			"garnet", "harry", "indigo",
		},
		RowLabels: []string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		},
		Grid: [][]float64{
			[]float64{1, 27, 3, 4, 135, 64, 7, 8, 9, 101},
			[]float64{11, 12, 13, 14, 15, 16, 17, 15, 19, 20},
			[]float64{1, 12, 113, 14, 98, 16, 1, 18, 19, 20},
			[]float64{11, 12, 73, 14, 15, 6, 17, 18, 19, 20},
			[]float64{13, 12, 3, 14, 15, 16, 17, 18, 193, 20},
			[]float64{111, 922, 3, 14, 35, 16, 17, 18, 493, 20},
			[]float64{0, 2, 3, 14, 5, 16, 17, 18, 193, 760},
			[]float64{67, 2, 3, 14, 150, 146, 07, 18, 191, 2},
			[]float64{9, 902, 3, 14, 15, 46, 17, 18, 1, 110},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	fmt.Println("Running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
