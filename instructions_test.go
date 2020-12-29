package v60

import "os"

func ExamplePrintInstructions() {
	out := os.Stdout

	PrintPrep(out)
	PrintInstructions(
		NewBarista(out),
		newMockStopwatch(out),
		NewWaterWeights(15),
	)
	// Output: Warm the filter paper with hot water (tap will do)
	//Add coffee grounds
	//Create well in the coffee
	//Saturate the coffee with 30g of water
	//Swirl the coffee slurry until evenly mixed
	//Wait 45s for next prompt (elapsed 45s)
	//Add water until scales read 150g
	//Wait 30s for next prompt (elapsed 1m15s)
	//Add water until scales read 250g
	//Wait 30s for next prompt (elapsed 1m45s)
}
