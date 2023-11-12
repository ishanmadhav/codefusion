package api

import "gorm.io/gorm"

type Code struct {
	gorm.Model
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
	Output   string `json:"output"`
}
