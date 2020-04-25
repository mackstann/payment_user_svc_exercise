package models

type User struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Address    Address `json:"address"`
}

type Address struct {
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2"`
	City        string `json:"city"`
	Subdivision string `json:"subdivision"`
	PostalCode  string `json:"postal_code"`
}
