# Control Structures
## Table of Contents
- [Control Structures](#control-structures)
  - [Table of Contents](#table-of-contents)
  - [Blocks, Shadowing, Redeclaration, Reinitialization, Reassignment in go](#blocks-shadowing-redeclaration-reinitialization-reassignment-in-go)
    - [What is blocks in go ?](#what-is-blocks-in-go-)
    - [what shadowing of variables in go](#what-shadowing-of-variables-in-go)
    - [What is redeclaration of variables in go](#what-is-redeclaration-of-variables-in-go)
    - [What is reinitialization of variables in go?](#what-is-reinitialization-of-variables-in-go)
    - [What is reassignment of variables in go?](#what-is-reassignment-of-variables-in-go)
    - [what is the difference between shadowing and redeclaration](#what-is-the-difference-between-shadowing-and-redeclaration)
    - [what is the difference between shadowing and reassignment](#what-is-the-difference-between-shadowing-and-reassignment)
    - [what is the difference between shadowing and reinitialization](#what-is-the-difference-between-shadowing-and-reinitialization)
  - [if or if else loops in go](#if-or-if-else-loops-in-go)
    - [syntax of if loops in go](#syntax-of-if-loops-in-go)
    - [example of if loops in go](#example-of-if-loops-in-go)
    - [Different syntax of if loops in go](#different-syntax-of-if-loops-in-go)
    - [if with multiple conditions in go](#if-with-multiple-conditions-in-go)
    - [Scoping a variable to an if statement in go](#scoping-a-variable-to-an-if-statement-in-go)
  - [switch loops in go](#switch-loops-in-go)
    - [syntax of switch loops in go](#syntax-of-switch-loops-in-go)
    - [example of switch loops in go](#example-of-switch-loops-in-go)
    - [switch with multiple conditions in go](#switch-with-multiple-conditions-in-go)
    - [switch with no condition in go](#switch-with-no-condition-in-go)
    - [what is fallthrough in go ?](#what-is-fallthrough-in-go-)
    - [switch with fallthrough in go](#switch-with-fallthrough-in-go)
    - [switch with multiple conditions and fallthrough in go](#switch-with-multiple-conditions-and-fallthrough-in-go)
    - [switch with no condition and fallthrough in go](#switch-with-no-condition-and-fallthrough-in-go)
  - [for loops in go](#for-loops-in-go)
    - [syntax of for loops in go](#syntax-of-for-loops-in-go)
    - [example of for loops in go](#example-of-for-loops-in-go)
    - [for with multiple conditions in go](#for-with-multiple-conditions-in-go)
    - [for with no condition in go](#for-with-no-condition-in-go)
    - [syntax of for loops with no condition in go](#syntax-of-for-loops-with-no-condition-in-go)
    - [for loop with break in go](#for-loop-with-break-in-go)
    - [Difference between break and return in go](#difference-between-break-and-return-in-go)
    - [for loop with continue in go](#for-loop-with-continue-in-go)
    - [label in go ?](#label-in-go-)
    - [syntax of label in go](#syntax-of-label-in-go)
    - [for loop with goto in go](#for-loop-with-goto-in-go)





## Blocks, Shadowing, Redeclaration, Reinitialization, Reassignment in go

### What is blocks in go ?
- Blocks are a way to group multiple statements into a single unit.
-  Variables,constants, types, and functions declared outside of any functions are placed in the package block
-  Within a function, every set of  braces ({}) defines another block
-  

### what shadowing of variables in go

- Shadowing is the process of declaring a new variable with the same name as an existing variable in the same scope.


### What is redeclaration of variables in go
* Redeclaration does not introduce a new variable. 
* It just assigns a new value to the original.

example:
```
package main

import "fmt"

x := 10
x := 20 // redeclaration
fmt.Println(x)
```
### What is reinitialization of variables in go?
reinitialization is the process of assigning a new value to an existing variable in the same scope.
example:
```
package main
import "fmt"

i := 10
i := 20 // reinitialization ( this is not allowed in GO)
fmt.Println(i)
```


### What is reassignment of variables in go?
reassignment is the process of assigning a new value to an existing variable in the same scope.
example:
```
package main
import "fmt"

x := 10
x = 20 // reassignment
fmt.Println(x)
```
Note:  reassignment and reinitialization are similar but reinitialization is not allowed in go.


### what is the difference between shadowing and redeclaration
Difference between shadowing and redeclaration is that shadowing is allowed in go but redeclaration is not allowed in go.

### what is the difference between shadowing and reassignment
Difference between shadowing and reassignment is that shadowing is allowed in go but reassignment is not allowed in go.

### what is the difference between shadowing and reinitialization
Difference between shadowing and reinitialization is that shadowing is allowed in go but reinitialization is not allowed in go.


## if or if else loops in go
if loops are used to execute a block of code if a condition is true.

### syntax of if loops in go
```
if condition {
    // code block
}
```
### example of if loops in go
```
package main
import "fmt"

func main() {
    x := 10
    if x == 10 {
        fmt.Println("x is equal to 10")
    }
}
```

### Different syntax of if loops in go
```
if condition {
    // code block
} else {
    // code block
}
```


```
if condition1 {
    // code block
} else if condition2 {
    // code block
} else {
    // code block
}
```

```
if condition1 {
    // code block
} else if condition2 {
    // code block
} else if condition3 {
    // code block
} else if condition4 {
    // code block
} else {
    // code block
}
```

### if with multiple conditions in go
```
if condition1 && condition2 {
    // code block
}
```
### Scoping a variable to an if statement in go
```
package main
import (
    "fmt"
    "math/rand"
)
func main() {
    if n := rand.Intn(10); n == 0 {
        fmt.Println("That's too low")
    } else if n > 5 {
        fmt.Println("That's too big:", n)
    } else {
        fmt.Println("That's a good number:", n)
    }
}
```

## switch loops in go
switch loops are used to execute a block of code based on the value of a variable.

### syntax of switch loops in go
```
switch variable {
    case value1:
        // code block
    case value2:
        // code block
    case value3:
        // code block
    default:
        // code block
}
```
### example of switch loops in go
```
package main
import "fmt"

city := "Delhi"

switch city {
    case "Delhi":
        fmt.Println("Capital of India")
    case "Mumbai":
        fmt.Println("Financial Capital of India")
    case "Bangalore":
        fmt.Println("Silicon Valley of India")
    default:
        fmt.Println("No information available for this city")
}
```

### switch with multiple conditions in go
```
switch variable {
    case value1, value2, value3:
        // code block
    default:
        // code block
}
```

### switch with no condition in go
```
switch {
    case variable1 == value1:
        // code block
    case variable2 == value2:
        // code block
    default:
        // code block
}
```

### what is fallthrough in go ?

fallthrough is used to execute the next case statement in a switch statement.

### switch with fallthrough in go
```
switch variable {
    case value1:
        // code block
        fallthrough
    case value2:
        // code block
    default:
        // code block
}
```

### switch with multiple conditions and fallthrough in go
```
switch variable {
    case value1, value2, value3:
        // code block
        fallthrough
    default:
        // code block
}
```

### switch with no condition and fallthrough in go
```
switch {
    case variable1 == value1:
        // code block
        fallthrough
    case variable2 == value2:
        // code block
    default:
        // code block
}
```

## for loops in go
for loops are used to execute a block of code multiple times.

### syntax of for loops in go
```
for init; condition; post {
    // code block
}
```

### example of for loops in go
```
package main
import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

### for with multiple conditions in go
```
for condition1 && condition2 {
    // code block
}
```

### for with no condition in go

for loops with no condition are used to execute a block of code infinitely.

### syntax of for loops with no condition in go
```
for {
    // code block
}
```

### for loop with break in go
* break is used to exit a for loop.
* break can be used with for loops with or without a condition.

```
for {
    // code block
    break
}
```

### Difference between break and return in go
* break is used to exit a for loop.
* return is used to exit a function.



### for loop with continue in go
* continue is used to skip the current iteration of a for loop.

syntax of for loop with continue in go
```
for {
    // code block
    if condition {
        continue
    }
```
NOTE: 
    * break & continue can be used with for loops with or without a condition.
    * 

### label in go ?
label is a keyword in go which is used to define a label in a function.
* usecase of label ?
  


### syntax of label in go
```
label:
```

### for loop with goto in go
* goto is used to jump to a label in a function.

syntax of for loop with goto in go
```
for {
    // code block
    if condition {
        goto label
    }
}
label:
* labels are used to define a label in a function.
* goto is used to jump to a label in a function.
* goto can be used with for loops with or without a condition.
* goto can be used with if loops with or without a condition.
* goto can be used with switch loops with or without a condition.

### example of for loop with goto in go
```
package main
import "fmt"

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        if i == 10 {
            goto end
        }
    }

end:
    fmt.Println("end")

}
```
 







