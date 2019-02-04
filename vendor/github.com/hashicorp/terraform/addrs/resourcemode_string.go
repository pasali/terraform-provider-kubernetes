// Code generated by "stringer -type ResourceMode"; DO NOT EDIT.

package addrs

import "strconv"

const (
	_ResourceMode_name_0 = "InvalidResourceMode"
	_ResourceMode_name_1 = "DataResourceMode"
	_ResourceMode_name_2 = "ManagedResourceMode"
)

func (i ResourceMode) String() string {
	switch {
	case i == 0:
		return _ResourceMode_name_0
	case i == 68:
		return _ResourceMode_name_1
	case i == 77:
		return _ResourceMode_name_2
	default:
		return "ResourceMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}