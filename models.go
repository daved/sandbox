package main

import "github.com/jinzhu/gorm"

type customer struct {
	gorm.Model
	Name   string
	Orders []order
}

type order struct {
	gorm.Model
	Note string
}
