
## Blocks, Shadowing, Redeclaration, Reinitialization, Reassignment in go

### What is blocks in go ?
- Blocks are a way to group multiple statements into a single unit.
-  Variables,constants, types, and functions declared outside of any functions are placed in the package block
-  Within a function, every set of  braces ({}) defines another block
-  

### what shadowing of variables in go

- Shadowing is the process of declaring a new variable with the same name as an existing variable in the same scope.


### What is redeclaration of variables in go
redeclaration is the process of declaring a new variable with the same name as an existing variable in the same scope.
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





