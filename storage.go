package main

type BirthdayStorage struct {
	file      string
	birthdays []*Birthday
}

type Birthday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
