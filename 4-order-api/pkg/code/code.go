package code

import (
	"fmt"
	"math/rand"
)

const CODE_LENGHT = 10000

func GenerateCode() string {
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
