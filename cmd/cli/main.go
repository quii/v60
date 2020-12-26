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
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter coffee weight: ")
	reader.Scan()
	coffeeWeight, err := strconv.ParseFloat(reader.Text(), 64)

	if err != nil {
		log.Fatal("Please enter a number for the coffee weight", err)
	}

	v60.PrintInstructions(os.Stdout, coffeeWeight)
}


