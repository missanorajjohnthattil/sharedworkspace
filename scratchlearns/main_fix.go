package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"PersonName"`
	Age  int    `json:"Age"`
}

// ---------------------------------------------- Group Registery---------------

func (grp GroupRegistry) addPerson() {
	var groupId int
	fmt.Printf("Enter the group id : ")
	fmt.Scanf("%d", &groupId)
	s, ok := Gr[groupId]
	if ok {
		fmt.Printf("\n\n\tPlease fill up the below details to add new member to %s :", s.GName)
		mId := len(s.GMembers) + 1
		fmt.Printf("\n\n\t\t%v -\tlength : %v", s.GMembers, mId)
		ss, _ := s.GMembers[mId]
		fmt.Printf("\n\t Enter your Name : ")
		fmt.Scanf("%s", &ss.Name)
		fmt.Printf("\n\t Enter your age : ")
		fmt.Scanf("%d", &ss.Age)
		fmt.Printf("\t %s of age %d is added to %s with id %d \n", ss.Name, ss.Age, s.GName, mId)
	} else {
		fmt.Println("The group ID you entered doesn't exist")
	}
}

// ---------------------------------------------- Group Registery---------------

type GroupRegistry struct {
	GName    string         `json:"groupName"`
	GMembers map[int]Person `json:"groupMembers"`
	// gAdmin   map[int]person // will have to create it later
}

var Gr = make(map[int]GroupRegistry)

var grp GroupRegistry

func (grp GroupRegistry) createGroup() {
	var gId int
	gId = len(Gr)
	fmt.Printf("\n\nEnter the new group name : ")
	fmt.Scanf("%s", &grp.GName)
	gId += 1
	Gr[gId] = GroupRegistry{GName: grp.GName}
	fmt.Printf("\n\n\t New group %s is registered with group id %d \n\n", grp.GName, gId)
	fmt.Printf("\n\n\n\t %v", Gr)
}

func (grp GroupRegistry) viewGroup() {
	var grpId int
	fmt.Printf("\n\n Enter your group id : ")
	fmt.Scanf("%d", &grpId)
	s, ok := Gr[grpId]
	if ok {
		fmt.Printf("The %s group found", s.GName)
		fmt.Printf("\n\n\t %v", s.GMembers)
	} else {
		fmt.Println("The group ID you entered doesn't exist")
	}
}

// ---------------------------------------------- Json codes ---------------

func (grp GroupRegistry) saveJson() {
	fmt.Printf("\n\n\n\t %v", Gr)
	mapB, _ := json.Marshal(Gr)
	fmt.Println(string(mapB))

	fmt.Printf("\n\n The below mentioned entiries are to be written :  \n\t %v ", Gr)
}

// ---------------------------------------------- Main Function ---------------

func main() {
	fmt.Printf("\n\t\tWelcome registry manager\n\n")
	for {
		var option int
		fmt.Printf(`
		Action No -- Option detail
		1 ---------- Add a member to group
		2 ---------- View members in group
		3 ---------- Add new group
		4 ---------- View group details
		10 --------- Get the Magic code
		15 --------- Delete a member from the group
		16 --------- Delete the group
		20 --------- Save as json
		0 ---------- Exit from the function

		Select your action
		`)
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			grp.addPerson()
		case 3:
			grp.createGroup()
		case 4:
			grp.viewGroup()

		case 20:
			grp.saveJson()
		case 0:
			println("Happy to see you later")
			return
		default:
			println("\n\n You have entered a wrong entry \n\n")
		}
	}
}
