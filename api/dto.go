package main

type RegisterBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}
