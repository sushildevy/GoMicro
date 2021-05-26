package model

type Employee struct {
	Empid    int64  `json:"emp_id"`
	EmpName  string `json:"emp_name"`
	EmpEmail string `json:"emp_email"`
}
