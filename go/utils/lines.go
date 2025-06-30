package utils

import "fmt"

type LineArr struct {
	lines []string
	seen  map[string]bool
}

func (la *LineArr) AddLine(line string) {
	if !la.Has(line) {
		la.lines = append(la.lines, line)
	}
}

func (la *LineArr) Has(line string) bool {
	return la.seen[line]
}

func (la *LineArr) Print() {
	for _, line := range la.lines {
		fmt.Println(line)
	}
}
