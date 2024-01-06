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

type Parser struct {
	lexparse.BaseParser
}

func (p *Parser) Init(s string) {
	l := &lexer{}
	l.Init(s)
	p.BaseParser.Init(l)
}

func (p *Parser) ParseMustacheExpression() (*MustacheExpression, error) {
	switch p.PeekTokenType() {
	case tokenTypeLeftBracket:
		p.Expect(tokenTypeLeftBracket)
		t := p.Expect(tokenTypeOther)
		p.Expect(tokenTypeRightBracket)
		return &MustacheExpression{Expression: t.Value}, nil
	default:
		return nil, p.Unexpected()
	}
}

func (p *Parser) ParseLiteralExpression() (*LiteralExpression, error) {
	switch p.PeekTokenType() {
	case tokenTypeOther:
		t := p.Expect(tokenTypeOther)
		return &LiteralExpression{Literal: t.Value}, nil

	default:
		return nil, p.Unexpected()
	}
}

func (p *Parser) ParseExpression() (Expression, error) {
	switch p.PeekTokenType() {
	case tokenTypeOther:
		return p.ParseLiteralExpression()
	case tokenTypeLeftBracket:
		return p.ParseMustacheExpression()

	default:
		return nil, p.Unexpected()
	}
}

func (p *Parser) ParseExpressionList() (*ExpressionList, error) {
	el := &ExpressionList{}

	first, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}
	el.Expressions = append(el.Expressions, first)

	for {
		switch p.PeekTokenType() {
		case tokenTypeLeftBracket, tokenTypeOther:
			e, err := p.ParseExpression()
			if err != nil {
				return nil, err
			}
			el.Expressions = append(el.Expressions, e)

		case tokenTypeEOF:
			return el, nil

		default:
			return nil, p.Unexpected()
		}
	}
}
