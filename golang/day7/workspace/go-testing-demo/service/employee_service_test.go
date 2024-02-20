package service

import (
	"errors"
	"myapp/model"
	"testing"
)

// ---- mock implementation of EmployeeDao interface
type MockEmployeeDao struct {
}

func (dao MockEmployeeDao) FindById(id int) (model.Employee, error) {
	if id <= 0 {
		return model.Employee{}, errors.New("invalid id")
	}
	if id == 1 {
		return model.Employee{Id: 1, Name: "John Doe", Salary: 50000}, nil
	}

	return model.Employee{}, errors.New("employee not found")
}

func (dao MockEmployeeDao) Save(employee model.Employee) (int, error) {
	if employee.Name == "John" {
		return 7654, nil
	}
	if employee.Name == "000" {
		return 0, errors.New("db connect failed")
	}
	return 0, nil
}

//---- end of mock implementation of EmployeeDao interface

func TestAddNewEmployee(t *testing.T) {
	var mockDao MockEmployeeDao
	service := EmployeeService{dao: mockDao}

	t.Run("valid employee added", func(t *testing.T) {
		emp := model.Employee{Name: "John", Salary: 44000}
		savedEmp := service.AddNewEmployee(emp)
		if savedEmp.Id != 7654 {
			t.Error("saved id is incorrect")
		}
	})

	t.Run("employee with blank name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		emp := model.Employee{Name: "    ", Salary: 44000}
		service.AddNewEmployee(emp)
	})
	t.Run("employee with too large name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		emp := model.Employee{Name: "vinodvinodvinodvinodvinodvinodvinodvinodvinodvinodvinodvinod", Salary: 44000}
		service.AddNewEmployee(emp)
	})
	t.Run("employee with low salary", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		emp := model.Employee{Name: "Kishore", Salary: 4000}
		service.AddNewEmployee(emp)
	})
	t.Run("employee with high salary", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		emp := model.Employee{Name: "Kishore", Salary: 500001}
		service.AddNewEmployee(emp)
	})
	t.Run("database error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		emp := model.Employee{Name: "000", Salary: 50000}
		service.AddNewEmployee(emp)
	})
}

func TestGetEmployee(t *testing.T) {
	var mockDao MockEmployeeDao
	service := EmployeeService{dao: mockDao}

	t.Run("invalid id", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("was expecting a panic; did not get one")
			}
		}()
		service.GetEmployee(0)
	})

	t.Run("employee not found", func(t *testing.T) {
		emp := service.GetEmployee(123)
		if emp != nil {
			t.Errorf("was expecting `nil` but got %v", emp)
		}
	})

	t.Run("existing employee", func(t *testing.T) {
		emp := service.GetEmployee(1)
		if emp == nil {
			t.Error("was not expecting `nil`")
		} else if emp.Name != "John Doe" || emp.Salary != 50000 {
			t.Errorf("wanted employee with name `John Doe` and salary `50000` but got `%v` and `%v`", emp.Name, emp.Salary)
		}
	})
}

// go test -v -cover ./service
