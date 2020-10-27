package main

import (
	"fmt"
	"github.com/zach-klippenstein/goadb"
	"log"
)

var (
	client *adb.Adb
)

type device struct {
	device string
	ip     string
	name   string
}

func InitAdb() {
	var err error
	client, err = adb.New()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Starting server…")
	client.StartServer()

	serverVersion, err := client.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server version:", serverVersion)

}
