package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	eval := "3 + 2"

	tokens := tokenize(eval)

	assert.Equal(t, 3, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())

	eval = "3 + 2 + 1"

	tokens = tokenize(eval)

	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())
	assert.Equal(t, "+", tokens[3].String())
	assert.Equal(t, "1", tokens[4].String())

	eval = "3 - 2"

	tokens = tokenize(eval)

	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "-", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())

	eval = "3 + 2 * 1"

	tokens = tokenize(eval)

	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())
	assert.Equal(t, "*", tokens[3].String())
	assert.Equal(t, "1", tokens[4].String())

	eval = "3 + 2 * 1 / 2"

	tokens = tokenize(eval)

	assert.Equal(t, 7, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())
	assert.Equal(t, "*", tokens[3].String())
	assert.Equal(t, "1", tokens[4].String())
	assert.Equal(t, "/", tokens[5].String())
	assert.Equal(t, "2", tokens[6].String())
}

func TestPostfix(t *testing.T) {
	eval := "3 + 2"

	tokens := postfix(eval)
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "2", tokens[1].String())
	assert.Equal(t, "+", tokens[2].String())

	eval = "3 + 2 - 1"

	tokens = postfix(eval)
	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "2", tokens[1].String())
	assert.Equal(t, "+", tokens[2].String())
	assert.Equal(t, "1", tokens[3].String())
	assert.Equal(t, "-", tokens[4].String())
}

func TestEval(t *testing.T) {
	eval := "3 + 2"

	result, success := Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 5, result)

	eval = "3 + 2 - 1"

	result, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 4, result)
}
