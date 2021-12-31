package model

type TradingRecord struct {
	Id int 					`json:"id"`
	Username string 		`json:"username"`
	Forusername string 		`json:"forusername"`
	Time string 			`json:"time"`
	TranserRMB int 			`json:"transerrmb"`
	AccountBalanceNow int 	`json:"accountbalancenow"`
	Txt string 				`json:"txt"`
}