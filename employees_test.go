package main

import "testing"

func Test_directory_updateSalary(t *testing.T) {
	type fields struct {
		employees []employee
	}
	type args struct {
		id        int
		newSalary int
	}
	var tests = []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test case 1",fields: fields{},
			, args: args{1, 20000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &employee{
				id:        tt.fields.id,
				name:		tt.fields.name,
				age:		tt.fields.age,
				phone:     tt.fields.phone,
				email:     tt.fields.email,
				salary:    tt.fields.salary,
			}
			e.updateSalary(tt.args.newSalary)
			//assert.Equal(t, e.salary, tt.args.newSalary)
		})

		//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		d := &directory{
	//			employees: tt.fields.employees,
	//		}
	//	})
	}
}