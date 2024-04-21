package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type RolePolicy struct {
	PolicyName     string `json:"PolicyName"`
	PolicyDocument string `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version string
}

func main() {
	rp := RolePolicy{"policydoc", "policyname"}
	fmt.Println(rp)

	jsonFile, err := os.Open("input/mock.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &rp)

	fmt.Println("opened json" + rp.PolicyName + rp.PolicyDocument)
	defer jsonFile.Close()
}
