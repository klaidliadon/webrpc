package ridl

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	buf := `


		ridl						 =  v1

				+     foo=bar

					-baz   = 56 # a comment

							         		version=                    v0.0.1


foo=bar`

	tokens, err := Tokenize(buf)
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	log.Printf("tokens: %v", tokens)
}

func TestParse(t *testing.T) {
	fp, err := os.Open("example1.ridl")
	assert.NoError(t, err)

	buf, err := ioutil.ReadAll(fp)
	assert.NoError(t, err)

	tokens, err := Tokenize(string(buf))
	assert.NoError(t, err)

	log.Printf("buf: %v", string(buf))
	log.Printf("tokens: %v", tokens)
}