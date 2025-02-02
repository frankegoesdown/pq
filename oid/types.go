// Code generated by gen.go. DO NOT EDIT.

package oid

const (
	ColTypeBoolean       Oid = 5
	ColTypeInt64         Oid = 6
	ColTypeFloat64       Oid = 7
	ColTypeChar          Oid = 8
	ColTypeVarChar       Oid = 9
	ColTypeTimestamp     Oid = 12
	ColTypeTimestampTZ   Oid = 13
	ColTypeVarBinary     Oid = 17
	ColTypeUUID          Oid = 20
	ColTypeLongVarChar   Oid = 115
	ColTypeLongVarBinary Oid = 116
	ColTypeBinary        Oid = 117
)

var TypeName = map[Oid]string{
	ColTypeBoolean:       "BOOL",
	ColTypeInt64:         "INT4",
	ColTypeFloat64:       "NUMERIC",
	ColTypeChar:          "CHAR",
	ColTypeVarChar:       "VARCHAR",
	ColTypeTimestamp:     "TIMESTAMP",
	ColTypeTimestampTZ:   "TIMESTAMPTZ",
	ColTypeVarBinary:     "VARBINARY",
	ColTypeUUID:          "UUID",
	ColTypeLongVarChar:   "LONGVARCHAR",
	ColTypeLongVarBinary: "LONGVARBINARY",
	ColTypeBinary:        "BINARY",
}
