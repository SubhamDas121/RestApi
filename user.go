package main

import "gorm.io/gorm"

type USER struct {
	gorm.Model

	NAME        string `json:"name"`
	DOB         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
