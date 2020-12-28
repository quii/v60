package v60

import "os"

func ExamplePrintInstructions() {
	out := os.Stdout
	PrintPrep(out)
	PrintInstructions(out, newMockStopwatch(out), 15)
	// Output: Warm the filter paper with hot water (tap will do)
	//Add coffee grounds
	//Create well in the coffee
	//Saturate the coffee with 30g of water
	//Swirl the coffee slurry until evenly mixed
	//Wait 45s for next prompt (elapsed 45s)
	//Add more water until scales read 150g
	//Wait 30s for next prompt (elapsed 1m15s)
	//Add remaining water up to 250g, pour slowly
	//Wait 30s for next prompt (elapsed 1m45s)
	//Stir clockwise once and then anticlockwise
	//Wait for the drawdown, around 3 minutes 30s
}
