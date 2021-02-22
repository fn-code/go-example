package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	b := []byte{0x42, 0xF6, 0xE6, 0x66}
	//bb := []uint16{66, 246, 230, 102}

	//bits := math.Float64bits(123.45)
	//bytes := make([]byte, 8)
	//binary.LittleEndian.PutUint64(bytes, bits)
	//
	//fmt.Println(bytes)

	res := binary.LittleEndian.Uint16(b)
	//float := math.Float64frombits(res)
	fmt.Println(res)

}
