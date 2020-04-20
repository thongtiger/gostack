package util

import (
	"math/rand"

	"github.com/google/uuid"
)

func GenUUID() string {
	val, _ := uuid.NewRandom()
	return val.String()
}

func GenUint32() uint32 {
	return rand.Uint32()
}
