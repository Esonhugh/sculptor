package pkg

/*
import "reflect"

// ValueSelector is the class for selecting values in structs
type ValueSelector struct {
	// ColumnName means How to select the data in files
	// such as "ID" in "raw data" you import means "身份证" in basic info "type" field
	// and put the "ID" or the json query in this "ColumnName" field
	ColumnName string
	// MapNameInBasicInfo means How to name selected data in files
	// such as "ID" in "raw data" you import means "身份证" in basic info "type" field
	// and put the "Info_TYPE_ID" in this "MapNameInBasicInfo" field
	MapNameInBasicInfo string
	// Index of the columnName in the file This is useful in csv and xls data.
	Index int // Index of the columnName in the file This is useful in csv and xls data.
	// value is Selector selected Value
	value reflect.Value
}

// Selector is the class for selectors
type Selector []ValueSelector

func (v *ValueSelector) SetValue(value string) {
	v.value.SetString(value)
}

func (v *ValueSelector) GetValue() string {
	return v.value.String()
}
*/
