package v60

import (
	"os"
	"testing"
	"time"
)

func ExampleBrew() {
	out := os.Stdout

	barista := NewBarista(out)
	barista.PrepV60()
	Brew(
		barista,
		newMockStopwatch(out),
		NewWaterWeights(15),
	)
	// Output: Warm the filter paper with hot water
	//Add coffee grounds
	//Create well in the coffee
	//Saturate the coffee with 30g of water
	//Swirl the coffee slurry until evenly mixed
	//Wait 45s for next prompt (elapsed 45s)
	//Add water until scales read 150g
	//Wait 30s for next prompt (elapsed 1m15s)
	//Add water until scales read 250g
	//Stir clockwise once and then anticlockwise
}

func TestBrew(t *testing.T) {
	t.Run("it gives instructions for brewing a coffee with a v60 according to James Hoffman's guide", func(t *testing.T) {
		ratios := WaterWeights{
			BloomWaterWeight: 30,
			FirstPourWeight:  160,
			FinalPourWeight:  250,
		}
		spyBarista := &V60BaristaMock{
			BloomCoffeeFunc: func(amountOfWater float64) {},
			AddWaterFunc: func(amountOfWater float64) {},
			StirFunc: func() {},
		}
		//todo: maybe the waiting should live with Barista rather than it being separate, feels better
		dummyStopwatch := &StopwatchMock{
			WaitFunc: func(duration time.Duration) {},
		}

		Brew(spyBarista, dummyStopwatch, ratios)

		t.Run("Barista is told to bloom the water with the correct amount of water", func(t *testing.T) {
			numberOfTimesBaristaBloomed := len(spyBarista.BloomCoffeeCalls())
			if numberOfTimesBaristaBloomed != 1 {
				t.Errorf("should've bloomed once but actually bloomed %d times", numberOfTimesBaristaBloomed)
			}
			actualAmountOfWaterToBloomWith := spyBarista.BloomCoffeeCalls()[0].AmountOfWater
			if actualAmountOfWaterToBloomWith != ratios.BloomWaterWeight {
				t.Errorf(
					"barista was given the wrong amount of water to bloom with, got %f, wanted %f",
					actualAmountOfWaterToBloomWith,
					ratios.BloomWaterWeight,
				)
			}
		})

		t.Run("barista adds water twice", func(t *testing.T) {
			got := len(spyBarista.AddWaterCalls())
			if got != 2 {
				t.Fatalf("barista was supposed to add water twice but actually did it %d times", got)
			}
		})

		t.Run("Barista's first pour amount is correct", func(t *testing.T) {
			got := spyBarista.AddWaterCalls()[0].AmountOfWater
			if got != ratios.FirstPourWeight {
				t.Errorf("got %f, wanted %f", got, ratios.FirstPourWeight)
			}
		})

		t.Run("Barista's final pour amount is correct", func(t *testing.T) {
			got := spyBarista.AddWaterCalls()[1].AmountOfWater
			if got != ratios.FinalPourWeight {
				t.Errorf("got %f, wanted %f", got, ratios.FinalPourWeight)
			}
		})

	})
}
