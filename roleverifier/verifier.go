package roleverifier

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
	Version         string
	Id              string
	Statement       []Statement
	StatementSingle Statement
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

func (pd *PolicyDocument) UnmarshalJSON(data []byte) error {
	var temp struct {
		Version   string
		Id        string
		Statement json.RawMessage
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if err := json.Unmarshal(temp.Statement, &pd.Statement); err == nil {
		return nil
	}
	if err := json.Unmarshal(temp.Statement, &pd.StatementSingle); err == nil {
		return nil
	}

	return fmt.Errorf("Unable to unmarshal Statement")
}

func Verify(path string) (bool, error) {
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

	if result.PolicyDocument.StatementSingle.Resource == "*" {
		return false, nil
	}

	for _, statement := range result.PolicyDocument.Statement {
		if statement.Resource == "*" {
			return false, nil
		}
	}
	return true, nil
}
