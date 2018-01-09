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
			"a", "b", "c", "d", "e",
		},
		RowLabels: []string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		},
		Grid: [][]float64{
			[]float64{1, 27, 3, 4, 135, 64, 7, 8, 9, 101},
			[]float64{11, 12, 13, 14, 15, 16, 17, 15, 19, 20},
			[]float64{1, 12, 113, 14, 98, 16, 1, 18, 19, 20},
			[]float64{11, 12, 73, 14, 15, 6, 17, 18, 19, 20},
			[]float64{1, 92, 3, 14, 15, 16, 17, 18, 193, 20},
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
