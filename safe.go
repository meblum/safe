package safe

import "math"

type Int struct{}

func (i Int) Add(x, y int) (int, bool) {
	r := x + y
	// Overflow if both arguments have the opposite sign of the result
	if ((x ^ r) & (y ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int) Subtract(x, y int) (int, bool) {
	r := x - y
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if ((x ^ y) & (x ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int) Muliply(x, y int) (int, bool) {
	r := x * y
	ax := abs(int64(x))
	ay := abs(int64(y))
	if (ax|ay)>>31 != 0 {
		// Some bits greater than 2^31 that might cause overflow
		// Check the result using the divide operator
		// and check for the special case of Long.MIN_VALUE * -1
		if ((y != 0) && (r/y != x)) ||
			(x == math.MinInt64 && y == -1) {
			return r, false
		}
	}
	return r, true
}

func (i Int) Increment(v int) (int, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Int) Deccrement(v int) (int, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func (i Int) Negate(v int) (int, bool) {
	r := -v
	if v == math.MinInt64 || -int64(v) != int64(r) {
		return r, false
	}

	return r, true
}

type Int32 struct{}

func (i Int32) Add(x, y int32) (int32, bool) {
	r := x + y
	// Overflow if both arguments have the opposite sign of the result
	if ((x ^ r) & (y ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int32) Subtract(x, y int32) (int32, bool) {
	r := x - y
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if ((x ^ y) & (x ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int32) Muliply(x, y int32) (int32, bool) {
	r := int64(x) * int64(y)
	if int64(int32(r)) != r {
		return int32(r), false
	}
	return int32(r), true
}

func (i Int32) Increment(v int32) (int32, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Int32) Deccrement(v int32) (int32, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func (i Int32) Negate(v int32) (int32, bool) {
	r := -v
	if v == math.MinInt32 {
		return r, false
	}

	return r, true
}

type Int64 struct{}

func (i Int64) Add(x, y int64) (int64, bool) {
	r := x + y
	// Overflow if both arguments have the opposite sign of the result
	if ((x ^ r) & (y ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int64) Subtract(x, y int64) (int64, bool) {
	r := x - y
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if ((x ^ y) & (x ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int64) Muliply(x, y int64) (int64, bool) {
	r := x * y
	ax := abs(x)
	ay := abs(y)
	if (ax|ay)>>31 != 0 {
		// Some bits greater than 2^31 that might cause overflow
		// Check the result using the divide operator
		// and check for the special case of Long.MIN_VALUE * -1
		if ((y != 0) && (r/y != x)) ||
			(x == math.MinInt64 && y == -1) {
			return r, false
		}
	}
	return r, true
}

func (i Int64) Increment(v int64) (int64, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Int64) Deccrement(v int64) (int64, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func (i Int64) Negate(v int64) (int64, bool) {
	r := -v
	if v == math.MinInt64 {
		return r, false
	}

	return r, true
}

type Int16 struct{}

func (i Int16) Add(x, y int16) (int16, bool) {
	r := x + y
	// Overflow if both arguments have the opposite sign of the result
	if ((x ^ r) & (y ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int16) Subtract(x, y int16) (int16, bool) {
	r := x - y
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if ((x ^ y) & (x ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int16) Muliply(x, y int16) (int16, bool) {
	r := int64(x) * int64(y)
	if int64(int16(r)) != r {
		return int16(r), false
	}
	return int16(r), true
}

func (i Int16) Increment(v int16) (int16, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Int16) Deccrement(v int16) (int16, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func (i Int16) Negate(v int16) (int16, bool) {
	r := -v
	if v == math.MinInt16 {
		return r, false
	}

	return r, true
}

type Int8 struct{}

func (i Int8) Add(x, y int8) (int8, bool) {
	r := x + y
	// Overflow if both arguments have the opposite sign of the result
	if ((x ^ r) & (y ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int8) Subtract(x, y int8) (int8, bool) {
	r := x - y
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if ((x ^ y) & (x ^ r)) < 0 {
		return r, false
	}
	return r, true
}

func (i Int8) Muliply(x, y int8) (int8, bool) {
	r := int64(x) * int64(y)
	if int64(int8(r)) != r {
		return int8(r), false
	}
	return int8(r), true
}

func (i Int8) Increment(v int8) (int8, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Int8) Deccrement(v int8) (int8, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func (i Int8) Negate(v int8) (int8, bool) {
	r := -v
	if v == math.MinInt8 {
		return r, false
	}

	return r, true
}

type Byte struct{}

func (i Byte) Add(x, y byte) (byte, bool) {

	r := x + y

	ir := int(x) + int(y)
	// Overflow if both arguments have the opposite sign of the result
	if int(r) != ir {
		return r, false
	}
	return r, true
}

func (i Byte) Subtract(x, y byte) (byte, bool) {
	r := x - y
	ir := int(x) - int(y)
	// Overflow iff the arguments have different signs and
	// the sign of the result is different than the sign of x
	if int(r) != ir {
		return r, false
	}
	return r, true
}

func (i Byte) Muliply(x, y byte) (byte, bool) {
	r := int64(x) * int64(y)
	if int64(byte(r)) != r {
		return byte(r), false
	}
	return byte(r), true
}

func (i Byte) Increment(v byte) (byte, bool) {
	r := v + 1
	if r < v {
		return r, false
	}

	return r, true
}

func (i Byte) Deccrement(v int) (int, bool) {
	r := v - 1
	if r > v {
		return r, false
	}

	return r, true
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
