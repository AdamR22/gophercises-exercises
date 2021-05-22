package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Struct holding default arguments
// Golang functions can't hold default arg values
type defaultFuncArgmunets struct {
	newFileName string
	quizTime    int
}

func quizGame(args defaultFuncArgmunets) {

	filePath := "/home/adam/Desktop/gophercises-exercises/files/problems.csv"
	var newFilePath string
	var input string

	quizScore := 0

	//Change file name if user chooses so
	if args.newFileName != "nil" {
		newFilePath = fmt.Sprintf("../files/%s", args.newFileName)
		err := os.Rename(filePath, newFilePath)

		if err == nil {
			filePath = newFilePath
		} else {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	recordLength := len(records)

	doneChannel := make(chan bool)

	go func() {
		for i := 0; i < recordLength; i++ {
			problem := records[i]

			question := problem[0]
			answer := problem[1]

			fmt.Printf("What's the solution to %s? ", question)
			fmt.Scan(&input)

			if input == answer {
				quizScore++
			}

			if i == recordLength-1 {
				doneChannel <- true
				fmt.Println("Final Score: ", quizScore)
				os.Exit(0)
			}
		}
	}()

	go func() {
		timer := time.NewTimer(time.Duration(args.quizTime) * time.Second)
		<-timer.C
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Your time is up...")
		fmt.Println("Final Score: ", quizScore)
		os.Exit(0)
	}()

	<-doneChannel
}

func commandLineParser() defaultFuncArgmunets {
	args := defaultFuncArgmunets{}

	fileName := flag.String("filename", "nil", "Change Filename")
	quizDuration := flag.Int("duration", 30, "Quiz Duration")

	flag.Parse()

	args.newFileName = *fileName
	args.quizTime = *quizDuration

	return args
}

func main() {

	functionArgs := commandLineParser()
	quizGame(functionArgs)

}
