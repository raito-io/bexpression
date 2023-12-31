// Code generated by "enumer -type=EntityType -values -gqlgen -yaml -json -trimprefix=EntityType"; DO NOT EDIT.

package datacomparison

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const _EntityTypeName = "DataObjectColumnReferenceByName"

var _EntityTypeIndex = [...]uint8{0, 10, 31}

const _EntityTypeLowerName = "dataobjectcolumnreferencebyname"

func (i EntityType) String() string {
	if i < 0 || i >= EntityType(len(_EntityTypeIndex)-1) {
		return fmt.Sprintf("EntityType(%d)", i)
	}
	return _EntityTypeName[_EntityTypeIndex[i]:_EntityTypeIndex[i+1]]
}

func (EntityType) Values() []string {
	return EntityTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EntityTypeNoOp() {
	var x [1]struct{}
	_ = x[EntityTypeDataObject-(0)]
	_ = x[EntityTypeColumnReferenceByName-(1)]
}

var _EntityTypeValues = []EntityType{EntityTypeDataObject, EntityTypeColumnReferenceByName}

var _EntityTypeNameToValueMap = map[string]EntityType{
	_EntityTypeName[0:10]:       EntityTypeDataObject,
	_EntityTypeLowerName[0:10]:  EntityTypeDataObject,
	_EntityTypeName[10:31]:      EntityTypeColumnReferenceByName,
	_EntityTypeLowerName[10:31]: EntityTypeColumnReferenceByName,
}

var _EntityTypeNames = []string{
	_EntityTypeName[0:10],
	_EntityTypeName[10:31],
}

// EntityTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EntityTypeString(s string) (EntityType, error) {
	if val, ok := _EntityTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EntityTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EntityType values", s)
}

// EntityTypeValues returns all values of the enum
func EntityTypeValues() []EntityType {
	return _EntityTypeValues
}

// EntityTypeStrings returns a slice of all String values of the enum
func EntityTypeStrings() []string {
	strs := make([]string, len(_EntityTypeNames))
	copy(strs, _EntityTypeNames)
	return strs
}

// IsAEntityType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EntityType) IsAEntityType() bool {
	for _, v := range _EntityTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for EntityType
func (i EntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EntityType
func (i *EntityType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EntityType should be a string, got %s", data)
	}

	var err error
	*i, err = EntityTypeString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for EntityType
func (i EntityType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EntityType
func (i *EntityType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = EntityTypeString(s)
	return err
}

// MarshalGQL implements the graphql.Marshaler interface for EntityType
func (i EntityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(i.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for EntityType
func (i *EntityType) UnmarshalGQL(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("EntityType should be a string, got %T", value)
	}

	var err error
	*i, err = EntityTypeString(str)
	return err
}
