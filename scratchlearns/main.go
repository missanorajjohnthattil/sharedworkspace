package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string `json:"PersonName"`
	age  int    `json:"Age"`
}

// ---------------------------------------------- Group Registery---------------

func (grp groupRegistry) addPerson() {
	var groupId int
	fmt.Printf("Enter the group id : ")
	fmt.Scanf("%d", &groupId)
	s, ok := gr[groupId]
	if ok {
		fmt.Printf("\n\n\tPlease fill up the below details to add new member to %s :", s.gName)
		mId := len(s.gMembers) + 1
		fmt.Printf("\n\n\t\t%v -\tlength : %v", s.gMembers, mId)
		ss, _ := s.gMembers[mId]
		fmt.Printf("\n\t Enter your Name : ")
		fmt.Scanf("%s", &ss.name)
		fmt.Printf("\n\t Enter your age : ")
		fmt.Scanf("%d", &ss.age)
		fmt.Printf("\t %s of age %d is added to %s with id %d \n", ss.name, ss.age, s.gName, mId)
	} else {
		fmt.Println("The group ID you entered doesn't exist")
	}
}

// ---------------------------------------------- Group Registery---------------

type groupRegistry struct {
	gName    string         `json:"groupName"`
	gMembers map[int]person `json:"groupMembers"`
	// gAdmin   map[int]person // will have to create it later
}

var gr = make(map[int]groupRegistry)

var grp groupRegistry

func (grp groupRegistry) createGroup() {
	var gId int
	gId = len(gr)
	fmt.Printf("\n\nEnter the new group name : ")
	fmt.Scanf("%s", &grp.gName)
	gId += 1
	gr[gId] = groupRegistry{gName: grp.gName}
	fmt.Printf("\n\n\t New group %s is registered with group id %d \n\n", grp.gName, gId)
	fmt.Printf("\n\n\n\t %v", gr)
}

func (grp groupRegistry) viewGroup() {
	var grpId int
	fmt.Printf("\n\n Enter your group id : ")
	fmt.Scanf("%d", &grpId)
	s, ok := gr[grpId]
	if ok {
		fmt.Printf("The %s group found", s.gName)
		fmt.Printf("\n\n\t %v", s.gMembers)
	} else {
		fmt.Println("The group ID you entered doesn't exist")
	}
}

// ---------------------------------------------- Json codes ---------------

func (grp groupRegistry) saveJson() {
	fmt.Printf("\n\n\n\t %v", gr)
	mapB, _ := json.Marshal(gr)
	fmt.Println(string(mapB))

	fmt.Printf("\n\n The below mentioned entiries are to be written :  \n\t %v ", gr)
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
