// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lexparse

import (
	"io"
	"strings"
)

type BaseLexer struct {
	reader      *strings.Reader
	unreadRunes []rune
	err         error
	eof         bool
}

const (
	LexerRuneEOF   rune = -1
	LexerRuneError rune = -2
)

func (l *BaseLexer) Init(s string) {
	l.reader = strings.NewReader(s)
	l.eof = false
	l.err = nil
	l.unreadRunes = l.unreadRunes[:0]
}

func (l *BaseLexer) Unread(r rune) {
	if l.err != nil {
		return
	}
	l.unreadRunes = append(l.unreadRunes, r)
}

func (l *BaseLexer) Read() rune {
	if l.err != nil {
		return LexerRuneError
	}

	if n := len(l.unreadRunes); n != 0 {
		r := l.unreadRunes[n-1]
		l.unreadRunes = l.unreadRunes[:n-1]
		return r
	}

	if l.eof {
		return LexerRuneEOF
	}

	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.eof = true
			return LexerRuneEOF
		} else {
			l.err = err
			return LexerRuneError
		}
	}
	return r
}

func (l *BaseLexer) Peek() rune {
	r := l.Read()
	l.Unread(r)
	return r
}

func (l *BaseLexer) Err() error {
	if l.err == nil && l.eof {
		return io.EOF
	}
	return l.err
}
