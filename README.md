# AWS::IAM::Role Verifier
Go module for verifying whether an AWS::IAM::Role Policy JSON file doesn't contain a wildcard (*) in the Resource field.
## Installation
This module can be used both as a standalone CLI application and imported as a module into your Go project.
### CLI
Just clone this repository and make sure you have Go installed.
### Go module
Add this module as a dependency:
```bash
go get github.com/mrcxmrj/aws-iam-role-verifier/roleverifier
```
## Usage
### CLI
```bash
go run cmd/main.go <path_to_your_json_file>
```
Running this command will print false if "*" is present in one of the Resource fields and true otherwise.
### Go module
Import `"github.com/mrcxmrj/aws-iam-role-verifier/roleverifier"` into your code. This package exposes `Verify(path string) (bool error)` function, that works analogously to the CLI.\
Example usage:
```go
package main

import (
	"fmt"
	"os"
	"github.com/mrcxmrj/aws-iam-role-verifier/roleverifier"
)

func main() {
	result, err := roleverifier.Verify("input.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
```
Appropriate types for working with AWS::IAM::Role JSONs are also available:
```go
type RolePolicy struct {
	PolicyName     string
	PolicyDocument PolicyDocument
}

type PolicyDocument struct {
	Version         string
	Id              string
	Statement       []Statement
	StatementSingle Statement
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
```
> [!NOTE]  
> Since Statement field can contain either an array of Statement JSONs or just a singular JSON, they get stored after unmarshalling in either Statement or Statement Single
fields, respectively.
## Testing
To add custom tests for the Verifier function - add a json file you want to test to `roleverifier/test_input` directory, then add a new element to the `tests` table
in `roleverifier/verifier_test.go`:
```go
tests := map[string]struct {
    path     string
    expected bool
}{
    ...
    "new test name": {
        path:     "test_input/<new_test_filename>.json",
        expected: <expected_value>,
    },
}
```
