package ridl

func parserStateBasicMeta(en *BasicNode) parserState {
	return func(p *parser) parserState {

		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			_, err := p.match(tokenPlusSign, tokenWhitespace)
			if err != nil {
				return p.stateError(err)
			}

			// tag
			left, err := p.expectMetadataKey()
			if err != nil {
				return p.stateError(err)
			}

			// =
			_, err = p.match(tokenWhitespace, tokenEqual, tokenWhitespace)
			if err != nil {
				return p.stateError(err)
			}

			// - or value
			right, err := p.expectMetadataValue()
			if err != nil {
				return p.stateError(err)
			}

			en.meta = append(en.meta, &DefinitionNode{
				leftNode:  newTokenNode(left),
				rightNode: newTokenNode(right),
			})

			return parserStateBasicMetaDefinition(en)

		default:
			p.emit(en)
			return parserDefaultState

		}

		return parserStateBasicMeta(en)
	}
}

func parserStateBasicMetaDefinition(et *BasicNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateBasicMetaDefinition(et)

		default:
			return parserStateBasic(p)

		}

		p.emit(et)
		return parserDefaultState

	}
}

func parserStateBasic(p *parser) parserState {
	// enum <name>: <type>[<# comment>]
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenColon, tokenWhitespace)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordBasic {
		return p.stateError(errUnexpectedToken)
	}

	typeToken, err := p.expectType()
	if err != nil {
		return p.stateError(err)
	}

	return parserStateBasicMeta(&BasicNode{
		name:      newTokenNode(matches[2]),
		basicType: newTokenNode(typeToken),
		comment:   parseComments(p.comments, matches[0].line),
	})
}
