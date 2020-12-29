package v60

import (
	"time"
)

//go:generate moq -out stopwatch_moq_test.go . Stopwatch
type Stopwatch interface {
	Wait(duration time.Duration)
}

//go:generate moq -out v60_barista_moq_test.go . V60Barista
type V60Barista interface {
	PrepV60()
	BloomCoffee(amountOfWater float64)
	AddWater(amountOfWater float64)
	Stir()
}

func Brew(barista V60Barista, stopwatch Stopwatch, ratios WaterWeights) {
	barista.BloomCoffee(ratios.BloomWaterWeight)
	stopwatch.Wait(45 * time.Second)

	barista.AddWater(ratios.FirstPourWeight)
	stopwatch.Wait(30 * time.Second)

	barista.AddWater(ratios.FinalPourWeight)
	barista.Stir()
}
