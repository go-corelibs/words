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

func TestWords(t *testing.T) {
	Convey("Custom Settings", t, func() {
		w := Default()
		w.AverageWPM = -1
		w.RelaxedWPM = -1
		w.Punctuation = []rune{'!'}
		w.PunctuationAsBreaker = true
		So(w.List(`they're one two`), ShouldEqual, []string{
			"they", "re", "one", "two",
		})
		m := w.Metrics(`one two`)
		So(m.WordCount, ShouldEqual, 2)
		So(m.Average.Minutes, ShouldEqual, 0)
		So(m.Relaxed.Minutes, ShouldEqual, 1)
	})
}
