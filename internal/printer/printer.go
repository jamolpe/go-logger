package printer

import (
	"fmt"
	"gologger/pkg/models"
)

// PrintMessage prints the log with corresponding format
func PrintMessage(log models.LogModel) {
	formatedMessage := log.Time.String() + " " + log.FormatedHead + " : " + log.Message
	fmt.Println(formatedMessage)
}
