## Struct in go

### What is Struct in go ?
- Struct is a collection of fields. 
- Structs are similar to classes in other languages. 
- Structs are used to represent a record.
- Structs can contain fields which can be any type including other structs.
- Structs can also contain methods which are functions with a special receiver argument. 
- Structs are useful for grouping data together to form records.
- Its a blueprint of a data type.
- Structs types are created during compile time.
- Field name and types are declared during compile time.
- Field types cannot be changed once declared.
- Field values can be changed during runtime.
- Structs cannot dynamically grow or shrink. but they have different sset of types 

## Features of Struct in go
- Structs are value types.
- Cannot add new fields to a struct once it is defined (runtime also).

## How to declare a struct in go ?
```
type structName struct {
    field1 type1
    field2 type2
    ...
    fieldN typeN
}

Example: 

type person struct {
    name string
    age int
}
```

## How to declare a struct with a short declaration ?
```
structName := struct {
    field1 type1
    field2 type2
    ...
    fieldN typeN
}{
    field1: value1,
    field2: value2,
    ...
    fieldN: valueN,
}

Example:
person := struct {
    name string
    age int
}{
    name: "NavaGanesh",
    age: 30,
}
```

## How to access struct fields ?
```
structName.fieldName 

Example:
person.name
person.age
```


## Ananymous Struct
- Ananymous struct is a struct without a name.

syntax:
```
struct {
    field1 type1
    field2 type2
    ...
    fieldN typeN
}
```

use case:
- Ananymous struct is used to create a struct without a name.

Example:
```
person := struct {
    name string
    age int
}{
    name: "NavaGanesh",
    age: 30,
}
```

## Comparing structs
- Two structs are equal if they have the same type and if their corresponding fields are equal.
  
 ```
    type person struct {
        name string
        age int 
    }

    p1 := person{
        name: "NavaGanesh",
        age: 30,
    }

    p2 := person{
        name: "NavaGanesh",
        age: 30,
    }

    if p1 == p2 {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
```

## Embedding in Go
- Embedding is a way to include one struct into another struct. Embedding is also known as composition. 
- The embedded struct is called the inner struct and the struct that contains the inner struct is called the outer struct.
- 

has-a relationship:
- The outer struct has an inner struct.

in-a-relationship:
- The outer struct is in a relationship with the inner struct.

Composition through embedding:
The outer struct is composed of the inner struct. Composition is a more general term than embedding. 

Example:
```
type person struct {
    name string
    age int
}

type employee struct {
    person
    id int
}

```

## How to access embedded struct fields ?
```
outerStruct.innerStruct.fieldName

Example:
employee.person.name
employee.person.age
employee.id
```

## Printing struct fields
- Structs can be printed using the fmt package.

Syntax:
```
fmt.Printf("%+v", structName)
```



  
