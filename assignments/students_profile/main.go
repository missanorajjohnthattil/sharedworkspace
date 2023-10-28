package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Pid       int    `json:"MemberID"`
	Name      string `json:"MemberName"`
	Age       int    `json:"MemberAge"`
	PhoneNo   string `json:"PhoneNo"`
	MagicCode int    `json:"MagicCode"`
}

type Group struct {
	Gid    int            `json:"GroupID"`
	Gname  string         `json:"GroupName"`
	Groups map[int]Person `json:"Group"`
}

type Registery struct {
	Rid    int           `json:"RegisterID"`
	RName  string        `json:"RegisterName"`
	RGroup map[int]Group `json:"groupRegister"`
}

var Rid, Gid, Pid int
var Rgstrs = make(map[int]Registery)

func init() {
	getJsonValues("sample.json")
}

func main() {
	defer writeJsonValues(Rgstrs, "sample.json")
	var exitCheck bool = false
	if exitCheck {
		return
	} else {
		for exitCheck == false {
			exitCheck = RegisterUI()
		}
	}
}

func RegisterUI() bool {
	fmt.Printf("\n\t\tWelcome registry manager\n\n")
	for {
		var option string
		fmt.Printf(`
		Action No -- Option detail
		1 ---------- !Add a member to group
		2 ---------- !View members in group
		3 ---------- !Delete Member
		5 ---------- !Add new group
		6 ---------- !View group details
		7 ---------- !Delete group
		10 --------- !Add Registry details
		11 --------- !View Registry details
		12 --------- !Delete Registry
		15 --------- !Create the Magic code
		16 --------- Get the Magic code
		20 --------- Save json file
		21 --------- View json file
		99 ---------- Exit from the function

		Select your action
		`)
		fmt.Scanf("%s", &option)
		optionStoA, err := strconv.Atoi(option)
		if err != nil {
			fmt.Printf("\nYou have not entered a valid option.\nPlease try again")
			return false
		} else {
			switch optionStoA {
			case 16:
				viewMagicCode()
			case 20:
				writeJsonValues(Rgstrs, "sample.json")
			case 21:
				getJsonValues("sample.json")
			case 99:
				println("Happy to see you later")
				return true
			default:
				println("\n\n You have entered a wrong entry \n\n")
			}
		}
	}
}

// Function to Open the json file and unmarshal the values
func getJsonValues(fileName string) {
	println("Json file reader is fetching the details")
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n %s Successfully Opened \n\n", fileName)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var readValue map[int]Registery

	json.Unmarshal(byteValue, &readValue)
	fmt.Printf("\n\n Lenght : \t%d\n", len(readValue))

	for i := 1; i <= len(readValue); i++ {
		Rgstrs[i] = readValue[i]
		fmt.Printf("\tRegister ID: %d \n", readValue[i].Rid)
		for j := 1; j <= len(readValue[i].RGroup); j++ {
			fmt.Printf("\t\tGroup ID: %d \n", readValue[i].RGroup[j].Gid)
			for k := 1; k <= len(readValue[i].RGroup[j].Groups); k++ {
				fmt.Printf("\t\t\tMember ID : %d \n", readValue[i].RGroup[j].Groups[k].Pid)
				fmt.Printf("\t\t\tMember Name: %s \n", readValue[i].RGroup[j].Groups[k].Name)
				fmt.Printf("\t\t\tMember Age : %d \n", readValue[i].RGroup[j].Groups[k].Age)
				fmt.Printf("\t\t\tPhone Number : %s \n", readValue[i].RGroup[j].Groups[k].PhoneNo)
				fmt.Printf("\t\t\tMagic Code : %d \n", readValue[i].RGroup[j].Groups[k].MagicCode)
			}
		}
	}
	defer fmt.Println("Reading completed the file is closing")
	defer jsonFile.Close()
}

// Function to write and save the json values
func writeJsonValues(Rgstrs map[int]Registery, fileName string) {
	println("Json file writer iss saving your entries")
	mapB, err := json.MarshalIndent(Rgstrs, "", " ")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(mapB))
		ioutil.WriteFile(fileName, mapB, 0644)
		fmt.Println("------------------- File saved successfully------------")
	}
}

// Function to create a new student
func createMagicCode(phoneNumber string, Pid int) int {
	code := 0
	codeCount := 0
	if phoneNumber == "" {
		Rid, Gid, Pid = getstudentGroup()
		phoneNumber = Rgstrs[Rid].RGroup[Gid].Groups[Pid].PhoneNo
	}
	for runeCheck := 0; runeCheck < 10; runeCheck++ {
		count := strings.Count(phoneNumber, strconv.Itoa(runeCheck))
		/*

		   --------------------------Magic code iterative value checker------------------------
		   Printf("Found %s in %d times \n", strconv.Itoa(runeCheck), count)

		*/
		if count >= 2 {
			codeCount = (codeCount + runeCheck*Pid)
			code = code + (count * codeCount)
		} else if count == 0 {
			codeCount = (codeCount + runeCheck) + Pid
		}
	}
	return code % 1000
}

func viewMagicCode() {
	_, gok := Rgstrs[Rid].RGroup[Gid].Groups[Pid]
	if gok {
		fmt.Printf("/n/t Your magic code is : %d", Rgstrs[Rid].RGroup[Gid].Groups[Pid].MagicCode)
	} else {
		getstudentGroup()
	}
}

func getstudentGroup() (Rid int, Gid int, Pid int) {
	var checkFlag bool
	for {
		fmt.Printf("\n\nEnter your register ID : \t")
		Rid = readInt()
		_, rok := Rgstrs[Rid]
		if rok {
			checkFlag = true
		} else {
			fmt.Printf("\n ID doen't exist \t")
			checkFlag = false
			return
		}
		fmt.Printf("\n\nEnter your group ID : \t")
		Gid = readInt()
		_, rgok := Rgstrs[Rid].RGroup[Gid]
		if rgok || checkFlag {
			checkFlag = true
			return
		} else {
			fmt.Printf("\n ID doen't exist \t")
			checkFlag = false
		}
		fmt.Printf("\n\nEnter your student ID : \t")
		Pid = readInt()
		_, gok := Rgstrs[Rid].RGroup[Gid].Groups[Pid]
		if gok || checkFlag {
			checkFlag = true
			return Rid, Gid, Pid
		} else {
			fmt.Printf("\n ID doen't exist \t")
			checkFlag = false
			return
		}
	}
}

func readInt() (optionStoA int) {
	var option string
	fmt.Scanf("%s", &option)
	optionStoA, err := strconv.Atoi(option)
	if err != nil {
		fmt.Printf("\nYou have not entered a valid number.\n")
		return 0
	} else {
		return optionStoA
	}
}
