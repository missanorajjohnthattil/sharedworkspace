package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Pid  int
	Name string
}

type Group struct {
	Gid    int
	Groups map[int]Person
}

type Registery struct {
	Rid    int
	RGroup map[int]Group
}

var prsns = make(map[int]Person)

var Grps = make(map[int]Group)

var Rgstrs = make(map[int]Registery)

func main() {
	prsns[1] = Person{Pid: 1, Name: "Missan"}
	prsns[2] = Person{Pid: 2, Name: "Midhun"}
	prsns[3] = Person{Pid: 3, Name: "Reema"}
	Grps[1] = Group{Gid: 1, Groups: prsns}
	Grps[2] = Group{Gid: 2, Groups: prsns}
	Rgstrs[1] = Registery{Rid: 1, RGroup: Grps}

	Grps[1].Groups[3] = Person{Pid: 3, Name: "Johnson"}

	fmt.Printf("\n\t\t Grps[1].Groups[2]  \t : %v \n", Grps[1].Groups[2])

	fmt.Printf(" Rgstrs : \t %v \n ", Grps)
	fmt.Printf(" \t Length : \t %v \n ", len(Grps))
	fmt.Printf("\t \t Rgstrs[1] : \t %v \n ", Grps[1])
	fmt.Printf("\t \t \t Rgstrs[1].Group : \t %v \n ", Grps[1].Groups)
	fmt.Printf("\t \t \t \t Length : \t %v \n ", len(Grps[1].Groups))
	fmt.Printf("\t \t \t \t \t Rgstrs[1].Group[1] : \t %v \n ", Grps[1].Groups[1])
	fmt.Printf("\t \t \t \t \t \t Rgstrs[1].Group[2] : \t %v \n ", Grps[1].Groups[2])
	fmt.Printf("\t \t \t \t \t \t \t Rgstrs[1].Group[1].Name : \t %v \n ", Grps[1].Groups[1].Name)
	fmt.Println("")
	mapB, err := json.Marshal(Rgstrs)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(mapB))
	}
}
