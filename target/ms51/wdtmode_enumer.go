// Code generated by "enumer -type=WDTMode -trimprefix=WDT -transform=snake -json -text"; DO NOT EDIT

package ms51

import (
	"encoding/json"
	"fmt"
)

const _WDTModeName = "disabledenabledenabled_always"

var _WDTModeIndex = [...]uint8{0, 8, 15, 29}

func (i WDTMode) String() string {
	if i >= WDTMode(len(_WDTModeIndex)-1) {
		return fmt.Sprintf("WDTMode(%d)", i)
	}
	return _WDTModeName[_WDTModeIndex[i]:_WDTModeIndex[i+1]]
}

var _WDTModeValues = []WDTMode{0, 1, 2}

var _WDTModeNameToValueMap = map[string]WDTMode{
	_WDTModeName[0:8]:   0,
	_WDTModeName[8:15]:  1,
	_WDTModeName[15:29]: 2,
}

// WDTModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func WDTModeString(s string) (WDTMode, error) {
	if val, ok := _WDTModeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to WDTMode values", s)
}

// WDTModeValues returns all values of the enum
func WDTModeValues() []WDTMode {
	return _WDTModeValues
}

// IsAWDTMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i WDTMode) IsAWDTMode() bool {
	for _, v := range _WDTModeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for WDTMode
func (i WDTMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for WDTMode
func (i *WDTMode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WDTMode should be a string, got %s", data)
	}

	var err error
	*i, err = WDTModeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for WDTMode
func (i WDTMode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for WDTMode
func (i *WDTMode) UnmarshalText(text []byte) error {
	var err error
	*i, err = WDTModeString(string(text))
	return err
}
