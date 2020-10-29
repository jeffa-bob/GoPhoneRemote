package main

import (
	"io"
	"os/exec"
	"log"
)

var stream *io.ReadCloser

func StartStream(){
	locserial, err := device.Serial()
	if err != nil {
		log.Fatal(err)
	}
	//adb exec-out screenrecord --bit-rate=16m --output-format=h264 --size 1920x1080 -|ffplay -framerate 60 -framedrop -bufsize 16M -

	cmd := exec.Command("adb","-s",locserial,"exec-out","screenrecord","--bit-rate=16m","--output-format=h264","--size 1920x1080 -")
	str, err := cmd.StdoutPipe()
	stream = &str
	if err != nil {
		log.Fatal(err)
	}
	lierr := cmd.Start()
	if lierr != nil {
		log.Fatal(lierr)
	}
}
/*
func DisplayStream(){
	stream.
}*/