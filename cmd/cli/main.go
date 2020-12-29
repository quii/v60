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
	fmt.Print("Enter coffee weight: ")
	reader.Scan()
	coffeeWeight, err := strconv.ParseFloat(reader.Text(), 64)

	if err != nil {
		log.Fatal("Please enter a number for the coffee weight", err)
	}

	v60.PrintPrep(out)

	fmt.Print("Hit return when prepped")
	reader.Scan()

	v60.PrintInstructions(v60.NewBarista(out), v60.NewRealStopwatch(out), coffeeWeight)
}
