package main

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func RunGame(pA Strategy, pB Strategy, numIter int) (scoresA []float64, scoresB []float64, err error) {

	fmt.Println("Starting Game!")

	scoreA := 0.0
	scoreB := 0.0

	scoresA = make([]float64, 0)
	scoresB = make([]float64, 0)

	for i := 0; i < numIter; i++ {
		var pAAction Action
		var pBAction Action

		fmt.Printf("Iteration %d\n", i)

		pAAction, err = pA.Decide()
		if err != nil {
			return
		}

		pBAction, err = pB.Decide()
		if err != nil {
			return
		}

		examine := NewExamine()

		sentence := examine.CrossExamine(pAAction, pBAction)

		scoreA += sentence.forA
		scoreB += sentence.forB

		scoresA = append(scoresA, scoreA)
		scoresB = append(scoresB, scoreB)

		fmt.Printf("Prisoner A sentenced to %g current total %g\n", sentence.forA, scoreA)
		fmt.Printf("Prisoner B sentenced to %g current total %g\n", sentence.forB, scoreB)

		pA.SetOpponentDecision(pBAction)
		pB.SetOpponentDecision(pAAction)
	}

	return
}

func main() {
	upA := NewRandom()
	upB := NewAlways(false)

	scoresA, scoresB, err := RunGame(upA, upB, 10)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Iter\t%s\t%s\n", upA.prisoner.name, upB.prisoner.name)
	for i := 0; i < len(scoresA); i++ {
		fmt.Printf("%g\t%g\t%g\n", float64(i), scoresA[i], scoresB[i])
	}

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title:    upA.prisoner.name + " vs " + upB.prisoner.name,
			Subtitle: "Accumulation of sentence per iteration of Prisoners Dilemma",
		}))

	// Put data into instance
	line.SetXAxis([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
		AddSeries("Category A", generateLineItems(scoresA)).
		AddSeries("Category B", generateLineItems(scoresB)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Step: true}))

	f, _ := os.Create("bar.html")
	line.Render(f)
}

func generateLineItems(data []float64) []opts.LineData {
	fmt.Println(data)
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}
