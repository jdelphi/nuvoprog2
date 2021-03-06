// Code generated by "enumer -type=BootSelect -trimprefix=BootFrom -transform=snake -json -text"; DO NOT EDIT

package ms51

import (
	"encoding/json"
	"fmt"
)

const _BootSelectName = "ldromaprom"

var _BootSelectIndex = [...]uint8{0, 5, 10}

func (i BootSelect) String() string {
	if i < 0 || i >= BootSelect(len(_BootSelectIndex)-1) {
		return fmt.Sprintf("BootSelect(%d)", i)
	}
	return _BootSelectName[_BootSelectIndex[i]:_BootSelectIndex[i+1]]
}

var _BootSelectValues = []BootSelect{0, 1}

var _BootSelectNameToValueMap = map[string]BootSelect{
	_BootSelectName[0:5]:  0,
	_BootSelectName[5:10]: 1,
}

// BootSelectString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func BootSelectString(s string) (BootSelect, error) {
	if val, ok := _BootSelectNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to BootSelect values", s)
}

// BootSelectValues returns all values of the enum
func BootSelectValues() []BootSelect {
	return _BootSelectValues
}

// IsABootSelect returns "true" if the value is listed in the enum definition. "false" otherwise
func (i BootSelect) IsABootSelect() bool {
	for _, v := range _BootSelectValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for BootSelect
func (i BootSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BootSelect
func (i *BootSelect) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BootSelect should be a string, got %s", data)
	}

	var err error
	*i, err = BootSelectString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for BootSelect
func (i BootSelect) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for BootSelect
func (i *BootSelect) UnmarshalText(text []byte) error {
	var err error
	*i, err = BootSelectString(string(text))
	return err
}
