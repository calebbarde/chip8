package cpu

import (
	"chip8/pkg/display"
	"encoding/binary"
	"log"
)

const addressableMemorySize = 0xFFF - 0x200

type CPU struct {
	memory  [addressableMemorySize]byte
	vs      [0xF]byte
	i       uint16
	sp      byte
	pc      uint16
	Display display.Displayer
	// How to keyboard?
	// Display Display
}

func (c *CPU) LoadMemory(data []byte) {
	// what happens if file is too big?
	for i, v := range data {
		c.memory[i] = v
	}
}

func (c *CPU) Run() {
	for {
		opcode := c.readOpcode()

		c.handleOpcode(opcode)
	}
}

func (c *CPU) readOpcode() uint16 {
	bytes := []byte{c.memory[c.pc], c.memory[c.pc+1]}
	var doubleWord = binary.BigEndian.Uint16(bytes)
	c.pc += 2
	return doubleWord
}

func (c *CPU) handleOpcode(opcode uint16) {
	if opcode&0x6000 == 0x6000 {
		registerIndex := readHighFourBits(opcode)
		value := readLowByte(opcode)

		c.vs[registerIndex] = value
	} else if opcode&0xA000 == 0xA000 {
		address := readLow12Bits(opcode)

		c.i = address
	} else if opcode&0xD000 == 0xD000 {
		log.Printf("opcode = %x", opcode)
		log.Printf("i = %x", c.i)
		xRegisterIndex := readHighFourBits(opcode)
		yRegisterIndex := readMidFourBits(opcode)
		nBytes := readLowFourBits(opcode)

		log.Printf("X = %x", xRegisterIndex)
		log.Printf("Y = %x", yRegisterIndex)
		log.Printf("N = %x", nBytes)

		c.Display.Draw(c.vs[xRegisterIndex], c.vs[yRegisterIndex], c.memory[c.i:c.i+uint16(nBytes)])
	} else {
		log.Fatalf("Unhandled opcode: %x", opcode)
	}
}

func readHighFourBits(opcode uint16) byte {
	return byte(opcode >> 8 & 0x0F)
}

func readMidFourBits(opcode uint16) byte {
	return byte((opcode & 0x00F0) >> 4)
}

func readLowFourBits(opcode uint16) byte {
	return byte(opcode & 0x000F)
}

func readLowByte(opcode uint16) byte {
	return byte(opcode & 0xFF)
}

func readLow12Bits(opcode uint16) uint16 {
	return opcode & 0x0FFF
}
