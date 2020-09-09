package tests

import (
	"testing"
)

func TestDemo(t *testing.T) {
	file := "./files/gear.json"
	Demo(file, 25, 10)
}
