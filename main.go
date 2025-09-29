package main

import (
	"fmt"
	"os"
)

func generateReadmeContent(name string) string {
	return `## Question
Add the question here.

### Example
` + "```" + `go
// Example input
// 5
// Example output
// 120
` + "```" + `

### Solution
- [Implementation in Go](./` + name + `.go)
- [Unit Test Cases](./` + name + `_test.go)
`
}

func generateGoFileContent(name string) string {
	return `package ` + name + `

import "fmt"

func main() {
	fmt.Println("Hello from ` + name + `!")
}
`
}

func generateTestFileContent(name string) string {
	return `package ` + name + `

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing ` + name + ` package")
}
`
}

func createChallangeDir(name string) error {
	fmt.Printf("Creating directory: %s\n", name)
	return os.Mkdir(name, 0755)
}

func createGoFile(dir, name string) error {
	filePath := fmt.Sprintf("%s/%s.go", dir, name)
	fileContent := generateGoFileContent(name)
	return os.WriteFile(filePath, []byte(fileContent), 0644)
}

func createReadmeFile(dir, name string) error {
	filePath := fmt.Sprintf("%s/README.md", dir)
	fileContent := generateReadmeContent(name)
	return os.WriteFile(filePath, []byte(fileContent), 0644)
}

func createTestFile(dir, name string) error {
	filePath := fmt.Sprintf("%s/%s_test.go", dir, name)
	fileContent := generateTestFileContent(name)
	return os.WriteFile(filePath, []byte(fileContent), 0644)
}

func main() {
	fmt.Println("Enter the challanege name without spaces to created a new directory for it: (e.g. fireworkschallange)")
	var challangeName string
	fmt.Scanln(&challangeName)
	err := createChallangeDir(challangeName)
	if err != nil {
		fmt.Println("Error creating challange directory:", err)
		return
	}
	fmt.Println("Challange directory created successfully. Creating the solution file template...")
	err = createGoFile(challangeName, challangeName)
	if err != nil {
		fmt.Println("Error creating Go file:", err)
		return
	}
	err = createReadmeFile(challangeName, challangeName)
	if err != nil {
		fmt.Println("Error creating README file:", err)
		return
	}
	err = createTestFile(challangeName, challangeName)
	if err != nil {
		fmt.Println("Error creating test file:", err)
		return
	}
	fmt.Println("All files created successfully.")
	fmt.Println("You can now implement your solution in", challangeName+"/"+challangeName+".go")
}