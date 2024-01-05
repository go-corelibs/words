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

package words

// List returns a list of words that were separated by spaces using the
// Default Words configuration
func List(input string) (list []string) {
	list = Default().List(input)
	return
}

// Range iterates over the list of parsed words using the Default Words
// configuration
func Range(input string, fn func(word string)) {
	Default().Range(input, fn)
	return
}

// Count returns the total number of parsed words using the Default Words
// configuration
func Count(input string) (count int) {
	count = Default().Count(input)
	return
}

// Parse returns the list of parsed words using the Default Words
// configuration
func Parse(input string) (words []string) {
	words = Default().Parse(input)
	return
}

// Search performs a very simple keyword search of the content using the
// Default Words configuration
func Search(query, content string) (score int, present []string) {
	score, present = Default().Search(query, content)
	return
}

// Metrics parses the contents and returns some interesting ReadingMetrics
// using the Default Words configuration
func Metrics(content string) (m ReadingMetrics) {
	m = Default().Metrics(content)
	return
}
