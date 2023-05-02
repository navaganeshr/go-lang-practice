
## Type Conversion

### What is type conversion?

Go is a statically typed language. This means that every variable has a type and that type cannot be changed. However, sometimes we need to convert a value from one type to another. This is called type conversion.


### How to convert a type?

- Convert int to float

```
var integer3 int = int(float4)
```

- Convert float to int

```
var float4 float64 = float64(integer3)
```

- Convert string to int

```
var integer5 int = strconv.Atoi(string6)
```

- Convert int to string

```
var string6 string = strconv.Itoa(integer5)
```

- Convert string to float

```
var float7 float64 = strconv.ParseFloat(string8, 64)
```

- Convert float to string

```
var string8 string = strconv.FormatFloat(float7, 'f', 6, 64)
```

- Convert string to bool

```
var bool9 bool = strconv.ParseBool(string10)
```

- Convert bool to string

```
var string10 string = strconv.FormatBool(bool9)
```




