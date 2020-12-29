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

func (b Barista) PrepV60() {
	fmt.Fprintln(b.out, "Warm the filter paper with hot water")
	fmt.Fprintln(b.out, "Add coffee grounds")
	fmt.Fprintln(b.out, "Create well in the coffee")
}

func (b Barista) BloomCoffee(amountOfWater float64) {
	fmt.Fprintf(b.out, "Saturate the coffee with %.0fg of water\n", amountOfWater)
	fmt.Fprintln(b.out, "Swirl the coffee slurry until evenly mixed")
}

func (b Barista) AddWater(amountOfWater float64) {
	fmt.Fprintf(b.out, "Add water until scales read %.0fg\n", amountOfWater)
}

func (b Barista) Stir() {
	fmt.Fprintln(b.out, "Stir clockwise once and then anticlockwise")
}
