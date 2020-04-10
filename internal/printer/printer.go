package printer

import (
	"fmt"

	"github.com/jamolpe/gologger/pkg/models"
)

// PrintMessage prints the log with corresponding format
func PrintMessage(log models.LogModel) {
	formatedMessage := log.Time.String() + " " + log.FormatedHead + " : " + log.Message
	fmt.Println(formatedMessage)
}
