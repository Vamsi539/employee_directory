package main

import "fmt"

func main(){
	var dir = directory{employees: []employee{}}
	 e := employee{
		name:   "Vamsi",
		id:     1,
		phone:  "12345678",
		age:    25,
		salary: 30000,
		email:  "Vamsi.gmail.com",
	}
	dir.addEmployee(e)
	dir.getid("Vamsi")
	dir.updateSalary(1,15000)
	f :=employee{
		"Dheeraj",
		2,
		"987654321",
		25,
		5000,
		"Dheeraj@gmail.com",
	}
	dir.addEmployee(f)
	dir.getinfo(e)
	fmt.Println(dir)

}