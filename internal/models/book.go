package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title   string
	Authors []string
	ISBN    string
	Price   int
	Pages   int
	Created time.Time
}
