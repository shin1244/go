package calculator

import "strconv"

type token interface {
	String() string
	Evaluate(tokens *[]token) (int, bool)
}

type number int
type plus struct{}
type minus struct{}
type multiply struct{}
type divide struct{}
type parser struct {
	eval         []rune
	idx          int
	parsedTokens token
}

func (p *parser) parse() bool {
	for {
		if p.idx >= len(p.eval) {
			return false
		}

		if p.eval[p.idx] != ' ' {
			break
		}

		p.idx++
	}

	switch p.eval[p.idx] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		var value int
		for p.idx < len(p.eval) {
			if p.eval[p.idx] >= '0' && p.eval[p.idx] <= '9' {
				value *= 10
				value += int(p.eval[p.idx] - '0')
				p.idx++
			} else {
				break
			}
		}
		p.parsedTokens = number(value)
		return true
	case '+':
		p.parsedTokens = plus{}
		p.idx++
		return true
	case '-':
		p.parsedTokens = minus{}
		p.idx++
		return true
	case '*':
		p.parsedTokens = multiply{}
		p.idx++
		return true
	case '/':
		p.parsedTokens = divide{}
		p.idx++
		return true
	}

	return false
}
func (n number) String() string {
	return strconv.Itoa(int(n))
}

func (n number) Evaluate(tokens *[]token) (int, bool) {
	return int(n), true
}

func (p plus) String() string {
	return "+"
}

func (p plus) Evaluate(tokens *[]token) (int, bool) {
	return 0, false
}

func (m minus) String() string {
	return "-"
}

func (m minus) Evaluate(tokens *[]token) (int, bool) {
	return 0, false
}

func (m multiply) String() string {
	return "*"
}

func (m multiply) Evaluate(tokens *[]token) (int, bool) {
	return 0, false
}

func (d divide) String() string {
	return "/"
}

func (d divide) Evaluate(tokens *[]token) (int, bool) {
	return 0, false
}

func tokenize(eval string) []token {
	tokens := []token{}
	p := &parser{eval: []rune(eval)}
	for p.parse() {
		tokens = append(tokens, p.parsedTokens)
	}

	return tokens
}

func postfix(eval string) []token {
	tokens := tokenize(eval)
	if len(tokens) == 0 {
		return tokens
	}

	postfix := make([]token, 0, len(tokens))
	var op token

	for i := 0; i < len(tokens); i++ {
		if num, ok := tokens[i].(number); ok {
			postfix = append(postfix, num)
		} else {
			if op != nil {
				postfix = append(postfix, op)
			}
			op = tokens[i]
		}
	}
	if op != nil {
		postfix = append(postfix, op)
	}

	return postfix
}

func Evaluate(eval string) (result int, success bool) {
	postfix := postfix(eval)

	for _, token := range postfix {
		result, success = token.Evaluate(&postfix)
		if !success {
			return 0, false
		}
	}

	return result, true
}
