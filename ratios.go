package v60

type WaterWeights struct {
	BloomWaterWeight float64
	FirstPourWeight  float64
	FinalPourWeight  float64
}

func NewWaterWeights(coffeeGroundWeight float64) WaterWeights {
	finalPourWeight := (coffeeGroundWeight / 6) * 100

	return WaterWeights{
		BloomWaterWeight: coffeeGroundWeight * 2,
		FirstPourWeight:  (finalPourWeight / 100) * 60,
		FinalPourWeight:  finalPourWeight,
	}
}
