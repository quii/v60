package v60

import (
	"fmt"
	"io"
	"time"
)

func PrintPrep(out io.Writer) {
	fmt.Fprintln(out, "Warm the filter paper with hot water (tap will do)")
	fmt.Fprintln(out, "Add coffee grounds")
	fmt.Fprintln(out, "Create well in the coffee")
}

type Stopwatch interface {
	Wait(duration time.Duration)
}

func PrintInstructions(out io.Writer, stopwatch Stopwatch, coffeeWeight float64) {
	ratios := NewWaterWeights(coffeeWeight)

	fmt.Fprintf(out, "Saturate the coffee with %.0fg of water\n", ratios.BloomWaterWeight)
	fmt.Fprintln(out, "Swirl the coffee slurry until evenly mixed")
	stopwatch.Wait(45 * time.Second)

	fmt.Fprintf(out, "Add more water until scales read %.0fg\n", ratios.FirstPourWeight)
	stopwatch.Wait(30 * time.Second)

	fmt.Fprintf(out, "Add remaining water up to %.0fg, pour slowly\n", ratios.FinalPourWeight)
	stopwatch.Wait(30 * time.Second)

	fmt.Fprintln(out, "Stir clockwise once and then anticlockwise")
	fmt.Fprintln(out, "Wait for the drawdown, around 3 minutes 30s")
}
