package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func praseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i].question = line[0]
		ret[i].answer = line[1]
	}
	return ret
}

func main() {
	// Open the file i can't find file?
	file, err := os.Open("Problems.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Read the file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	correct := 0
	wrong := 0
	problems := praseLines(lines)
	for i, problem := range problems {
		fmt.Printf("Problem #%d) %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == strings.TrimSpace(problem.answer) {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("Correct: %d, Wrong: %d", correct, wrong)

}
