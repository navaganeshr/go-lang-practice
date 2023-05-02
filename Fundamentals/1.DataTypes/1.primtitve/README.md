# Data Types

## Integer Data Type

### What is an integer?
integer is a whole number, positive or negative, without decimals, of unlimited length.

### How to declare an integer?
```
var integer1 int = 1
```

### How to declare an integer with a short declaration?
```
integer2 := 2
```

### Difference between int and int64?
int is a 32-bit integer and int64 is a 64-bit integer. int is the default type for integer values.

### Different integer types in Go
```
int8 -> -128 to 127
int16 -> -32768 to 32767
int32 -> -2147483648 to 2147483647
int64 -> -9223372036854775808 to 9223372036854775807
uint8 -> 0 to 255
uint16 -> 0 to 65535
uint32 -> 0 to 4294967295
```

### Examples
```
package main

import "fmt"

func main() {
	var integer_1 int = 1
	var integer_3 int
	integer_2 := 10
	integer_3 = 100
	fmt.Printf("integer_1: %v \n integer_2 %v \n integer_3 %v \n", integer_1, integer_2, integer_3)

}
```

## Float

Float is a data type that represents a floating point number. It is a 64-bit IEEE 754 floating point number.

### How to declare a float?

```
var float1 float64 = 1.1
```

### How to declare a float with a short declaration?

```
float2 := 2.2
```

### Difference between float and float64?
float is a 32-bit floating point number and float64 is a 64-bit floating point number. float is the default type for floating point values.

### Different float types in Go
```
float32 -> 1.18 × 10-38 to 3.40 × 1038
float64 -> 2.23 × 10-308 to 1.80 × 10308
```

### Example 
```
package main

import "fmt"

func main() {
	var float_1 float64 = 1.123
	var float_2 float64
	float_3 := 10.90
	float_2 = 100.345678
	fmt.Printf("integer_1: %v \n integer_2 %v \n integer_3 %v \n", float_1, float_2, float_3)
	fmt.Println("Sum of 3 numbers: ", sum(float_1, float_2, float_3))
	fmt.Println("Multiply of 3 numbers: ", multiply(float_1, float_2, float_3))

}

// sum of 3 numbers
func sum(a float64, b float64, c float64) float64 {
	return a + b + c
}

func multiply(a float64, b float64, c float64) float64 {
	return a * b * c
}
```

## Byte Data Type

### What is a byte?
byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.

### How to declare a byte?
```
var byte1 byte = 1
```

### How to declare a byte with a short declaration?
```
byte2 := 2
```

### Difference between byte and uint8?

### Example
```
package main

import "fmt"

func main() {
	Charstart := 65
	Charend := 90
	var someByte byte = 'A'
	fmt.Printf("value: %v \n byte: %c \t Type: %T \n", someByte, someByte, someByte)
	GenerateAlphabets(Charstart, Charend)
}

// Generate English alphabets
func GenerateAlphabets(charstart int, charend int) {
	// %c	the character represented by the corresponding Unicode code point
	// %v the value in a default format
	// %T	a Go-syntax representation of the type of the value
	for i := charstart; i <= charend; i++ {
		fmt.Printf("%v \t %c \t %T ", i, i, i)
		fmt.Println()
	}

}

```