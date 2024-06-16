package main

import (
	"encoding/json"
	"fmt"
)

// JwtData 구조체 정의
type JwtData struct {
	AuthTokenCD   string `json:"authTokenCD"`
	CallServiceCD string `json:"callServiceCD"`
	//Permissions   map[string]map[string][]string `json:"permissions"`
	Permissions map[string][]string `json:"permissions"`
}

// Response 구조체 정의
type Response struct {
	Jwt JwtData `json:"jwt"`
}

func main() {
	//jsonStr := `{"jwt":{"authTokenCD":"01","callServiceCD":"02","permissions":{"*":["*"]}}}`
	//jsonStr := `{"jwt":{"authTokenCD":"01","callServiceCD":"02","permissions":{"/jwt/verify":["POST","GET"],"/contract":["*"]}}}`
	jsonStr := `{"jwt":{"authTokenCD":"01","callServiceCD":"02","permissions":{}}}`

	response := &Response{}
	if err := json.Unmarshal([]byte(jsonStr), response); err != nil {
		panic(err)
	}
	jwtData := response.Jwt
	for k, v := range jwtData.Permissions {
		fmt.Println(k)
		fmt.Println(v)
	}
}
