package main

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Position int `json:"position"`
	Complete bool `json:"complete"`
}
