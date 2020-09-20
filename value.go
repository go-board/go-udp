package goudp

import "fmt"

//go:generate stringer -type VType
// VType is value type
type VType int

const (
	// Undefined is unknown value type
	Undefined VType = iota
	// I64 is int type, can present int/int8/int16/int32/int64
	I64
	// U64 is uint type, can present uint/uint8/uint16/uint32/uint64
	U64
	// F64 is float type, can present float32/float64
	F64
	// Str is string type, can present string
	Str
	// Bool is bool type
	Bool
	// Array is array type, can present any of []*Value
	Array
	// Object is map type, can present any of map[string]*Value
	Object
	// UDT is user-defined-type.
	UDT
)

// Value represent an uniform go data value.
type Value struct {
	vType VType
	i     int64
	u     uint64
	f     float64
	s     string
	b     bool
	udt   interface{}       // user defined type
	arr   []*Value          // array of value
	obj   map[string]*Value // map of value
}

// NewUndefined create new undefined value.
func NewUndefined() *Value { return &Value{} }

// New create new value with given VType.
func New(vt VType) *Value { return &Value{vType: vt} }

func (v *Value) valid() bool { return v.vType != Undefined }

func (v *Value) release() {
	switch v.vType {
	case UDT:
		v.udt = nil
	case Array:
		v.arr = nil
	case Object:
		v.obj = nil
	}
}

// Type return value type.
func (v *Value) Type() VType { return v.vType }

// SetInt set int64 value, if inner vtype is not Undefined nor I64, return error.
func (v *Value) SetInt(i int64) error {
	if v.vType != Undefined && v.vType != I64 {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", I64, v.vType)
	}
	v.vType = I64
	v.i = i
	return nil
}

// SetUint set uint64 value, if inner vtype is not Undefined nor U64, return error.
func (v *Value) SetUint(u uint64) error {
	if v.vType != Undefined && v.vType != U64 {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", U64, v.vType)
	}
	v.vType = U64
	v.u = u
	return nil
}

// SetFloat set float64 value, if inner vtype is not Undefined nor F64, return error.
func (v *Value) SetFloat(f float64) error {
	if v.vType != Undefined && v.vType != F64 {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", F64, v.vType)
	}
	v.vType = F64
	v.f = f
	return nil
}

// SetString set string value, if inner vtype is not Undefined nor Str, return error.
func (v *Value) SetString(s string) error {
	if v.vType != Undefined && v.vType != Str {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", Str, v.vType)
	}
	v.vType = Str
	v.s = s
	return nil
}

// SetBool set bool value, if inner vtype is not Undefined nor Bool, return error.
func (v *Value) SetBool(b bool) error {
	if v.vType != Undefined && v.vType != Bool {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", Bool, v.vType)
	}
	v.vType = Bool
	v.b = b
	return nil
}

// SetUdt set user-defined-type value, if inner vtype is not Undefined nor UDT, return error.
func (v *Value) SetUdt(vType int, udt interface{}) error {
	if v.vType != Undefined && v.vType != UDT {
		return fmt.Errorf("err: bad vtype, want %s, but got %s", UDT, v.vType)
	}
	v.vType = UDT
	v.udt = udt
	return nil
}

// SetArray return element i of inner array, if inner vtype is not Undefined nor Array, return error.
func (v *Value) SetArray(i int) (*Value, error) {
	if i < 0 {
		return nil, fmt.Errorf("err: bad index, got %d", i)
	}
	if v.vType != Undefined && v.vType != Array {
		return nil, fmt.Errorf("err: bad vtype, want %s, but got %s", Array, v.vType)
	}
	if len(v.arr) < i {
		old := v.arr
		v.arr = make([]*Value, i+1)
		copy(v.arr, old)
	}
	tmp := NewUndefined()
	v.arr[i] = tmp
	return tmp, nil
}

// SetObject return key of inner object, if inner vtype is not Undefined nor Object, return error.
func (v *Value) SetObject(key string) (*Value, error) {
	if v.vType != Undefined && v.vType != Object {

	}
	v.vType = Object
	if v.obj == nil {
		v.obj = make(map[string]*Value)
	}
	tmp := NewUndefined()
	v.obj[key] = tmp
	return tmp, nil
}
