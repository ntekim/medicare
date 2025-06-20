package helpers

type Response struct {
	Code	int 	`json:"code" dc:"Response Code"`
	Message string      `json:"message" dc:"Response message"`
	Data    interface{} `json:"data" dc:"Response payload"`
}
