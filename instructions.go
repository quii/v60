package v60

import (
	"fmt"
	"io"
	"time"
)

type Stopwatch interface {
	Wait(duration time.Duration)
}

type V60Barista interface {
	BloomCoffee(amountOfWater float64)
	AddWater(amountOfWater float64)
	Stir()
}

func PrintInstructions(barista V60Barista, stopwatch Stopwatch, ratios WaterWeights) {
	barista.BloomCoffee(ratios.BloomWaterWeight)
	stopwatch.Wait(45 * time.Second)

	barista.AddWater(ratios.FirstPourWeight)
	stopwatch.Wait(30 * time.Second)

	barista.AddWater(ratios.FinalPourWeight)
	stopwatch.Wait(30 * time.Second)
}

func PrintPrep(out io.Writer) {
	fmt.Fprintln(out, "Warm the filter paper with hot water (tap will do)")
	fmt.Fprintln(out, "Add coffee grounds")
	fmt.Fprintln(out, "Create well in the coffee")
}
