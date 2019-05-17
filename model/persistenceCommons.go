package model

import (
	"time"
)

type PersistenceCommons struct {
	Id               string    `json:"id,omitempty" gorm:"primary_key, size:36"`
	CreationTime     time.Time `json:"creationTime,omitempty"`
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
	CreatedBy        string    `json:"createdBy,omitempty" gorm:"size:100"`
	LastModifiedBy   string    `json:"lastModifiedBy,omitempty" gorm:"size:100"`
}
