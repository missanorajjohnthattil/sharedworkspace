package main

// Milestone 1: Read the csv file. Define a structure and load the values into objects. Hint: https://github.com/gocarina/gocsv
import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type ArrayValue struct {
	Name         string `csv:"name"`
	Mpg          int    `csv:"mpg"`
	Cylinders    int    `csv:"cylinders"`
	Displacement int    `csv:"displacement"`
	Horsepower   int    `csv:"horsepower"`
	Weight       int    `csv:"weight"`
	Acceleration int    `csv:"acceleration"`
	Model_year   int    `csv:"model_year"`
	Origin       string `csv:"origin"`
}

func main() {
	fmt.Println("Work Pool menu ")
	readFile("automobile.csv")
}

func readFile(fileName string) {
	fileValues, fileErr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if fileErr != nil {
		panic(fileErr)
	}
	defer fileValues.Close()

	ArrayValues := []ArrayValue{}

	if fileErr := gocsv.UnmarshalFile(fileValues, &ArrayValues); fileErr != nil { // Load clients from file
		panic(fileErr)
	}
}

// Milestone 2: Create a job(logic to calculate car with max magic value)

func (ar ArrayValue) getMagic() {

}

// Milestone 3: Create a workpool. Spawn n number of go routines and pass m number of records to process. Each job(routine) can calculate the max value and corresponding car name with the batch of values it processed.
// Milestone 4: Wait for all jobs to complete.
// Milestone 5: Given that each job has calculated individual max magic value, find a way to get the aggregated max.
// Milestone 6: Implement a jobs count limiter. There should be hold on max number of jobs that can run concurrently at a time.
// Milestone 7: Make workpool related code as separate package. You can design the package such that any csv file should work (given structure definition) and any logic should work.
