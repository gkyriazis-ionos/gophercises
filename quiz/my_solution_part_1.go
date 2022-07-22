package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./quiz/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(inputFile)
	problemCounter := 0
	correctAnsCounter := 0
	for {
		problemCounter += 1
		data, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Problem #" + strconv.Itoa(problemCounter) + ":" + " " + data[0])
		var ans string
		_, err = fmt.Scanf("%s", &ans)
		if err != nil {
			log.Fatal(err)
		}
		if ans == data[1] {
			correctAnsCounter += 1
		}
	}
	fmt.Println("Correct Answers: " + strconv.Itoa(correctAnsCounter))
	fmt.Println("Wrong Answers: " + strconv.Itoa(problemCounter-correctAnsCounter-1))

}
