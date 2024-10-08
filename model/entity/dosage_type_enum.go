// Code generated by "enumer -trimprefix=DosageType -type=DosageType -json -output dosage_type_enum.go"; DO NOT EDIT.

package entity

import (
	"encoding/json"
	"fmt"
)

const _DosageTypeName = "UnknownMg"

var _DosageTypeIndex = [...]uint8{0, 7, 9}

func (i DosageType) String() string {
	if i < 0 || i >= DosageType(len(_DosageTypeIndex)-1) {
		return fmt.Sprintf("DosageType(%d)", i)
	}
	return _DosageTypeName[_DosageTypeIndex[i]:_DosageTypeIndex[i+1]]
}

var _DosageTypeValues = []DosageType{0, 1}

var _DosageTypeNameToValueMap = map[string]DosageType{
	_DosageTypeName[0:7]: 0,
	_DosageTypeName[7:9]: 1,
}

// DosageTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DosageTypeString(s string) (DosageType, error) {
	if val, ok := _DosageTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DosageType values", s)
}

// DosageTypeValues returns all values of the enum
func DosageTypeValues() []DosageType {
	return _DosageTypeValues
}

// IsADosageType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DosageType) IsADosageType() bool {
	for _, v := range _DosageTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for DosageType
func (i DosageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for DosageType
func (i *DosageType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("DosageType should be a string, got %s", data)
	}

	var err error
	*i, err = DosageTypeString(s)
	return err
}
