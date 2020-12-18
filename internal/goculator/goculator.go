package goculator

import (
	"strings"

	"github.com/j1mmyson/Goculator/internal/calculator"
	"github.com/j1mmyson/Goculator/internal/check"
	"github.com/j1mmyson/Goculator/internal/display"
)

// Goculator is main function of Goculator
func Goculator() {
	storage := make([]float64, 0)
	var (
		input  string
		result float64
	)

	for {
		display.Display(storage, result, input)
		input = check.ReadInput()
		input = check.CallStorage(input, storage) // 저장된 값을 불러오는지 체크하고 인풋을 그에맞게 변경해줌.
		if strings.HasPrefix(input, "x") {
			display.ShowEndPage()
			break
		}
		calculator.Calculate(input, &result, &storage)
	}
}
