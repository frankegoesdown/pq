package pq

import (
	"math"
	"reflect"
	"time"

	"github.com/frankegoesdown/pq/oid"
)

const headerSize = 4

type fieldDesc struct {
	// The object ID of the data type.
	OID oid.Oid
	// The data type size (see pg_type.typlen).
	// Note that negative values denote variable-width types.
	Len int
	// The type modifier (see pg_attribute.atttypmod).
	// The meaning of the modifier is type-specific.
	Mod int
}

func (fd fieldDesc) Type() reflect.Type {
	switch fd.OID {
	case oid.ColTypeInt64:
		return reflect.TypeOf(int64(0))
	case oid.ColTypeChar, oid.ColTypeVarChar, oid.ColTypeLongVarChar:
		return reflect.TypeOf("")
	case oid.ColTypeBoolean:
		return reflect.TypeOf(false)
	case oid.ColTypeTimestampTZ, oid.ColTypeTimestamp:
		return reflect.TypeOf(time.Time{})
	case oid.ColTypeBinary:
		return reflect.TypeOf([]byte(nil))
	default:
		return reflect.TypeOf(new(interface{})).Elem()
	}
}

func (fd fieldDesc) Name() string {
	return oid.TypeName[fd.OID]
}

func (fd fieldDesc) Length() (length int64, ok bool) {
	switch fd.OID {
	case oid.ColTypeLongVarChar, oid.ColTypeBinary:
		return math.MaxInt64, true
	case oid.ColTypeVarChar, oid.ColTypeChar:
		return int64(fd.Mod - headerSize), true
	default:
		return 0, false
	}
}

func (fd fieldDesc) PrecisionScale() (precision, scale int64, ok bool) {
	switch fd.OID {
	case oid.ColTypeFloat64:
		mod := fd.Mod - headerSize
		precision = int64((mod >> 16) & 0xffff)
		scale = int64(mod & 0xffff)
		return precision, scale, true
	default:
		return 0, 0, false
	}
}

// ColumnTypeScanType returns the value type that can be used to scan types into.
func (rs *rows) ColumnTypeScanType(index int) reflect.Type {
	return rs.colTyps[index].Type()
}

// ColumnTypeDatabaseTypeName return the database system type name.
func (rs *rows) ColumnTypeDatabaseTypeName(index int) string {
	return rs.colTyps[index].Name()
}

// ColumnTypeLength returns the length of the column type if the column is a
// variable length type. If the column is not a variable length type ok
// should return false.
func (rs *rows) ColumnTypeLength(index int) (length int64, ok bool) {
	return rs.colTyps[index].Length()
}

// ColumnTypePrecisionScale should return the precision and scale for decimal
// types. If not applicable, ok should be false.
func (rs *rows) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	return rs.colTyps[index].PrecisionScale()
}
