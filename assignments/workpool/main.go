package main

//(done) Milestone 1: Read the csv file. Define a structure and load the values into objects. Hint: https://github.com/gocarina/gocsv
import (
	"fmt"
	"os"
	"sync"
	"time"

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

var largest int
var name string
var wg sync.WaitGroup
var wrkrLmt int

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
	workPlanner(ArrayValues, 20, 5)
}

//(done) Milestone 2: Create a job(logic to calculate car with max magic value)

func getMagicCodes(arR []ArrayValue) {
	defer stopWork()
	largest = 0
	for _, ar := range arR {
		name = ""
		var magicCodeInt int
		magicCodeInt = ar.Acceleration + ar.Cylinders + ar.Displacement + ar.Horsepower + ar.Model_year + ar.Mpg + ar.Weight + len(ar.Name)
		//get the unique character count
		if magicCodeInt > largest {
			largest = magicCodeInt
			name = ar.Name
		}
	}
	fmt.Printf("\n\t\t Largest of pool is %s with value : %d \n", name, largest)
}

//(done) Milestone 3: Create a workpool. Spawn n number of go routines and pass m number of records to process. Each job(routine) can calculate the max value and corresponding car name with the batch of values it processed.

func workPlanner(ArrayValues []ArrayValue, spliceLength int, workerLimit int) {
	var spliceEnd int
	for spliceHead := 0; spliceHead <= len(ArrayValues); spliceHead += spliceLength {
		spliceEnd = spliceHead + spliceLength
		if spliceEnd > len(ArrayValues) {
			spliceEnd = len(ArrayValues)
		}
		fmt.Printf(" \n Length : %d , Head : %d, Tail : %d \n", spliceLength, spliceHead, spliceEnd)
		startWork(workerLimit)
		go getMagicCodes(ArrayValues[spliceHead:spliceEnd])
	}
	wg.Wait()
}

func startWork(workerLimit int) {
	if wrkrLmt >= workerLimit {
		time.Sleep(1 * time.Second)
	}
	wrkrLmt++
	wg.Add(1)
	fmt.Println("Work added : ", wrkrLmt)
}

func stopWork() {
	wrkrLmt--
	wg.Done()
	fmt.Println("Work done : ", wrkrLmt)
}

//(done) Milestone 4: Wait for all jobs to complete.
//(done) Milestone 5: Given that each job has calculated individual max magic value, find a way to get the aggregated max.
// Milestone 6: Implement a jobs count limiter. There should be hold on max number of jobs that can run concurrently at a time.
// Milestone 7: Make workpool related code as separate package. You can design the package such that any csv file should work (given structure definition) and any logic should work.
