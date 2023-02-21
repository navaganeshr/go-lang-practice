package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to string
	var integer0 int = 42
	var string0 string = strconv.Itoa(integer0)
	fmt.Println("Convert int to string:")
	fmt.Printf("%v, %T \n", string0, string0)

	// string to int
	var string1 string = "42"
	integer1, err := strconv.Atoi(string1)
	fmt.Println("Convert string to int:")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v, %T \n", integer1, integer1)

	// float to string
	var float1 float64 = 42.42
	var float_string string = strconv.FormatFloat(float1, 'f', 2, 64)
	fmt.Println("Convert float to string:")
	fmt.Printf("%s, %T \n", float_string, float_string)

	// string to float
	var float_string1 string = "42.42"
	float2, err := strconv.ParseFloat(float_string1, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Convert string to float:")
	fmt.Printf("%v, %T \n", float2, float2)

	// int to float
	var integer2 int = 42
	var float3 float64 = float64(integer2)
	fmt.Println("Convert int to float:")
	fmt.Printf("%v, %T \n", float3, float3)

	// float to int
	var float4 float64 = 42.42
	var integer3 int = int(float4)
	fmt.Println("Convert float to int:")
	fmt.Printf("%v, %T \n", integer3, integer3)

}
