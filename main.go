package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var FormatError = errors.New("expected format letter:number, i.e. M:100, D:80, B:70, P:50")

type Attack int
type Defend int
type Action interface{}

func Parse(str string) (action Action, err error) {
	var name string
	var value int
	n, err := fmt.Sscanf(str, "%1s:%d", &name, &value)
	if n != 2 {
		err = FormatError
		return
	}
	name = strings.ToUpper(name)
	if value < 0 {
		value = 0
	}
	if name == "M" {
		action = Attack(value)
	} else if name == "D" {
		action = Defend(value)
	} else if name == "B" {
		action = Defend(value)
	} else if name == "P" {
		action = Defend(value)
	} else {
		err = FormatError
	}
	return
}

func calculate(attacks []Attack, actions []Action, defended bool) (hits, total uint64) {

	if len(actions) > 0 {
		action1, actions2 := actions[0], actions[1:]

		switch action := action1.(type) {
		case Attack:
			for a := Attack(0); a < action; a++ {
				h, t := calculate(append(attacks, a), actions2, defended)
				hits += h
				total += t
			}
			return

		case Defend:
			if len(attacks) > 0 {
				attack1, attacks2 := attacks[0], attacks[1:]
				for d := Defend(0); d < action; d++ {
					if attack1 > Attack(d) {
						h, t := calculate(attacks, actions2, defended)
						hits += h
						total += t
					} else {
						h, t := calculate(attacks2, actions2, true)
						hits += h
						total += t
					}
				}
				return
			}

			h, t := calculate(attacks, actions2, defended)
			hits += h * uint64(action)
			total += t * uint64(action)
			return
		}
	}

	if len(attacks) > 0 {
		// hit
		hits += 1
		total += 1
	} else if defended {
		// dodged or parried
		total += 1
	}
	return
}

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: <actions>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, " Actions can be: letter:number, i.e. M:100, D:80, B:70, P:50.\n")
		fmt.Fprintf(os.Stderr, " The order matters.\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal(FormatError)
	}

	actions := make([]Action, 0, 10)
	for _, arg := range flag.Args() {
		if action, err := Parse(arg); err != nil {
			log.Fatal(err)
		} else {
			actions = append(actions, action)
		}
	}

	hits, total := calculate(nil, actions, false)
	fmt.Printf("%d hits (%2.3f%%), %d misses (%2.3f%%), %d total\n",
		hits, 100*float32(hits)/float32(total),
		total-hits, 100*float32(total-hits)/float32(total),
		total,
	)
}
