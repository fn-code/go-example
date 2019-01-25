package main

import (
	"fmt"
)

func main() {
	// Use bitwise OR | to get bits that are in 1 OR 2
	fmt.Println("OR : ", 1|2)

	fmt.Println("OR : ", 1|5)

	// Use bitwise XOR ^ to get bits that are in 3 OR 6 but NOT BOTH
	// 3 		= 00000011
	// 6 		= 00000110
	// 3 ^ 6 	= 00000101 = 5
	fmt.Println("XOR : ", 3^6)
	// 3 		= 00000011
	// 7 		= 00000111
	// 3 ^ 7 	= 00000100 = 4
	fmt.Println("XOR : ", 3^7)
	// 3 		= 00000011
	// 11		= 00001011
	// 3 ^ 11 	= 00001000 = 8
	fmt.Println("XOR : ", 3^11)

	// 3 		= 00000011
	// 12		= 00001100
	// 3 ^ 12 	= 00001111 = 15
	fmt.Println("XOR : ", 3^12)

	// Use bitwise AND & to get the bits that are in 3 AND 6
	// 3     = 00000011
	// 6     = 00000110
	// 3 & 6 = 00000010 = 2
	fmt.Println("AND : ", 3&6)

	// Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
	// 26     = 00001110
	// 6      = 00000110
	// 3 &^ 6 = 00001000 = 1
	fmt.Println("AND NOT2 : ", 14&^6)

	// Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
	// 26     = 00011010
	// 6      = 00000110
	// turn     11111001
	// 3 &^ 6 = 00011000 = 1
	fmt.Println("AND NOT : ", 26&^6)

	// 5  00000101
	//-5  11111010+1=1111011 (two's complement of 5, the value was -5)

	// 6  00000110
	//-6  11111001+1=11111010

	// 5  00000101
	//-6  11111010 (one's complement of 5, the value was -6)
	fmt.Println("dd : ", ^6)

	a := 1
	a |= 1 << 2
	// 0000100 | 00000001 = 0000101 =5

	fmt.Println(a)
	a |= 1 << 6
	// 0000101 | 0100000 = 0100101 =69
	fmt.Println(a)

	fmt.Println(1 << 6)
	//0000001 << 6 = 1000000 = 64

}
