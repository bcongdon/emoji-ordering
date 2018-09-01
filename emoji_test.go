package ordering

import (
	"fmt"
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
		slice := EmojiSlice([]string{test.a, test.b})
		assert.Equal(t, slice.Less(0, 1), test.expected)
	}
}

func TestNewEmojiSlice(t *testing.T) {
	valid := []string{"😀", "😃", "😄"}
	emojiSlice, err := NewEmojiSlice(valid)
	assert.Equal(t, EmojiSlice(valid), emojiSlice)
	assert.Nil(t, err)

	invalid := []string{"foo", "😃", "😄"}
	emojiSlice, err = NewEmojiSlice(invalid)
	assert.Nil(t, emojiSlice)
	assert.Equal(t, err, fmt.Errorf("'foo' is not an emoji"))
}
