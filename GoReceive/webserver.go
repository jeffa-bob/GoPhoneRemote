package main

import (
	"math/rand"
	"strconv"
)

var temp string

func GenerateCode() string {
	temp = strconv.FormatInt(rand.Int63(), 10)
	return temp
}
