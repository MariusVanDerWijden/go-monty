package monty

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonty(t *testing.T) {
	a := []uint32{942} // in monty
	b := []uint32{813} // in monty
	n := []uint32{977}
	nInv := []uint32{7}
	result := monty(a, b, nInv, n)
	assert.Equal(t, []uint32{765846}, result)
	fmt.Print(result)
}
