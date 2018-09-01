package ordering

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

func Less(a, b string) bool {
	var aIdx int
	var bIdx int

	if IsEmoji(a) {
		aIdx = emojiIndices[a]
	}
	if IsEmoji(b) {
		bIdx = emojiIndices[b]
	}

	return aIdx < bIdx
}
