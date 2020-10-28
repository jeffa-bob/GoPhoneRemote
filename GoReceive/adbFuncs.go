package main

import (
	//"fmt"
	"github.com/zach-klippenstein/goadb"
	"log"
	"strconv"
	//"runtime/debug"
)

var (
	client *adb.Adb
	device *adb.Device
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

func SelectDevice(serial string) {
	device = client.Device(adb.DeviceWithSerial(serial))
	//adb exec-out screenrecord --bit-rate=16m --output-format=h264 --size 1920x1080 -|ffplay -framerate 60 -framedrop -bufsize 16M -
}

func TapSpot(x int64, y int64) {
	a, err := device.RunCommand("input touchscreen tap " + strconv.FormatInt(x, 10) + " " + strconv.FormatInt(y, 10))
	println(a)
	if err != nil {
		log.Fatal(err)
	}
}

func ListDevices() []*adb.DeviceInfo {
	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}
	return devices
}
