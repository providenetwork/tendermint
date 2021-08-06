package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/providenetwork/tendermint/types"
)

func TestMakeHeader(t *testing.T) {
	_, err := MakeHeader(&types.Header{})
	assert.NoError(t, err)
}
