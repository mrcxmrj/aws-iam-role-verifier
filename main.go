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

// NOTE: edge case: Resource can be nil (when NotResource is being used)
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

func verifier(path string) (bool, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return true, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result RolePolicy
	if err := json.Unmarshal(byteValue, &result); err != nil {
		return true, err
	}

	for _, statement := range result.PolicyDocument.Statement {
		if statement.Resource == "*" {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	args := os.Args
	path := args[1]
	result, err := verifier(path)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
