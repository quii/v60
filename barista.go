package v60

import (
	"fmt"
	"io"
)

type Barista struct {
	out io.Writer
}

func NewBarista(out io.Writer) *Barista {
	return &Barista{out: out}
}

func (v Barista) BloomCoffee(amountOfWater float64) {
	fmt.Fprintf(v.out, "Saturate the coffee with %.0fg of water\n", amountOfWater)
	fmt.Fprintln(v.out, "Swirl the coffee slurry until evenly mixed")
}

func (v Barista) AddWater(amountOfWater float64) {
	fmt.Fprintf(v.out, "Add water until scales read %.0fg\n", amountOfWater)
}

func (v Barista) Stir() {
	fmt.Fprintln(v.out, "Stir clockwise once and then anticlockwise")
}
