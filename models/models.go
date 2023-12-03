package models

import "gorm.io/gorm"

// custom struct Fact
// We will use json struct tags to describe the corresponding
// lowercase json keys for each of our fields
// we will also add GORM struct tags to our Fact fields
// this tags will represent as TEXT columns in the database
// set the defaul value to NULL so we can manage error when
// field are not created.
type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;defult:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}
