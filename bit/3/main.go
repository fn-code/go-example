package main

import (
	"encoding/json"
	"fmt"
)

type msg struct {
	m   int
	msg string
	c   int
	d   string
	h   hobi
}

type hobi struct {
	id int
	hb []string
}

func main() {
	// var a uint = 31
	// fmt.Printf("bits.Len(%d) = %d \n", a, bits.Len(a))

	hb := hobi{id: 2, hb: []string{"bola", "music"}}
	m := msg{m: 60, msg: "malam semuanya", c: 200, d: "masa olo eh iyo aaaa iyo masa aaaaaaaa", h: hb}
	buf, _ := json.Marshal(m)

	fmt.Println(128 & 128)

	mode := byte(0x81)
	mask := byte(0x80)

	frame := []byte{mode, mask | byte(len(buf))}
	data := append(frame, buf...)
	fmt.Println(data)

}
