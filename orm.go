package main

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Text string `json:"text"`
}
