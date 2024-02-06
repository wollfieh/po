package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gregdel/pushover"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("%s <messsage>\n", os.Args[0])

	}
	credentials := getEnv()
	// Create a new pushover app with a token
	app := pushover.New(credentials["userKey"])

	// Create a new recipient
	recipient := pushover.NewRecipient(credentials["apiToken"])

	// message comes from stding, title is the arguments passed
	message := pushover.NewMessageWithTitle(limitBody(readBody(), 1024), strings.Join(os.Args[1:], " "))

	// Send the message to the recipient
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response if you want
	log.Println(response)
}

// readBody reads the body from the standard input and returns it as a string.
func readBody() string {
	res, _ := io.ReadAll(os.Stdin)
	if len(res) == 0 {
		return "No body"
	}
	return string(res)
}

// func limitBody truncates the body to the given length. it takes a string and an int as arguments and returns a string.
func limitBody(body string, length int) string {
	appended := fmt.Sprintf("\n\nMessage truncated to %d characters", length)
	remainder := length - len(appended)
	if len(body) > length {
		return body[:remainder] + appended
	}
	return body
}

// get environment variables PO_userKey and PO_apiToken
func getEnv() map[string]string {
	return map[string]string{"userKey": os.Getenv("PO_userKey"), "apiToken": os.Getenv("PO_apiToken")}
}
