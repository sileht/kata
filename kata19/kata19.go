package kata19

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func charsInCommon(word1 string, word2 string) int {
	// Assume words have same length
	nb := 0
	for i := 0; i < len(word1); i += 1 {
		if word1[i] == word2[i] {
			nb += 1
		}
	}
	return nb
}

func wordAlreadyInChain(word string, chain []string) bool {
	for i := range chain {
		if chain[i] == word {
			return true
		}
	}
	return false
}

func loadWords(wordSize int) []string {
	// Load words
	dat, err := ioutil.ReadFile("wordlist-utf8.txt")
	check(err)
	words := strings.Split(string(dat), "\n")

	// Kept only useful words
	usefulWords := []string{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) == wordSize {
			usefulWords = append(usefulWords, word)
		}
	}
	return usefulWords
}

func timeTrack(tStart time.Time, name string, start string, end string) {
	elapsed := time.Since(tStart)
	fmt.Printf("%s took %s for ['%s', '%s']\n", name, elapsed, start, end)
}

func findWordChains(start string, end string) []string {
	wordSize := len(start)
	if wordSize != len(end) {
		panic("start and end must have same length")
	}

	usefulWords := loadWords(wordSize)
	chains := [][]string{{start}}
	neededCharsInCommon := wordSize - 1

	defer timeTrack(time.Now(), "FindWordChainsV1", start, end)

	// Beware time complexcity here, we have two loop, the first one
	// grow expotentialy, more the chain is long between the two words is long
	// more it will take times

	for len(chains) != 0 {
		chain := chains[0]
		lastWord := chain[len(chain)-1]

		for _, word := range usefulWords {
			if !wordAlreadyInChain(word, chain) {
				n := charsInCommon(lastWord, word)
				if n >= neededCharsInCommon {

					n = charsInCommon(end, word)
					if n >= neededCharsInCommon {
						// We found the short chain
						return append(chain, word, end)
					}

					// Create a new possible chain to evaluate
					chains = append(chains, append(chain, word))
				}
			}
		}
		chains = chains[1:]
	}

	return []string{}
}
