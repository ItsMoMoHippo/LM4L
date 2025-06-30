package utils

import "fmt"

type LineArr struct {
	Lines []string
	Seen  map[string]bool
}

func (la *LineArr) AddLine(line string) {
	if line == "" {
		return
	}
	if !la.Has(line) {
		la.Lines = append(la.Lines, line)
		la.Seen[line] = true
	}
}

func (la *LineArr) Has(line string) bool {
	return la.Seen[line]
}

func (la *LineArr) Print() {
	for _, line := range la.Lines {
		fmt.Println(line)
	}
}
