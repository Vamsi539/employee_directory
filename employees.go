package main

type employee struct {
	id int
	name string
	age int
	phone string
	salary int
	email string
}

type directory struct{
	employees []employee
}

func (d *directory) addEmployee(emp employee) {
	d.employees = append(d.employees, emp)
}

func (d *directory) deleteEmployee(id int) {
	d.deleteEmployee(id)
}


func (d *directory) updateSalary(id int,newSalary int) {
	e := d.employees[id]
	e.salary = newSalary
	d.employees[id]=e

	//emp := d.employees[id]
	//emp.salary = newSalary
	//d.employees[id] = emp
}
func (d *directory) updateName(id int,newName string){
	e := d.employees[id]
	e.name = newName
	d.employees[id]=e
}
func (d *directory) getid(Fname string)int {
	for _, e:= range d.employees {
		if e.name == Fname{
			return e.id
		}
	}
	return 0
}
func (d *directory) getname(id int)string{
	emp := d.employees[id]
	return emp.name
}

func (d *directory) getinfo(e employee)(string, int, string, int ,int, string){
	return e.name,e.id,e.phone,e.age,e.salary,e.email
}