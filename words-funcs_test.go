// Copyright (c) 2024  The Go-Curses Authors
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

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFuncs(t *testing.T) {
	Convey("List", t, func() {
		words := List(`one two`)
		So(len(words), ShouldEqual, 2)
		words = List(` `)
		So(len(words), ShouldEqual, 0)
	})

	Convey("Range", t, func() {
		var count int
		Range(`one two`, func(word string) {
			count += 1
		})
		So(count, ShouldEqual, 2)
	})

	Convey("Count", t, func() {
		count := Count(`さらに「やり遂げる」ためのEnjin`)
		So(count, ShouldEqual, 12)
	})

	Convey("Parse", t, func() {
		words := Parse(`さらに「やり遂げる」ためのEnjin`)
		So(len(words), ShouldEqual, 12)
	})

	Convey("Search", t, func() {
		score, present := Search(`word`, `one word two word`)
		So(score, ShouldEqual, 2)
		So(present, ShouldEqual, []string{"word"})
		score, present = Search("enjin", "さらに「やり遂げる」ためのEnjin")
		So(score, ShouldEqual, 1)
		So(present, ShouldEqual, []string{"enjin"})
		score, present = Search("enjin", "さらに「やり遂げる」ためのEnjinさらに「やり遂げる」ためのEnjin")
		So(score, ShouldEqual, 2)
		So(present, ShouldEqual, []string{"enjin"})
	})

	Convey("Metrics", t, func() {
		m := Metrics(`one two`)
		So(m.WordCount, ShouldEqual, 2)
		So(m.Average.Minutes, ShouldEqual, 0)
		So(m.Relaxed.Minutes, ShouldEqual, 1)
	})
}
