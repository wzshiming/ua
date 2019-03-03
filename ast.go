package ua

import (
	"fmt"
	"strings"
)

type Pairs []*Pair

type Pair struct {
	Name    string // Mozilla
	Col     string // /
	Subname string // 5.0
	Other   Pairs  // (Linux; Android 4.4.2; Infinix X509 Build/KOT49H)
	End     string // ;
}

func (p Pairs) String() string {
	ss := []string{}
	for _, v := range p {
		ss = append(ss, v.String())
	}
	return strings.Join(ss, " ")
}

func (p *Pair) String() string {
	if len(p.Other) == 0 {
		return fmt.Sprintf("%s%s%s%s", p.Name, p.Col, p.Subname, p.End)
	}
	if p.Name == "" {
		return fmt.Sprintf("(%s)%s", p.Other, p.End)
	}
	return fmt.Sprintf("%s%s%s (%s)%s", p.Name, p.Col, p.Subname, p.Other, p.End)
}
