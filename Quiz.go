package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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

   correct:=0
   wrong:=0
   for i,problem:= range lines{
	   fmt.Printf("Problem %d) %s =",i+1,problem[0])
	   var answer string
	   fmt.Scanf("%s \n",&answer)
	   if answer == problem[1]{
		   correct++
	   }else{
			wrong++
	   }
   }
   fmt.Printf("Correct: %d, Wrong: %d",correct,wrong)
	

	
}
