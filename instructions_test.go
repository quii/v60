package v60

import "os"

func ExamplePrintInstructions() {
	PrintInstructions(os.Stdout, 15)
	// Output: Warm the filter paper with hot water (tap will do)
	//Add coffee grounds
	//Create well in the coffee
	//Start timer
	//Satuate the coffee with 30g of water
	//Swirl the coffee slurry until evenly mixed
	//Wait until 45s
	//Add more water, scales should now be at 150g
	//Wait until 1 min 15s
	//Add remaining water up to 250g, pour slowly
	//At 1 min 45s stir clockwise once and then anticlockwise
	//Wait for the drawdown, around 3 minutes 30s
}
