
## Whay is Arrays in go ?
* Arrays are a collection of elements of the same type. 
* The elements of an array are indexed and can be accessed using the index of the element. 
* The index of an array starts from 0 and ends at n-1 where n is the number of elements in the array.
* 

## Array Syntax
```
var arrayName [size]type
```

## Multidimensional Array Syntax
```
var arrayName [size1][size2]...[sizeN]type
```

## make to initialize array
```
var a = make([]int, 5)
```


## Declaring an array

```
var a [3]int
```

##  Different ways to initialize an array

```
var a [3]int = [3]int{1, 2, 3}
var a = [3]int{1, 2, 3}
a := [3]int{1, 2, 3}
```

##  Accessing array elements

```
a := [3]int{1, 2, 3}
fmt.Println(a[0])
fmt.Println(a[1])
fmt.Println(a[2])
```

##  Iterating over an array

```
a := [3]int{1, 2, 3}
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}
```

## Range to iterate over an array

```
a := [3]int{1, 2, 3}
for i, v := range a {
    fmt.Println(i, v)
}
```
## Comparing arrays

```
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}

if a == b {
    fmt.Println("equal")
} else {
    fmt.Println("not equal")
}
```

##  Multidimensional arrays

```
a := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
fmt.Println(a)
```



