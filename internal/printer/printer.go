package printer

import (
	"fmt"
	"gologger/pkg/models"
)

func PrintMessage(log models.LogModel) {
	formatedMessage := log.Time.String() + " " + log.Head + " : " + log.Message
	fmt.Println(formatedMessage)
}
