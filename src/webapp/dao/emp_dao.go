package dao

import (
	"errors"
	"fmt"
	"webapp/model"
)

var (
	employees = map[int64]*model.Employee{
		100: {
			Empid:    101,
			EmpName:  "Paras",
			EmpEmail: "paras@gmail.com",
		},
	}
)

func GetEmployee(empId int64) (*model.Employee, error) {
	if employee := employees[empId]; employee != nil {
		return employee, nil
	}
	return nil, errors.New(fmt.Sprintf("Employee %v not found", empId))
}
