package main

import "strings"
import "fmt"

var Studentname = "default"
var Studentid = 1000

var Studentaddress = make(map[string]string)
var Phone_number = 9876543210

var Magiccode = luckcodegen(phone_number,studentid)

func main() {


}

func deletestudent(Studentid,Magiccode){

}


func getstudent(studentid) string{
	fmt.printf(" id: %d \n name: %s \n phone no: %d \n code: %d", studentid, studentname, phone_number, magiccode)
	for(int i = 1,address := range Studentaddress){
	fmt.println(i,")",address)
	}
}

func luckcodegen(phone_number, key := studentid) int{
	codearray := make([10]int,10,10)
	codearray = phone_number
	return codearray%key
}
	