package fixer

import "fmt"

type ResponseBody struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}
type ResponseBodySymbols struct {
	Success bool              `json:"success"`
	Symbols map[string]string `json:"symbols"`
}

type ResponseBodyConversation struct {
	Success    bool    `json:"success"`
	Query      Query   `json:"query"`
	Info       Info    `json:"info"`
	Historical string  `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}
type Query struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
type Info struct {
	Timestamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
}

func (r ResponseBody) GetResponseInfo() string {
	resultString := fmt.Sprintf("Success: %t\nTimestamp: %d\nBase: %s\nDate: %s\n", r.Success, r.Timestamp, r.Base, r.Date)
	for key, value := range r.Rates {
		resultString += fmt.Sprintf("		%s: %f\n", key, value)
	}
	return resultString
}

func (rc ResponseBodyConversation) GetResponseConversationInfo() string {
	resultString := fmt.Sprintf("Success: %t\nFrom: %s\nTo: %s\nAmount: %f\nHistorical: %s\nDate: %s\nResult: %f\n",
		rc.Success, rc.Query.From, rc.Query.To, rc.Query.Amount, rc.Historical, rc.Date, rc.Result)
	return resultString
}
func (rs ResponseBodySymbols) GetAllSymbolsInfo() string {
	resultString := fmt.Sprintf("Success: %t\n", rs.Success)
	for key, value := range rs.Symbols {
		resultString += fmt.Sprintf("		%s: %s\n", key, value)
	}
	return resultString

}
