package implementations

import (
	"log"
)

type DebugDisplay struct {
	screen [32]uint64
}

func (d *DebugDisplay) Draw(x, y byte, data []byte) {
	for v, i := range data {
		line = screen[y+i]

		log.Printf("% 08b", v)
	}
	// TODO WE NEED TO XOR AT THE RIGHT LOCATIONS
	// d.screen[x][y]
}

func (d *DebugDisplay) Update() {

}
