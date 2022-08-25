package book //nama package dari folder

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{ID: "1", Title: "Lost My Mind", Author: "Peanut", Quantity: 1},
	{ID: "2", Title: "Close Your Eyes", Author: "Stund", Quantity: 12},
	{ID: "3", Title: "I Don't Know", Author: "Crult", Quantity: 5},
}