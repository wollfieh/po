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
	message := pushover.NewMessageWithTitle(readbody(), strings.Join(os.Args[1:], " "))

	// Send the message to the recipient
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	log.Println(response)
}

// readbody reads the body from the standard input and returns it as a string.
func readbody() string {
	res, _ := io.ReadAll(os.Stdin)
	return string(res)
}
