package main

import (
	. "fmt"
	"strconv"
	"strings"
)

// -------------------------------------- defining global variables here------------------------------

var studentList = make(map[int]student)

type student struct {
	name        string
	phoneNumber string
	magicCode   int
}

var studentID int

// ----------------------------------- Main function -------------------------
func main() {
	for {
		Println("Welcome to student register")
		Println(`
		1: Create student
		2: View Student details
		3: Delete student
		4: Get magic code
		5: Find student ID
		0: Exit
		`)
		var option int
		Scanf("%d", &option)

		switch option {
		case 0:
			return
		case 1:
			createStudent()
		case 2:
			viewStudent()
		case 3:
			removeStudent()
		case 4:
			Printf("\n Enter your Contact number : ")
			var contactNo string
			var idNo int
			Scanf("%s", &contactNo)
			Printf("\n Enter your student ID : ")
			Scanf("%d", &idNo)
			Printf("\n Magic code is : \t %d \n", magicCode(contactNo, idNo))
		default:
			Printf("invalid option : %d \n", option)
		}
	}
}

// -------------------------- create strudent function ---------------------
func createStudent() {
	var name, confirmation, phoneNumber string

	Println("Fill up the details to create the student")
	Printf("Enter your name : ")
	Scanf("%s", &name)
	Printf("\nEnter your phone number : ")
	Scanf("%s", &phoneNumber)

	Printf(`We are creating your student id with follwing details:
		Name : %s
		Phone Number : %s
		
		For confirmation press y :
		`, name, phoneNumber)
	Scanf("%s", &confirmation)
	switch confirmation {
	case "y":
		studentID++
		newstudent := student{
			name:        name,
			phoneNumber: phoneNumber,
			magicCode:   magicCode(phoneNumber, studentID),
		}
		studentList[studentID] = newstudent
		Printf("New student ID of %d is registered \n \t with magic code : %d \n\n", studentID, magicCode(phoneNumber, studentID))
	default:
		Println("Aborting the student creation...")
	}
}

// -------------------------- view strudent function ---------------------
func viewStudent() {
	var studentID int
	Printf("Enter the student id to view : ")
	Scanf("%d", &studentID)
	s, ok := studentList[studentID]
	if ok {
		Printf(`
	Name : %s 
	Phone no. : %s

	`, s.name, s.phoneNumber)
	} else {
		Println("No student id found")
	}
}

// --------------------------------Delete Student-------------------------
func removeStudent() {
	studentID := 0
	magicCode := 0
	confirmation := "n"
	Printf("\nEnter your student id : \t")
	Scanf("%d", &studentID)
	Printf("\nEnter the magic code : \t")
	Scanf("%d", &magicCode)

	s, ok := studentList[studentID]
	if ok && s.magicCode == magicCode {
		Printf(`
		The student with details given below is to be removed :
	Student ID : %d
	Name : %s 
	Phone no. : %s

	`, studentID, s.name, s.phoneNumber)

		Print("Please confirm (y) : \t")
		Scanf("%s", &confirmation)
		if confirmation == "y" {
			delete(studentList, studentID)
			Println("Student unregistered successfully")
		} else {
			Println("Student deregisteration process is aborted")
		}
	} else {
		Println("No student id found")
	}
}

// --------------------------------generate  magic code-------------------------
func magicCode(phoneNumber string, studentID int) int {
	code := 0
	codeCount := 0
	for runeCheck := 0; runeCheck < 10; runeCheck++ {
		count := strings.Count(phoneNumber, strconv.Itoa(runeCheck))
		/*

		   --------------------------Magic code iterative value checker------------------------
		   Printf("Found %s in %d times \n", strconv.Itoa(runeCheck), count)

		*/
		if count >= 2 {
			codeCount = (codeCount + runeCheck*studentID)
			code = code + (count * codeCount)
		} else if count == 0 {
			codeCount = (codeCount + runeCheck) + studentID
		}
	}
	return code % 1000
}
