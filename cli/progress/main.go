package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var progressItem = `-\\|/`
var progressItem2 = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
var progressItem3 = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
var progressItem4 = []string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"}
var progressItem5 = []string{"◜", "◠", "◝", "◞", "◡", "◟"}
var progressItem8 = []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}
var Stdout io.Writer = os.Stdout

func main() {
	k := 0
	for i := 0; i < 100; i++ {
		if k > (len(progressItem) - 1) {
			k = 0
		}

		fmt.Fprintf(Stdout, "\x1b[34m%s TASK 1", progressItem2[k])
		time.Sleep(130 * time.Millisecond)
		fmt.Fprintf(Stdout, "\033[2K\r") // clear current line

		//what: screen => erase screen and go home
		//	line   => erase line and go to start of line
		//	bos    => erase to begin of screen
		//	eos    => erase to end of screen
		//	bol    => erase to begin of line
		//	eol    => erase to end of line

		//clear = {
		//	'screen': '\x1b[2J\x1b[H',
		//		'line': '\x1b[2K\x1b[G',
		//		'bos': '\x1b[1J',
		//		'eos': '\x1b[J',
		//		'bol': '\x1b[1K',
		//		'eol': '\x1b[K',
		//}
		k++

	}
	//fmt.Fprintf(Stdout, "Task 1 ")
}
