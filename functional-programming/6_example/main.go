package main

import (
	"errors"
	"fmt"
	"log"
)

const DASHES = "-----------------------"

type Pond struct {
	BugSupply       int
	StrokesRequired int
}

type StrokeBehavior interface {
	PaddleFoot(*int)
}

type EatBehavior interface {
	EatBug(*int)
}

type SurvivalBehavior interface {
	StrokeBehavior
	EatBehavior
}

type Duck struct{}

type Foot struct{}

func (Foot) PaddleFoot(ss *int) {
	fmt.Println("- Foot, paddle!")
	*ss--
}

type Bill struct{}

func (Bill) EatBug(ss *int) {
	*ss++
	fmt.Println("- Bill, eat a bug!")
}

func (Duck) Stroke(s StrokeBehavior, ss *int, p Pond) (err error) {
	for i := 0; i < p.StrokesRequired; i++ {
		if *ss < p.StrokesRequired-i {
			err = errors.New("Our duck died!")
		}
		s.PaddleFoot(ss)
	}
	return err
}

func (Duck) Eat(e EatBehavior, ss *int, p Pond) {
	for i := 0; i < p.BugSupply; i++ {
		e.EatBug(ss)
	}
}

func (d Duck) SwimAndEat(se SurvivalBehavior, ss *int, ps []Pond) {
	for i := range ps {
		pond := &ps[i]
		err := d.Stroke(se, ss, *pond)
		if err != nil {
			log.Fatal(err)
		}
		d.Eat(se, ss, *pond)
	}
}

type Capabilities struct {
	StrokeBehavior
	EatBehavior
	strokes int
}

func displayDuckStats(c *Capabilities, ps []Pond) {
	fmt.Printf("%s\n", DASHES)
	fmt.Printf("Ponds Processed:")
	for _, pond := range ps {
		fmt.Printf("\n\t%+v", pond)
	}
	fmt.Printf("\nStrokes remaining: %+v\n", c.strokes)
	fmt.Printf("%s\n\n", DASHES)
}

func main() {
	caps := Capabilities{
		StrokeBehavior: Foot{},
		EatBehavior:    Bill{},
		strokes:        2,
	}

	// ps := []Pond{
	// 	{BugSupply: 1, StrokesRequired: 3},
	// 	{BugSupply: 1, StrokesRequired: 2},
	// }
	ps := []Pond{
		{BugSupply: 2, StrokesRequired: 3},
	}

	duck := &Duck{}
	duck.SwimAndEat(&caps, &caps.strokes, ps)
	displayDuckStats(&caps, ps)
}
