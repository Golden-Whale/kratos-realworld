package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("123")
	spew.Dump(s)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$2As7wtY/Hy4RB4COXLkaV.drXq//Gq1vWl6SdBCTsDq5VQJRLQ05y", "1234"))
}
