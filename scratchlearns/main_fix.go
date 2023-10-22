package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"PersonName"`
	Age  int    `json:"Age"`
}

type GroupRegistry struct {
	GName    string         `json:"groupName"`
	GMembers map[int]Person `json:"groupMembers"`
	// gAdmin   map[int]person // will have to create it later
}

var Prsn Person

var prsns map[int]Person

var Grp GroupRegistry

var grps map[int]GroupRegistry

var Gr = make(map[int]GroupRegistry)

// ---------------------------------------------- Group Registery---------------

func (grp *GroupRegistry) addPerson() {
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
		fmt.Scanf("%s", &Prsn.Name)
		fmt.Printf("\n\t Enter your age : ")
		fmt.Scanf("%d", &Prsn.Age)
		fmt.Printf("\t %s of age %d is added to %s with id %d \n", Prsn.Name, Prsn.Age, s.GName, mId)
		fmt.Println("Type of Gr[groupId].GMembers[mId] is ", reflect.TypeOf(Gr[groupId].GMembers[mId]))
		fmt.Printf("Type of Gr[groupId].GMembers[mId] has value %v \n ", Gr[groupId].GMembers[mId])
		fmt.Println("ss is ", reflect.TypeOf(ss))
		fmt.Printf("ss has value %v \n", ss)
		fmt.Println("s is ", reflect.TypeOf(s))
		fmt.Printf("s has value %v \n", s)
		Gr[groupId].GMembers[mId] = Prsn
	} else {
		fmt.Println("The group ID you entered doesn't exist")
	}
}

// ---------------------------------------------- Group Registery---------------

func (Grp *GroupRegistry) createGroup() {
	var gId int
	gId = len(Gr)
	fmt.Printf("\n\nEnter the new group name : ")
	fmt.Scanf("%s", &Grp.GName)
	gId += 1
	fmt.Printf("\n\n\t New group %s is registered with group id %d \n\n", Grp.GName, gId)
	fmt.Printf("\n\n\n\t %v", Gr)
}

func (Grp *GroupRegistry) viewGroup() {
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

func (Grp *GroupRegistry) saveJson() {
	fmt.Printf("\n\n\n\t %v \n", Gr)
	mapB, err := json.Marshal(Gr)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(mapB))
	}

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
		2 ---------- !View members in group
		3 ---------- Add new group
		4 ---------- View group details
		10 --------- !Get the Magic code
		15 --------- !Delete a member from the group
		16 --------- !Delete the group
		20 --------- Save as json
		0 ---------- Exit from the function

		Select your action
		`)
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			Grp.addPerson()
		case 3:
			Grp.createGroup()
		case 4:
			Grp.viewGroup()

		case 20:
			Grp.saveJson()
		case 0:
			println("Happy to see you later")
			return
		default:
			println("\n\n You have entered a wrong entry \n\n")
		}
	}
}
