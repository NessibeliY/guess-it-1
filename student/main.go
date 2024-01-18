package main

import (
	"bufio"
	"fmt"
	"guess-it-1/student/prediction"
	"log"
	"os"
	"strconv"
)

func main() {
	consoleScanner := bufio.NewScanner(os.Stdin)
	slice := []int{}
	for consoleScanner.Scan() {
		input := consoleScanner.Text()
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			log.Print(err)
			return
		}
		if inputInt >= 100 && inputInt <= 200 {
			slice = append(slice, inputInt)
		}
		predictedRange := prediction.PredictRange(slice)
		fmt.Printf("%d %d\n", predictedRange[0], predictedRange[1])
	}
}
