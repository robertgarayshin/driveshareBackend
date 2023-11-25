package models

type Payload struct {
	Id int `json:"id"`
	//Role string `json:"role"`
	TokenType string `json:"token_type"`
	AuthToken string `json:"auth_token"`
}
