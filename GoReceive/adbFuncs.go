package main

import (
	//"fmt"
	"github.com/zach-klippenstein/goadb"
	"log"
	//"runtime/debug"
)

var (
	client *adb.Adb
	device *adb.DeviceDescriptor
)

func InitAdb() {
	var err error
	client, err = adb.New()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Starting server…")
	client.StartServer()
	/*
	serverVersion, err := client.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Server version:", serverVersion)
*/
}

func SelectDevice(serial string){
	for devices, _ : range client.ListDevices(){
		devices.
	}
	
}

func ListDevices() []*adb.DeviceInfo {
	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}
	return devices
}
