// Code generated by "enumer -type=TokenKind -json"; DO NOT EDIT.

package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TokenKindName = "TokenKindCharTokenKindDotTokenKindStarTokenKindPlusTokenKindQuestionTokenKindPipeTokenKindEscapeTokenKindCaretTokenKindDollarTokenKindLParenTokenKindRParenTokenKindLBraceTokenKindRBraceTokenKindLBracketTokenKindRBracketTokenKindEOFTokenKindNumberTokenKindCommaTokenKindHyphen"

var _TokenKindIndex = [...]uint16{0, 13, 25, 38, 51, 68, 81, 96, 110, 125, 140, 155, 170, 185, 202, 219, 231, 246, 260, 275}

const _TokenKindLowerName = "tokenkindchartokenkinddottokenkindstartokenkindplustokenkindquestiontokenkindpipetokenkindescapetokenkindcarettokenkinddollartokenkindlparentokenkindrparentokenkindlbracetokenkindrbracetokenkindlbrackettokenkindrbrackettokenkindeoftokenkindnumbertokenkindcommatokenkindhyphen"

func (i TokenKind) String() string {
	if i < 0 || i >= TokenKind(len(_TokenKindIndex)-1) {
		return fmt.Sprintf("TokenKind(%d)", i)
	}
	return _TokenKindName[_TokenKindIndex[i]:_TokenKindIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TokenKindNoOp() {
	var x [1]struct{}
	_ = x[TokenKindChar-(0)]
	_ = x[TokenKindDot-(1)]
	_ = x[TokenKindStar-(2)]
	_ = x[TokenKindPlus-(3)]
	_ = x[TokenKindQuestion-(4)]
	_ = x[TokenKindPipe-(5)]
	_ = x[TokenKindEscape-(6)]
	_ = x[TokenKindCaret-(7)]
	_ = x[TokenKindDollar-(8)]
	_ = x[TokenKindLParen-(9)]
	_ = x[TokenKindRParen-(10)]
	_ = x[TokenKindLBrace-(11)]
	_ = x[TokenKindRBrace-(12)]
	_ = x[TokenKindLBracket-(13)]
	_ = x[TokenKindRBracket-(14)]
	_ = x[TokenKindEOF-(15)]
	_ = x[TokenKindNumber-(16)]
	_ = x[TokenKindComma-(17)]
	_ = x[TokenKindHyphen-(18)]
}

var _TokenKindValues = []TokenKind{TokenKindChar, TokenKindDot, TokenKindStar, TokenKindPlus, TokenKindQuestion, TokenKindPipe, TokenKindEscape, TokenKindCaret, TokenKindDollar, TokenKindLParen, TokenKindRParen, TokenKindLBrace, TokenKindRBrace, TokenKindLBracket, TokenKindRBracket, TokenKindEOF, TokenKindNumber, TokenKindComma, TokenKindHyphen}

var _TokenKindNameToValueMap = map[string]TokenKind{
	_TokenKindName[0:13]:         TokenKindChar,
	_TokenKindLowerName[0:13]:    TokenKindChar,
	_TokenKindName[13:25]:        TokenKindDot,
	_TokenKindLowerName[13:25]:   TokenKindDot,
	_TokenKindName[25:38]:        TokenKindStar,
	_TokenKindLowerName[25:38]:   TokenKindStar,
	_TokenKindName[38:51]:        TokenKindPlus,
	_TokenKindLowerName[38:51]:   TokenKindPlus,
	_TokenKindName[51:68]:        TokenKindQuestion,
	_TokenKindLowerName[51:68]:   TokenKindQuestion,
	_TokenKindName[68:81]:        TokenKindPipe,
	_TokenKindLowerName[68:81]:   TokenKindPipe,
	_TokenKindName[81:96]:        TokenKindEscape,
	_TokenKindLowerName[81:96]:   TokenKindEscape,
	_TokenKindName[96:110]:       TokenKindCaret,
	_TokenKindLowerName[96:110]:  TokenKindCaret,
	_TokenKindName[110:125]:      TokenKindDollar,
	_TokenKindLowerName[110:125]: TokenKindDollar,
	_TokenKindName[125:140]:      TokenKindLParen,
	_TokenKindLowerName[125:140]: TokenKindLParen,
	_TokenKindName[140:155]:      TokenKindRParen,
	_TokenKindLowerName[140:155]: TokenKindRParen,
	_TokenKindName[155:170]:      TokenKindLBrace,
	_TokenKindLowerName[155:170]: TokenKindLBrace,
	_TokenKindName[170:185]:      TokenKindRBrace,
	_TokenKindLowerName[170:185]: TokenKindRBrace,
	_TokenKindName[185:202]:      TokenKindLBracket,
	_TokenKindLowerName[185:202]: TokenKindLBracket,
	_TokenKindName[202:219]:      TokenKindRBracket,
	_TokenKindLowerName[202:219]: TokenKindRBracket,
	_TokenKindName[219:231]:      TokenKindEOF,
	_TokenKindLowerName[219:231]: TokenKindEOF,
	_TokenKindName[231:246]:      TokenKindNumber,
	_TokenKindLowerName[231:246]: TokenKindNumber,
	_TokenKindName[246:260]:      TokenKindComma,
	_TokenKindLowerName[246:260]: TokenKindComma,
	_TokenKindName[260:275]:      TokenKindHyphen,
	_TokenKindLowerName[260:275]: TokenKindHyphen,
}

var _TokenKindNames = []string{
	_TokenKindName[0:13],
	_TokenKindName[13:25],
	_TokenKindName[25:38],
	_TokenKindName[38:51],
	_TokenKindName[51:68],
	_TokenKindName[68:81],
	_TokenKindName[81:96],
	_TokenKindName[96:110],
	_TokenKindName[110:125],
	_TokenKindName[125:140],
	_TokenKindName[140:155],
	_TokenKindName[155:170],
	_TokenKindName[170:185],
	_TokenKindName[185:202],
	_TokenKindName[202:219],
	_TokenKindName[219:231],
	_TokenKindName[231:246],
	_TokenKindName[246:260],
	_TokenKindName[260:275],
}

// TokenKindString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TokenKindString(s string) (TokenKind, error) {
	if val, ok := _TokenKindNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TokenKindNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TokenKind values", s)
}

// TokenKindValues returns all values of the enum
func TokenKindValues() []TokenKind {
	return _TokenKindValues
}

// TokenKindStrings returns a slice of all String values of the enum
func TokenKindStrings() []string {
	strs := make([]string, len(_TokenKindNames))
	copy(strs, _TokenKindNames)
	return strs
}

// IsATokenKind returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TokenKind) IsATokenKind() bool {
	for _, v := range _TokenKindValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TokenKind
func (i TokenKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TokenKind
func (i *TokenKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TokenKind should be a string, got %s", data)
	}

	var err error
	*i, err = TokenKindString(s)
	return err
}
