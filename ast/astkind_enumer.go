// Code generated by "enumer -type=ASTKind -json"; DO NOT EDIT.

package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _ASTKindName = "ASTKindCharASTKindUnionASTKindConcatASTKindStarASTKindGroup"

var _ASTKindIndex = [...]uint8{0, 11, 23, 36, 47, 59}

const _ASTKindLowerName = "astkindcharastkindunionastkindconcatastkindstarastkindgroup"

func (i ASTKind) String() string {
	if i < 0 || i >= ASTKind(len(_ASTKindIndex)-1) {
		return fmt.Sprintf("ASTKind(%d)", i)
	}
	return _ASTKindName[_ASTKindIndex[i]:_ASTKindIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ASTKindNoOp() {
	var x [1]struct{}
	_ = x[ASTKindChar-(0)]
	_ = x[ASTKindUnion-(1)]
	_ = x[ASTKindConcat-(2)]
	_ = x[ASTKindStar-(3)]
	_ = x[ASTKindGroup-(4)]
}

var _ASTKindValues = []ASTKind{ASTKindChar, ASTKindUnion, ASTKindConcat, ASTKindStar, ASTKindGroup}

var _ASTKindNameToValueMap = map[string]ASTKind{
	_ASTKindName[0:11]:       ASTKindChar,
	_ASTKindLowerName[0:11]:  ASTKindChar,
	_ASTKindName[11:23]:      ASTKindUnion,
	_ASTKindLowerName[11:23]: ASTKindUnion,
	_ASTKindName[23:36]:      ASTKindConcat,
	_ASTKindLowerName[23:36]: ASTKindConcat,
	_ASTKindName[36:47]:      ASTKindStar,
	_ASTKindLowerName[36:47]: ASTKindStar,
	_ASTKindName[47:59]:      ASTKindGroup,
	_ASTKindLowerName[47:59]: ASTKindGroup,
}

var _ASTKindNames = []string{
	_ASTKindName[0:11],
	_ASTKindName[11:23],
	_ASTKindName[23:36],
	_ASTKindName[36:47],
	_ASTKindName[47:59],
}

// ASTKindString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ASTKindString(s string) (ASTKind, error) {
	if val, ok := _ASTKindNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ASTKindNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ASTKind values", s)
}

// ASTKindValues returns all values of the enum
func ASTKindValues() []ASTKind {
	return _ASTKindValues
}

// ASTKindStrings returns a slice of all String values of the enum
func ASTKindStrings() []string {
	strs := make([]string, len(_ASTKindNames))
	copy(strs, _ASTKindNames)
	return strs
}

// IsAASTKind returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ASTKind) IsAASTKind() bool {
	for _, v := range _ASTKindValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ASTKind
func (i ASTKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ASTKind
func (i *ASTKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ASTKind should be a string, got %s", data)
	}

	var err error
	*i, err = ASTKindString(s)
	return err
}
