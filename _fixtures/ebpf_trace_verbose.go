package main

import "fmt"

// Test function with primitives for verbosity testing
// Only uses types supported by eBPF: int, uint, bool, string, byte
//
//go:noinline
func testPrimitives(i int64, u uint64, b bool, s string) int64 {
	fmt.Println(i, u, b, s)
	return i * 2
}

// Test function with actual small integer types (newly supported)
//
//go:noinline
func testIntegerTypes(i8 int8, i16 int16, i32 int32, i64 int64) int64 {
	return int64(i8) + int64(i16) + int64(i32) + i64
}

// Test function with actual unsigned small integer types (newly supported)
//
//go:noinline
func testUnsignedTypes(u8 uint8, u16 uint16, u32 uint32, u64 uint64) uint64 {
	return uint64(u8) + uint64(u16) + uint64(u32) + u64
}

// Test function with boolean and byte
//
//go:noinline
func testBoolByte(flag bool, ch byte) bool {
	fmt.Println(flag, ch)
	return !flag
}

// Test function with string parameter
//
//go:noinline
func testString(message string) int64 {
	fmt.Println(message)
	return int64(len(message))
}

// Test function with pointer types (newly supported)
//
//go:noinline
func testPointers(p *int, u *uint64) int64 {
	if p != nil && u != nil {
		fmt.Println(*p, *u)
		return int64(*p) + int64(*u)
	}
	return 0
}

// Test function with slice types (newly supported)
//
//go:noinline
func testSlices(s []byte, nums []int32) int {
	fmt.Println(s, nums)
	return len(s) + len(nums)
}

// Test function with mixed types for comprehensive verbosity testing
// Now includes newly supported types: small ints, pointers, slices
//
//go:noinline
func testMixed(i8 int8, u16 uint16, label string, p *int) string {
	fmt.Println(i8, u16, label, p)
	return label
}

func main() {
	// Test 1: Primitives
	testPrimitives(42, 100, true, "Hello eBPF")

	// Test 2: Small integer types (newly supported)
	testIntegerTypes(8, 16, 32, 64)

	// Test 3: Small unsigned types (newly supported)
	testUnsignedTypes(8, 16, 32, 64)

	// Test 4: Bool and byte
	testBoolByte(true, 'A')

	// Test 5: String
	testString("eBPF Trace Verbosity Test")

	// Test 6: Pointer types (newly supported)
	x := 42
	y := uint64(100)
	testPointers(&x, &y)

	// Test 7: Slice types (newly supported)
	testSlices([]byte{1, 2, 3}, []int32{10, 20, 30})

	// Test 8: Mixed types with new type support
	z := 99
	testMixed(7, 1000, "mixed", &z)
}
