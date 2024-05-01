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
	"fmt"

	"k8s.io/klog/v2"
)

type Lexer interface {
	Next() (Token, error)
}

type BaseParser struct {
	lexer        Lexer
	unreadTokens []Token
	err          error
}

func (p *BaseParser) Init(lexer Lexer) {
	p.lexer = lexer
	p.unreadTokens = nil
	p.err = nil
}

func (p *BaseParser) read() Token {
	if p.err != nil {
		return Token{TokenTypeError, ""}
	}
	if n := len(p.unreadTokens); n != 0 {
		t := p.unreadTokens[n-1]
		p.unreadTokens = p.unreadTokens[:n-1]
		return t
	}
	t, err := p.lexer.Next()
	if err != nil {
		p.err = err
		return Token{TokenTypeError, ""}
	}
	return t
}

func (p *BaseParser) PeekTokenType() TokenType {
	if n := len(p.unreadTokens); n != 0 {
		return p.unreadTokens[n-1].TokenType
	}
	t := p.peek()
	return t.TokenType
}

func (p *BaseParser) Expect(tt TokenType) Token {
	t := p.read()
	if t.TokenType != tt {
		if p.err == nil {
			p.err = fmt.Errorf("expected token type %v, got %v", tt, t.TokenType)
		}
		return Token{TokenTypeError, ""}
	}
	return t
}

func (p *BaseParser) Unread(t Token) {
	if p.err != nil {
		return
	}
	p.unreadTokens = append(p.unreadTokens, t)
}

func (p *BaseParser) peek() Token {
	t := p.read()
	p.Unread(t)
	return t
}

func (p *BaseParser) Unexpected() error {
	t := p.read()
	err := fmt.Errorf("unexpected token %v", t)
	p.err = err
	return err
}

func (p *BaseParser) Complete() error {
	p.PeekTokenType() // In case we just haven't touched the last token
	if p.err != nil {
		return p.err
	}

	// Check that we have consumed all input
	eof := true
	for _, t := range p.unreadTokens {
		if t.TokenType != TokenTypeEOF {
			eof = false
		}
	}
	if !eof {
		trailer := fmt.Sprintf("%v ...", p.peek())

		klog.Infof("unexpected tokens: %v", p.unreadTokens)
		return fmt.Errorf("extra input found (%s)", trailer)
	}
	return nil
}
