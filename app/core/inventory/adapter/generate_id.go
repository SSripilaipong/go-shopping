package adapter

import "github.com/google/uuid"

func GenerateItemId() string {
	return "INV-ITM-" + uuid.New().String()
}
