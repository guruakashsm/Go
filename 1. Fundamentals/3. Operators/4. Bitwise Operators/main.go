package main

import "fmt"

/*
In Go language, there are 6 bitwise operators that work at the bit level or are used to perform bit-by-bit operations. Following are the bitwise operators:

1. & (bitwise AND): Takes two numbers as operands and performs AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
2. | (bitwise OR): Takes two numbers as operands and performs OR on every bit of two numbers. The result of OR is 1 if any of the two bits is 1.
3. ^ (bitwise XOR): Takes two numbers as operands and performs XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
4. << (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
5. >> (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
6. &^ (AND NOT): This is a bit-clear operator.
*/

func main() {
	p := 34
	q := 20

	// & (bitwise AND)
	result1 := p & q
	fmt.Printf("Result of p & q = %d", result1)
	/*
		1. Bitwise AND (&)
		   The & operator performs a bitwise AND on two numbers. For each pair of corresponding bits in the two operands, the result will be 1 if both bits are 1. Otherwise, the result will be 0.

		Example:
		   p = 34, which is 00100010 in binary.
		   q = 20, which is 00010100 in binary.

		Operation:
		      p = 00100010  (34)
		      q = 00010100  (20)
		   -------------------------
		      p & q = 00000000  (0)
		   -------------------------

		Step-by-step:
		   1st bit: 0 & 0 = 0
		   2nd bit: 0 & 0 = 0
		   3rd bit: 1 & 0 = 0
		   4th bit: 0 & 1 = 0
		   5th bit: 0 & 0 = 0
		   6th bit: 0 & 1 = 0
		   7th bit: 0 & 0 = 0
		   8th bit: 1 & 0 = 0
		   So, result1 = 0.
	*/

	// | (bitwise OR)
	result2 := p | q
	fmt.Printf("\nResult of p | q = %d", result2)
	/*
		2. Bitwise OR (|)
		   The | operator performs a bitwise OR. For each corresponding pair of bits in the two operands, the result will be 1 if at least one of the bits is 1. If both bits are 0, the result is 0.

		Example:
		   p = 34, which is 00100010 in binary.
		   q = 20, which is 00010100 in binary.

		Operation:
		      p = 00100010  (34)
		      q = 00010100  (20)
		   -------------------------
		      p | q = 00110110  (54)
		   -------------------------

		Step-by-step:
		   1st bit: 0 | 0 = 0
		   2nd bit: 0 | 0 = 0
		   3rd bit: 1 | 0 = 1
		   4th bit: 0 | 1 = 1
		   5th bit: 0 | 0 = 0
		   6th bit: 0 | 1 = 1
		   7th bit: 0 | 0 = 0
		   8th bit: 1 | 0 = 1
		   So, result2 = 54.
	*/

	// ^ (bitwise XOR)
	result3 := p ^ q
	fmt.Printf("\nResult of p ^ q = %d", result3)
	/*
		3. Bitwise XOR (^)
		   The ^ operator performs a bitwise XOR (exclusive OR). For each pair of corresponding bits, the result will be 1 if the bits are different. If the bits are the same (both 0 or both 1), the result is 0.

		Example:
		   p = 34, which is 00100010 in binary.
		   q = 20, which is 00010100 in binary.

		Operation:
		      p = 00100010  (34)
		      q = 00010100  (20)
		   -------------------------
		      p ^ q = 00110110  (54)
		   -------------------------

		Step-by-step:
		   1st bit: 0 ^ 0 = 0
		   2nd bit: 0 ^ 0 = 0
		   3rd bit: 1 ^ 0 = 1
		   4th bit: 0 ^ 1 = 1
		   5th bit: 0 ^ 0 = 0
		   6th bit: 0 ^ 1 = 1
		   7th bit: 0 ^ 0 = 0
		   8th bit: 1 ^ 0 = 1
		   So, result3 = 54.
	*/

	// << (left shift)
	result4 := p << 1
	fmt.Printf("\nResult of p << 1 = %d", result4)
	/*
		4. Left Shift (<<)
		   The << operator performs a left shift on the bits of a number. It shifts all the bits to the left by a specified number of positions, and it fills with zeros from the right side. This operation is equivalent to multiplying the number by 2 for each shift.

		Example:
		   p = 34, which is 00100010 in binary.

		Operation:
		      p = 00100010  (34)
		      p << 1 = 01000100  (68)

		Step-by-step:
		   Shift all bits to the left by 1 position:
		   p = 00100010 becomes 01000100 (which is 68).
		   So, result4 = 68.
	*/

	// >> (right shift)
	result5 := p >> 1
	fmt.Printf("\nResult of p >> 1 = %d", result5)
	/*
		5. Right Shift (>>)
		   The >> operator performs a right shift on the bits of a number. It shifts all the bits to the right by a specified number of positions, and the leftmost bits are filled with the sign bit (in case of signed integers). For unsigned integers, they are filled with zeros.

		Example:
		   p = 34, which is 00100010 in binary.

		Operation:
		      p = 00100010  (34)
		      p >> 1 = 00010001  (17)

		Step-by-step:
		   Shift all bits to the right by 1 position:
		   p = 00100010 becomes 00010001 (which is 17).
		   So, result5 = 17.
	*/

	// &^ (AND NOT)
	result6 := p &^ q
	fmt.Printf("\nResult of p &^ q = %d", result6)
	/*
		6. AND NOT (&^)
		   The &^ operator performs a bitwise AND NOT. It works like a bitwise AND, but it sets the bits of the first operand to zero where the second operand has 1. In other words, it clears the bits of the first operand that correspond to 1 bits in the second operand.

		Example:
		   p = 34, which is 00100010 in binary.
		   q = 20, which is 00010100 in binary.

		Operation:
		      p = 00100010  (34)
		      q = 00010100  (20)
		      p &^ q = 00000010  (2)

		Step-by-step:
		   For each bit:
		   1st bit: 0 &^ 0 = 0
		   2nd bit: 0 &^ 0 = 0
		   3rd bit: 1 &^ 0 = 1
		   4th bit: 0 &^ 1 = 0
		   5th bit: 0 &^ 0 = 0
		   6th bit: 0 &^ 1 = 0
		   7th bit: 0 &^ 0 = 0
		   8th bit: 1 &^ 0 = 1
		   So, result6 = 2.
	*/
}
