package model

import (
	"time"
)

type Employee struct {
	PersistenceCommons
	Name        string    `json:"name,omitempty"`
	Age         int       `json:"age,omitempty"`
	JoiningDate time.Time `json:"joiningDate,omitempty"`
	Department  string    `json:"department,omitempty"`
}
