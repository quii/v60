package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"v60"
)

func main() {
	out := v60.NewSayItOutLoud(os.Stdout)

	reader := bufio.NewScanner(os.Stdin)
	fmt.Fprint(out, "Enter coffee weight: ")
	reader.Scan()
	coffeeWeight, err := strconv.ParseFloat(reader.Text(), 64)

	if err != nil {
		log.Fatal("Please enter a number for the coffee weight", err)
	}

	barista := v60.NewBarista(out)

	barista.PrepV60()
	fmt.Fprint(out, "Hit return when prepped")
	reader.Scan()

	v60.Brew(
		barista,
		v60.NewRealStopwatch(out),
		v60.NewWaterWeights(coffeeWeight),
	)

	fmt.Fprint(out, "Wait for the drawdown to finish, enjoy your coffee")
}
