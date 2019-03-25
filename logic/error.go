package logic

import (
	"log"
)

// SendError is a helper function to write a formatted error.
func SendError(err error, issue string) {
	log.Printf("issue: %v.\nerror: %v", issue, err)
}
