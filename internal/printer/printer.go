package printer

import "fmt"

func PrintMessage(head, message string) {
	formatedMessage := head + ": " + message
	fmt.Println(formatedMessage)
}
