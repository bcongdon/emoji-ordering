package ordering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmoji(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"foo", false},
		{"😀", true},
		{"🇺🇸", true},
		{"🤷‍♂️", true},
		{"", false},
		{"👍🏾", true},
	}

	for _, test := range tests {
		assert.Equal(t, IsEmoji(test.str), test.expected)
	}
}

func TestLess(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"😀", "🙂", true},
		{"🌕", "😞", false},
		{"😄", "😄", false},
		{"👍", "👍🏾", true},
		{"👍🏾", "👍", false},
	}

	for _, test := range tests {
		assert.Equal(t, Less(test.a, test.b), test.expected)
	}
}
