package v60

import "testing"

func TestInstructions(t *testing.T) {

	t.Run("final pour weight should be 60g per litre (or 100g water per 6 grams of coffee)", func(t *testing.T) {
		got := NewWaterWeights(15).FinalPourWeight
		want := 250.0
		AssertWeight(t, got, want)
	})

	t.Run("the bloom water weight is 2x the ground weight", func(t *testing.T) {
		got := NewWaterWeights(50).BloomWaterWeight
		want := 100.0
		AssertWeight(t, got, want)
	})

	t.Run("the first pour weight should be 60% of the final pour weight", func(t *testing.T) {
		got := NewWaterWeights(60).FirstPourWeight
		want := 600.0
		AssertWeight(t, got, want)
	})

	t.Run("example ratios", func(t *testing.T) {
		cases := []struct {
			Description  string
			GrindWeight  float64
			WantedRatios WaterWeights
		}{
			{"From the James Hoffman video", 30, WaterWeights{
				BloomWaterWeight: 60,
				FirstPourWeight:  300,
				FinalPourWeight:  500,
			}},
			{"Single coffee", 15, WaterWeights{
				BloomWaterWeight: 30,
				FirstPourWeight:  150,
				FinalPourWeight:  250,
			}},
		}

		for _, test := range cases {
			t.Run(test.Description, func(t *testing.T) {
				got := NewWaterWeights(test.GrindWeight)
				if got != test.WantedRatios {
					t.Errorf("got %v, want %v", got, test.WantedRatios)
				}
			})
		}
	})
}

func AssertWeight(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
