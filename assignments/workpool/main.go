package main

//(done) Milestone 1: Read the csv file. Define a structure and load the values into objects. Hint: https://github.com/gocarina/gocsv
import (
	"fmt"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
)

type carDetail struct {
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

var wg sync.WaitGroup
var largestMutex sync.Mutex
var largest int
var name string

func main() {
	fmt.Println("Work Pool menu ")
	carDetails := readFile("automobile.csv")
	workPlanner(carDetails, 20, 5)
}

func readFile(fileName string) []carDetail {
	fileValues, fileErr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if fileErr != nil {
		panic(fileErr)
	}
	defer fileValues.Close()
	carDetails := []carDetail{}

	if fileErr := gocsv.UnmarshalFile(fileValues, &carDetails); fileErr != nil { // Load clients from file
		panic(fileErr)
	}
	return carDetails
}

//(done) Milestone 3: Create a workpool. Spawn n number of go routines and pass m number of records to process. Each job(routine) can calculate the max value and corresponding car name with the batch of values it processed.

func workPlanner(ArrayValues []carDetail, spliceLength int, workerLimit int) {
	var spliceEnd int
	waitChan := make(chan struct{}, workerLimit)
	for spliceHead := 0; spliceHead <= len(ArrayValues); spliceHead += spliceLength {
		spliceEnd = spliceHead + spliceLength
		if spliceEnd > len(ArrayValues) {
			spliceEnd = len(ArrayValues)
		}
		fmt.Printf(" \n Length : %d , Head : %d, Tail : %d \n", spliceLength, spliceHead, spliceEnd)
		startWork(waitChan)
		go getMagicCodes(ArrayValues[spliceHead:spliceEnd], waitChan)
	}
	wg.Wait()
	fmt.Printf("\n largest of the pool is : \t %d, and the model is : \t %s \n\n", largest, name)
}

func startWork(waitChan chan (struct{})) {
	waitChan <- struct{}{}
	wg.Add(1)
}

func stopWork(waitChan chan (struct{})) {
	<-waitChan
	wg.Done()
}

//(done) Milestone 2: Create a job(logic to calculate car with max magic value)

func getMagicCodes(arR []carDetail, waitChan chan (struct{})) {
	defer stopWork(waitChan)

	localLargest := 0
	var localName string
	for _, ar := range arR {
		var magicCodeInt int
		//magicCodeInt = (ar.Acceleration + ar.Cylinders + ar.Displacement + ar.Horsepower + ar.Model_year + ar.Mpg + ar.Weight) * len(ar.Name)
		magicCodeInt = len(ar.Name)
		//get the unique character count
		if magicCodeInt > localLargest {
			localLargest = magicCodeInt
			localName = ar.Name
		}
	}
	checkAndUpdate(localLargest, localName)
}

func checkAndUpdate(localLargest int, localName string) {
	largestMutex.Lock()
	if localLargest > largest {
		largest = localLargest
		name = localName
	}
	largestMutex.Unlock()
}

//(done) Milestone 4: Wait for all jobs to complete.
//(done) Milestone 5: Given that each job has calculated individual max magic value, find a way to get the aggregated max.
//(done) Milestone 6: Implement a jobs count limiter. There should be hold on max number of jobs that can run concurrently at a time.
// Milestone 7: Make workpool related code as separate package. You can design the package such that any csv file should work (given structure definition) and any logic should work.
