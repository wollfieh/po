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
	userKey := "akwjcknmaxrtx4sven9x2oy4974pm7"
	apiToken := "u9ps3pf4e72o39dksnhwcztd287x3b"
	// Create a new pushover app with a token
	app := pushover.New(userKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(apiToken)

	// Create the message to send
	message := pushover.NewMessageWithTitle(limitBody(readBody(), 1024), strings.Join(os.Args[1:], " "))

	// Send the message to the recipient
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
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
	appended := fmt.Sprintf("Message truncated to %d characters", length)
	remainder := length - len(appended)
	if len(body) > length {
		return body[:remainder] + appended
	}
	return body
}
