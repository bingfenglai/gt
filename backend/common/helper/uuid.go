package helper

import (
	"github.com/google/uuid"
	"log"
	"strings"
)

func GenUUIDStr() string {

	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Default().Println(err.Error())
		return ""
	}

	uuidStr := newUUID.String()
	return strings.Replace(uuidStr, "-", "", -1)
}
