package conversor

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"username"`
	Online bool   `json:"active"`
}

func ConverterUserToJSON(user User) (string, error) {
	e, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(string(e))
	return string(e), nil
}

func ConverterUserFromJSON(user string) {
	model := &User{}

	err := json.Unmarshal([]byte(user), model)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("Name: %s, Online: %t", model.Name, model.Online)
}
