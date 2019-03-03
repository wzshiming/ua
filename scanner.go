package ua

type Scanner struct {
	buf []rune
	ch  rune
	off int
}


func NewScanner(buf []rune) *Scanner {
	s := &Scanner{
		buf: buf,
	}
	s.next()
	return s
}

func (s *Scanner) skipSpace() {
	for {
		switch s.ch {
		case ' ', '\r', '\t':
			s.next()
		default:
			return
		}
	}
}

func (s *Scanner) next() {
	if len(s.buf) <= s.off {
		s.ch = -1
		s.off = len(s.buf) + 1
		return
	}

	s.ch = s.buf[s.off]
	s.off++
	return
}

func (s *Scanner) Scan() (tok Token, val string) {
	s.skipSpace()
	switch {
	//	case s.ch >= '0' && s.ch <= '9', s.ch >= 'a' && s.ch <= 'z', s.ch >= 'A' && s.ch <= 'Z', s.ch == '_', s.ch == '~', s.ch == '-', s.ch == '.':
	//		tok = Literal
	//		val = s.scanIdent()
	//		return

	case s.ch == '/', s.ch == ';', s.ch == ':', s.ch == ',', s.ch == '(', s.ch == ')':
		val = string([]rune{s.ch})
		tok = Symbol
		s.next()
		return
	case s.ch == -1:
		val = ""
		tok = EOF
		return
	default:
		tok = Literal
		val = s.scanLiteral()
		return
	}
	return
}

func (s *Scanner) scanLiteral() string {
	off := s.off - 1
loop:
	for {
		switch s.ch {
		case '/',
			'\t', ' ',
			'(', ')',
			',', ';', -1:
			break loop
		}
		s.next()
	}
	return string(s.buf[off : s.off-1])
}

func (s *Scanner) scanLiteral2() string {
	off := s.off - 1
loop:
	for {
		switch s.ch {
		case '\t', ' ',
			'(', ')', -1:
			break loop
		}
		s.next()
	}
	return string(s.buf[off : s.off-1])
}
