package models

//Data ...
type Data struct {
	Body   string `json:"body"`
	Title  string `json:"title"`
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
}

//Link ...
type Link struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

//Pagination ...
type Pagination struct {
	Links Link `json:"links"`
	Total int  `json:"total"`
	Pages int  `json:"pages"`
	Page  int  `json:"page"`
	Limit int  `json:"limit"`
}

//Meta ...
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

//Response ...
type Response struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}
