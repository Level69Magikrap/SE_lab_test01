package entity

import (
	"time"
	"gorm.io/gorm"
	"github.com/Level69Magikrap/se_lab_test01"
)
type User struct{
	gorm.Model
	Name string `valid:"required~need name"`
	Age int `valid:"required~need age"`
	DOB time.Time `valid:"required~need time"`
}