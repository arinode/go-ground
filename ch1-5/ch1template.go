package main

import "fmt"

/*..
  go list : list all packages or modules
  go run : compile and run
  go run main.go
  go build main.go it compiles the code and generates the compiled file binary data
  go version
.*/

func main() {
	fmt.Println("starting server...")

	// create a new variable, it defaults to 0
	// (the "zero value" for ints)
	var mySkillIssues int
	// overwrite the zero value with 42
	mySkillIssues = 42
	fmt.Println("Your SMS sending limit is", mySkillIssues)

	var username string
	username = "eddie_cabot"

	var isAdmin bool
	isAdmin = true

	var permissions int
	permissions = 0x1F

	var costPerSMS float64
	costPerSMS = 0.05

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("permissions:", permissions)
	fmt.Println("costPerSMS:", costPerSMS)

	//GOATed variable declaration := type inference
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)

	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

}
