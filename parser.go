package ua

type parser struct {
	s   *Scanner
	tok Token
	val string
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Scan() {
	p.tok, p.val = p.s.Scan()
}

func (p *parser) Parse(u string) (Pairs, error) {
	p.s = NewScanner([]rune(u))
	p.Scan()
	ps, err := p.parsePairs()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (p *parser) parsePairs() (Pairs, error) {
	pairs := Pairs{}
	for p.tok != EOF {
		pair, err := p.parsePair()
		if err != nil {
			return nil, err
		}
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs, nil
}

func (p *parser) parsePair() (*Pair, error) {
	pair := &Pair{}
	if p.tok == EOF {
		return nil, nil
	}
loop:
	for {
		switch p.tok {
		case Literal:
			if pair.Name == "" {
				pair.Name = p.val
			} else {
				pair.Name += " " + p.val
			}
		case Symbol:
			break loop
		case EOF:
			return pair, nil
		}
		p.Scan()
	}

	switch p.val {
	case "/", ":":
		pair.Col = p.val
		p.Scan()
		switch p.tok {
		case Literal:
			pair.Subname = p.val
			p.Scan()
		default:

		}
	}

	switch p.tok {
	case Symbol:
		switch p.val {
		case "/":
			pair.Name += pair.Col
			pair.Name += pair.Subname
			pair.Col = ""
			pair.Subname = ""
			for p.tok != EOF && p.val != ")" {
				pair.Name += p.val
				p.Scan()
			}
			p.tok = EOF
			return pair, nil
		case "(":
			p.Scan()
			ps, err := p.parsePairs()
			if err != nil {
				return nil, err
			}
			pair.Other = ps
			p.Scan()
		case ")":
			p.tok = EOF
			return pair, nil
		default:

		}

	default:
	}

	switch p.tok {
	case Symbol:
		switch p.val {
		case "(", ")":
		default:
			pair.End = p.val
			p.Scan()
		}

	default:
	}

	return pair, nil
}
