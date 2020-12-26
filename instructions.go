package v60

import (
	"fmt"
	"io"
)

func PrintInstructions(out io.Writer, coffeeWeight float64) {
	ratios := NewWaterWeights(coffeeWeight)

	fmt.Fprintln(out,"Warm the filter paper with hot water (tap will do)")
	fmt.Fprintln(out,"Add coffee grounds")
	fmt.Fprintln(out,"Create well in the coffee")
	fmt.Fprintln(out, "Start timer")
	fmt.Fprintf(out,"Satuate the coffee with %.0fg of water\n", ratios.BloomWaterWeight)
	fmt.Fprintln(out,"Swirl the coffee slurry until evenly mixed")
	fmt.Fprintln(out,"Wait until 45s")
	fmt.Fprintf(out,"Add more water, scales should now be at %.0fg\n", ratios.FirstPourWeight)
	fmt.Fprintln(out, "Wait until 1 min 15s")
	fmt.Fprintf(out, "Add remaining water up to %.0fg, pour slowly\n", ratios.FinalPourWeight)
	fmt.Fprintln(out, "At 1 min 45s stir clockwise once and then anticlockwise")
	fmt.Fprintln(out, "Wait for the drawdown, around 3 minutes 30s")
}