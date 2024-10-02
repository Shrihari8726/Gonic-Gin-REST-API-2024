package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const prefix = "DAYT"

var counter = 1

func GenerateID() string {
	numberPart := strconv.Itoa(counter)
	numberPart = strings.Repeat("0", 4-len(numberPart)) + numberPart

	id := prefix + numberPart

	counter++

	fmt.Println(id)

	return id
}
