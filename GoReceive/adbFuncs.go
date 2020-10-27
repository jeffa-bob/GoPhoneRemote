package main

import (
	"fmt"
	"github.com/zach-klippenstein/goadb"
	"log"
	//"runtime/debug"
)

var (
	client *adb.Adb
)

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

func ListDevices() []*adb.DeviceInfo {
	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("\t%+v\n", *device)
	}
	return devices
}
