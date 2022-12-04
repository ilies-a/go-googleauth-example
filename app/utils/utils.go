package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateStringUUID() string {
	return fmt.Sprintf("%x", uuid.New())
}
