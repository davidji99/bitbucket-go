package command

import (
	"fmt"
	"strings"
)

func GetFlagCliString(cmdFlag string) string {
	return fmt.Sprintf("%s, %s", cmdFlag, GetCmdFlagShort(cmdFlag))
}

func GetCmdFlagShort(fullName string) string {
	shortName := ""
	fullNameSplit := strings.Split(fullName, "-")

	if len(fullNameSplit) == 2 {
		for _, word := range fullNameSplit {
			shortName += string(word[0])
		}
	} else if len(fullNameSplit) == 1 {
		shortName += string(fullNameSplit[0][0])
	} else {
		shortName = ""
	}

	return shortName
}
