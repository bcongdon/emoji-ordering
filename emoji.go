package ordering

import (
	"math"
)

var emojiIndices = make(map[string]int)

func init() {
	for idx, emoji := range Emojis {
		emojiIndices[emoji] = idx
	}
}

func IsEmoji(e string) bool {
	_, ok := emojiIndices[e]
	return ok
}

type EmojiSlice []string

func (e EmojiSlice) Less(a, b int) bool {
	aVal := math.MaxInt32
	bVal := math.MaxInt32

	if IsEmoji(e[a]) {
		aVal = emojiIndices[e[a]]
	}
	if IsEmoji(e[b]) {
		bVal = emojiIndices[e[b]]
	}

	return aVal < bVal
}

func (e EmojiSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e EmojiSlice) Len() int {
	return len(e)
}
