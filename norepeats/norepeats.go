package norepeats

type countWithIndex struct {
	Count     int
	LastIndex int
}

func NonRepeat(sentence string) rune {
	characterCount := make(map[rune]countWithIndex)

	for index, char := range sentence {
		characterCount[char] = countWithIndex{
			Count:     characterCount[char].Count + 1,
			LastIndex: index,
		}
	}

	highestIndex := -1
	resultChar := rune(0)

	for char, charData := range characterCount {
		if charData.Count == 1 {
			if charData.LastIndex > highestIndex {
				highestIndex = charData.LastIndex
				resultChar = char
			}
		}
	}

	return resultChar
}
