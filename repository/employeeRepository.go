package repository

import (
	"com.kaushal/rai/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type EmployeeRepository interface {
	Add(employee model.Employee) error
	Get(employeeId string) (model.Employee, error)
	Delete(employeeId string) error
	Update(employee model.Employee) error
	List(start int, limit int) ([]model.Employee, error)
}

type PostgresRepository struct {
	Db *gorm.DB
}

func (r *PostgresRepository) Add(employee model.Employee) error {
	// Create tables in public schema
	// Need to think on selecting specific Schema in DB and not public
	r.Db.AutoMigrate(&model.Employee{})
	r.Db.Create(&employee)

	return nil
}

func (r *PostgresRepository) Get(employeeId string) (model.Employee, error) {
	var employee model.Employee
	r.Db.First(&employee, "id = ?", employeeId)

	return employee, nil
}

func (r *PostgresRepository) Update(employee model.Employee) error {
	r.Db.Save(&employee)

	return nil
}

func (r *PostgresRepository) Delete(employee *model.Employee) error {
	r.Db.Delete(&employee)

	return nil
}

func (r *PostgresRepository) List(start, limit int) (*[]model.Employee, error) {
	r.Db.AutoMigrate(&model.Employee{})

	var employees []model.Employee
	r.Db.Offset(start).Limit(limit).Find(&employees)

	return &employees, nil
}
