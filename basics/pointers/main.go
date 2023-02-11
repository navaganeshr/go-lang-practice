package main

import "fmt"

func main() {
	/*
		    	* 1. A pointer is a variable that stores the memory address of another variable.
					* pointer syntax:
					var <name> *<type> = &<variable>
				* 2. The & operator is used to get the memory address of a variable.
				* 3. The * operator is used to get the value of a variable that is stored in a pointer.
				* 4. The * operator is also used to declare a pointer.
				* 5. How to deference a pointer?
					 <pointer> = *<pointer>
				* 6. deference a pointer: *<pointer>

	*/

	var name string = "John"
	var ptr_1 *string = &name
	fmt.Printf("the pointer value of \"name\" is: %v and the same for pointer is %v \n", &name, ptr_1)
	fmt.Printf("The value of ptr_1 is: %v and type is %T \n", name, ptr_1)
	fmt.Printf("Derefernced value of pointer ptr_1 is: %v \n", *ptr_1)

	/*
		 passing pointer to function syntax:
		 func <function_name> (<pointer_name> *<type>) {
		  body of the function
		}
	*/

	fmt.Printf("The value of name before calling the function is: %v \n", *ptr_1)
	pointerExample(&name)
	fmt.Printf("The value of name after calling the function is: %v \n", *ptr_1)

}

func pointerExample(ptr_2 *string) {
	*ptr_2 = "Doe"
}
