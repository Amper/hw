package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var wordsMatcher = regexp.MustCompile(`(\p{L}+([\p{L}-]+\p{L}+)*)`)

type wordStat struct {
	word  string
	count uint
}

func Top10(value string) []string {
	return TopN(value, 10)
}

func TopN(value string, count int) []string {
	words := matchWords(value)
	stat := getWordsStat(words)
	result := getTopWords(stat, count)
	return result
}

func matchWords(value string) []string {
	return wordsMatcher.FindAllString(value, -1)
}

func getWordsStat(words []string) []wordStat {
	// count of each word
	counter := map[string]uint{}
	for _, word := range words {
		counter[strings.ToLower(word)]++
	}

	// create array of statistics: pairs of word and count
	stat := make([]wordStat, 0)
	for word, count := range counter {
		stat = append(stat, wordStat{word, count})
	}

	return stat
}

func getTopWords(stat []wordStat, count int) []string {
	// sort statistics for top
	sort.Slice(stat, func(i, j int) bool {
		iCount, jCount, iWord, jWord := stat[i].count, stat[j].count, stat[i].word, stat[j].word
		if iCount == jCount {
			// in lexicographic orders - if the count matches
			return iWord < jWord
		}
		// in descending order of count
		return iCount > jCount
	})

	// create array with top words
	result := make([]string, 0, count)
	for num, pair := range stat {
		if num >= count {
			break
		}
		result = append(result, pair.word)
	}

	return result
}
