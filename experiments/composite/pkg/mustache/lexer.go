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

package mustache

import "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/lexparse"

type lexer struct {
	lexparse.BaseLexer
}

func (l *lexer) Init(s string) {
	l.BaseLexer.Init(s)
}

type token = lexparse.Token

const (
	tokenTypeLeftBracket  lexparse.TokenType = '{'
	tokenTypeRightBracket                    = '}'
	tokenTypeOther                           = '.'
	tokenTypeEOF                             = lexparse.TokenTypeEOF
	tokenTypeError                           = lexparse.TokenTypeError
)

func (l *lexer) lexOther(first rune) (token, error) {
	var s []rune
	s = append(s, first)
runeLoop:
	for {
		r := l.Read()

		switch r {
		case lexparse.LexerRuneEOF:
			break runeLoop

		case lexparse.LexerRuneError:
			return token{}, l.Err()

		case '$', '}':
			l.Unread(r)
			break runeLoop
		default:
			s = append(s, r)
		}
	}
	return token{TokenType: tokenTypeOther, Value: string(s)}, nil
}

func (l *lexer) Next() (token, error) {
	r := l.Read()

	switch r {
	case lexparse.LexerRuneEOF:
		return token{TokenType: tokenTypeEOF, Value: ""}, nil

	case lexparse.LexerRuneError:
		return token{}, l.Err()

	case '$':
		r2 := l.Peek()
		if r2 == '{' {
			l.Read()
			return token{TokenType: tokenTypeLeftBracket, Value: "${"}, nil
		}

	case '}':
		return token{TokenType: tokenTypeRightBracket, Value: "}"}, nil
	}

	return l.lexOther(r)
}
