// Copyright (c) 2023  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package words provides a means for counting the numbers of words and
// estimating the reading time of the given content.
//
// Words per minute values based on:
//
//	https://irisreading.com/what-is-the-average-reading-speed/
//	https://www.researchgate.net/publication/332380784_How_many_words_do_we_read_per_minute_A_review_and_meta-analysis_of_reading_rate
//	https://www.sciencedirect.com/science/article/abs/pii/S0749596X19300786
//
// Package words was inspired by:
//
//	https://github.com/byn9826/words-count/blob/master/src/globalWordsCount.js
package words

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-corelibs/maths"
	"github.com/go-corelibs/slices"
)

var (
	rxSpaces   = regexp.MustCompile(`\s+`)
	rxSymbols  = regexp.MustCompile(`\p{S}`)
	rxCharSets = regexp.MustCompile(`[\p{Han}\p{Katakana}\p{Hiragana}\p{Hangul}]`)
)

// Words is the definition for running customized word operations and is the
// implementation driving the normal package functions
type Words struct {
	// PunctuationAsBreaker specifies that punctuation characters should not
	// be removed and instead be replaced with a space. For example, "they're"
	// by default is collapsed to "theyre" for counting purposes. With
	// PunctuationAsBreaker set to true, "they're" would become "they re"
	PunctuationAsBreaker bool
	// DisableDefaultPunctuation specifies that only the Words.Punctuation
	// runes are to be considered punctuation
	DisableDefaultPunctuation bool
	// Punctuation defines the list of punctuation runes to use when parsing
	// words out of content
	Punctuation []rune
	// AverageWPM specifies the average words per minute to use
	// for calculating Metrics, default is 238.0, see: AverageWordsPerMinute
	AverageWPM float64
	// RelaxedWPM specifies the average words per minute to use
	// for calculating Metrics, default is 177.0, see: RelaxedWordsPerMinute
	RelaxedWPM float64

	combined []rune
}

// Default returns a new Words instance configured with sane defaults
func Default() (w *Words) {
	return &Words{
		AverageWPM: AverageWordsPerMinute,
		RelaxedWPM: RelaxedWordsPerMinute,
	}
}

func (w *Words) prepare() {
	if w.AverageWPM <= 0 {
		w.AverageWPM = AverageWordsPerMinute
	}
	if w.RelaxedWPM <= 0 {
		w.RelaxedWPM = RelaxedWordsPerMinute
	}

	if !w.DisableDefaultPunctuation {
		w.combined = append(w.combined, DefaultPunctuation...)
	}
	if len(w.Punctuation) > 0 {
		w.combined = append(w.combined, w.Punctuation...)
	}
	w.combined = slices.Unique(w.combined)
	return
}

// List returns a list of all the words detected within the given input that
// are separated by spaces, word characters not separated by spaces are
// clumped within individual items of the list returned. Use Words.Parse to
// derive a more accurate word list
func (w *Words) List(input string) (list []string) {
	w.prepare()

	var work, spacer string
	if work = strings.TrimSpace(input); work == "" {
		return
	} else if w.PunctuationAsBreaker {
		spacer = " "
	}

	// replace all punctuation with spacers
	if len(w.combined) > 0 {
		for _, r := range w.combined {
			work = strings.ReplaceAll(work, string(r), spacer)
		}
	}

	work = rxSymbols.ReplaceAllString(work, "")
	work = rxSpaces.ReplaceAllString(work, " ")
	work = strings.TrimSpace(work)

	list = strings.Split(work, " ")
	return
}

// Range iterates over all words detected within input, calling the given `fn`
// for each word found
func (w *Words) Range(input string, fn func(word string)) {
	w.prepare()

	list := w.List(input)

	for _, word := range list {
		if rxCharSets.MatchString(word) {

			// handle mixed-locales
			var latin string
			var carry []string
			for _, r := range word {
				if char := string(r); rxCharSets.MatchString(char) {
					if latin != "" {
						carry = append(carry, latin)
						latin = ""
					}
					carry = append(carry, char)
				} else if !slices.Within(r, w.combined) {
					latin += char
				}
			}
			if latin != "" {
				carry = append(carry, latin)
				latin = ""
			}
			for _, carriedWord := range carry {
				fn(carriedWord)
			}

		} else {

			fn(word)

		}
	}

}

// Count returns the total number of words detected within the given input
func (w *Words) Count(input string) (count int) {
	w.Range(input, func(_ string) {
		count += 1
	})
	return
}

// Parse returns the total list of words detected within the given input
func (w *Words) Parse(input string) (words []string) {
	w.Range(input, func(word string) {
		words = append(words, word)
	})
	return
}

// Search performs a case-insensitive search for the keywords within the given
// `query` string and returns the list of unique query keywords found along
// with a simple scoring metric weighing earlier keywords more than later
// keywords
func (w *Words) Search(query, content string) (score int, found []string) {
	w.prepare()

	keywords := w.Parse(strings.ToLower(query))
	keywordCount := len(keywords)
	haystack := w.Parse(strings.ToLower(content))

	for _, word := range haystack {
		for idx, keyword := range keywords {
			if word == keyword {
				weight := keywordCount - idx
				score += weight
				found = append(found, word)
			}
		}
	}

	found = slices.Unique(found)
	return
}

// Metrics gets the Words.Count and derives some estimated reading times
func (w *Words) Metrics(content string) (m ReadingMetrics) {
	w.prepare()

	m.WordCount = w.Count(content)

	avgTime := float64(m.WordCount) / w.AverageWPM
	m.Average.Minutes = maths.RoundDown(avgTime)
	m.Average.Duration = time.Duration(avgTime * float64(time.Minute))

	relTime := float64(m.WordCount) / w.RelaxedWPM
	m.Relaxed.Minutes = maths.RoundUp(relTime)
	m.Relaxed.Duration = time.Duration(relTime * float64(time.Minute))
	return
}
