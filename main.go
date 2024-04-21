package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type RolePolicy struct {
	PolicyName     string
	PolicyDocument PolicyDocument
}

// NOTE: edge case: statement can be singular instead of an array
type PolicyDocument struct {
	Version   string
	Id        string
	Statement []Statement
}

type Statement struct {
	Sid          string
	Effect       string
	Principal    string
	NotPrincipal string
	Action       []string
	NotAction    string
	Resource     string
	NotResource  string
}

func main() {
	jsonFile, err := os.Open("input/statements.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result RolePolicy
	if err := json.Unmarshal(byteValue, &result); err != nil {
		fmt.Println(err)
	}

	for _, statement := range result.PolicyDocument.Statement {
		fmt.Println("policy resource: " + statement.Resource)
		if statement.Resource == "*" {
			fmt.Println(true)
		}
	}
}
