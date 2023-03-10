package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}
