package main

import (
	cpu "chip8/pkg"
	"chip8/pkg/display/implementations"
	"flag"
	"log"
	"os"
)

func main() {
	romPath := flag.String("romPath", "", "Path where ROM is located")
	flag.Parse()
	if *romPath == "" {
		log.Fatalln("Must provide filepath!")
	}

	file, err := os.Open(*romPath)
	if err != nil {
		log.Fatalln(err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]byte, info.Size())

	_, err = file.Read(data)
	if err != nil {
		log.Fatalln(err)
	}

	cpu := cpu.CPU{Display: &implementations.DebugDisplay{}}
	cpu.LoadMemory(data)
	cpu.Run()

}
