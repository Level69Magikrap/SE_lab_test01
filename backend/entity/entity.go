package entity

import (
	"time"
	"gorm.io/gorm"
)
type User struct{
	gorm.Model
	Name string `valid:"required~need name"`
	Age int `valid:"required~need age"`
	DOB time.Time `valid:"required~need time"`
}