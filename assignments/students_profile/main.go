package main

import . "fmt"

// -------------------------------------- defining global variables here------------------------------

var studentList = make(map[int]student)

type student struct {
	name        string
	phoneNumber string
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
		newstudent := student{
			name:        name,
			phoneNumber: phoneNumber,
		}
		studentID++
		studentList[studentID] = newstudent
		Printf("New student ID of %d is registered \n", studentID)
	default:
		Println("Aborting the student creation...")
	}
}

// -------------------------- view strudent function ---------------------
func viewStudent() {
	var studentID int
	Printf("Enter the student id to view : ")
	Scanf("%d", studentID)
	Printf("%+v", studentList[studentID])
}
