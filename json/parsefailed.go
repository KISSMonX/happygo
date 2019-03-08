package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Timestamp    int64  `json:"ts"`
	ServiceID    int32  `json:"sid"`
	SubserivceID int32  `json:"ssid"`
	Body         string `json:"body"`
}

func main() {
	var it Item
	str := `{"ts":1548040280,"sid":3,"ssid":1,"body":"{\"msg\":\"Do you know how to  leave \\"Feedback\\"? \",\"msg_type\":0,\"name\":\"Miguiel G\",\"role\":20,\"sender\":\"72d29dd096ba4268be8af89fd2ed7663\"}"}`
	err := json.Unmarshal([]byte(str), &it)
	fmt.Println(err)
	fmt.Println(it)
}
