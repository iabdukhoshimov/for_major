package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	c := Load()

	ast := assert.New(t)

	ast.NotNil(c)
}
