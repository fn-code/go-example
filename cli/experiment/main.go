package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"math"
	"os"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

var Stdout io.Writer = os.Stdout

// BarFormat defines the bar design
var BarFormat = "[=>-]"

const fps = 25

type progress struct {
	title       string
	prependText string
	current     int
	total       int
	lastTimePrint time.Time
}

func main() {

	pgrs := &progress{
		title:   "Progsress 1",
		current: 0,
		total:   1000,
	}


	start := time.Now()
	for i := pgrs.current; i <= 100; i++ {
		pgrs.current = i * 10
		now := time.Now()
		if now.Sub(pgrs.lastTimePrint) > time.Second/fps {
			pgrs.lastTimePrint = now
		}
		pgrs.Run()
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("\n%v - %v\n", time.Since(pgrs.lastTimePrint).Seconds(), time.Since(start).Seconds())
}

func (p *progress) Run() {
	pct := percentage(p.current, p.total)
	pcts := fmt.Sprintf("%.2f%%", pct*100)
	p.prependText = fmt.Sprintf("%v of %s", pcts, "100%")

	clearCurrentLine()

	terminalWidth, _, err := terminal.GetSize(int(syscall.Stdin))
	if terminalWidth <= 0 || err != nil {
		// set terminal with to 40, if terminal width is less than 0 and err is not nil
		terminalWidth = 40
	}


	barLeftWidth := terminalWidth/3
	maxTitleWidth := barLeftWidth - utf8.RuneCountInString(p.prependText)
	if maxTitleWidth < 0 {
		maxTitleWidth = 0
	}

	title := p.title
	if utf8.RuneCountInString(title) > maxTitleWidth {
		titleDis := maxTitleWidth - 3
		if titleDis < 2 {
			titleDis = 0
			title = title[:titleDis]
		} else {
			title = title[:titleDis]+"..."
		}
	}

	titleLen := utf8.RuneCountInString(title)
	titleSpace := int(math.Max(0, float64(maxTitleWidth) - float64(titleLen)))  // (len(pcts) - 4)

	// show progress bar
	s := fmt.Sprintf("%s%s%s ",
		title,
		strings.Repeat(" ", titleSpace),
		p.prependText)

	fmt.Fprint(Stdout, s)

	barWidth := int(math.Floor(float64(terminalWidth - 7) - float64(barLeftWidth)))
	fill := int(math.Max(0, math.Floor(pct*float64(barWidth))))

	if fill > 0 {
		progChar := BarFormat[2]
		if p.current == p.total {
			progChar = BarFormat[1]
		}
		fmt.Fprintf(Stdout, "%c%s%c%s%c",
			BarFormat[0],
			strings.Repeat(string(BarFormat[1]), fill),
			progChar,
			strings.Repeat(string(BarFormat[3]), barWidth-fill),
			BarFormat[4])
	}

}

// percentage returns the percentage bound between 0.0 and 1.0
func percentage(current, total int) float64 {
	pct := float64(current) / float64(total)
	if total == 0 {
		if current == 0 {
			// When both Total and Current are 0, show a full progressbar
			pct = 1
		} else {
			pct = 0
		}
	}
	// percentage is bound between 0 and 1
	return math.Min(1, math.Max(0, pct))
}

func clearCurrentLine() {
	fmt.Fprintf(Stdout, "\033[2K\r")
}
