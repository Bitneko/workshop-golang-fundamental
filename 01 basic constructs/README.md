### Basics Constructs

#### Anatomy of a Go package
Every package being with a main package.
A import statement to bring in external packages.
And a `main` function

```
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World")
}
```

#### Variables
Go variable has a default value when declared without assigning a value. The default value of a `int` is `0`, and the default value of a `string` is a empty string
```
var foo int
var bar string

fmt.Println(foo) //Print 0
fmt.Println(bar) //Print ""
```

Go can infer a variable's type implicitly
```
var foo int = 6

bar := 7
foobar := foo + bar
fmt.Println(foobar) //Print 13
```

#### Condition
```
count := 7

if count > 6 {
  fmt.Println("More than 6")
} else if count < 3 {
  fmt.Println("Less than 3")
} else {
  fmt.Println("Between 3 and 6")
}
```

#### Array
Array has a fixed size. When you pass around an array in Go, a new copy of the array is created which could eat up a lot of memory,  however it does has a faster access speed.

Note: Array are actually pretty rare in Go code.

```
var foo [5]int
foo[2] = 7

bar := [5]int{1, 2, 3, 4, 5}

fmt.Println(foo) //Print {0 0 7 0 0}
fmt.Println(bar) //Print {1 2 3 4 5}
```

#### Slices
A slice is a reference to a contiguous segment of an array. Slice is basically a pointer, with additional properties about the array.

The zero value of a slice is `nil`.

A slice is a structure of 3 fields.
- a pointer to the underlying array
- length of the slice
- capacity of the slice

- #### Literals
```
var foo = []int{1,2,3,4,5}

fmt.Println(foo) //Print {1 2 3 4 5}
```

- #### Appending to a slice
```
foo = append(foo, 6)

fmt.Println(foo) //Print {1 2 3 4 5 6}
```

#### Map
Map types are reference types, like pointers or slices, and so the value of `m` below is `nil` as it doesn't point to an initialised map.

```
var m map[string]int

// Initialise: using map or with an empty map
m = make(map[string]int) // map
m = map[string]int{} // empty map

// Set the key 'route' to value '66'
m["route"] = 66

// Remove an entry from the map
delete(m, "route")
// Retrieving and testing the existence of a key
i, ok := m["route"]
fmt.Print(i, ok) //Print 66 true

// Testing the existence of a key and ignoring the value
_, ok := m["route"]
fmt.Print(ok) //Print true
```

```

// Initialising map with data
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}

fmt.Println(commits) //Print map[adg:912 gri:1908 r:2138 rsc:3711]
```

#### Loops

```
sum := 0
for i := 0; i < 100; i++ {
  sum += i
}
```

#### Range

```
numbers := []int{1,2,3,4,5}
sum := 0
for index, number := range numbers {
    sum += number
    fmt.Printf("[%d, %d, %d]\n", index, number, sum)
```

#### Struct
Struct are typed collections of fields

```
type Box struct {
  Length int
  Width int
}

boxA := Box{5, 5}
boxB := Box{
  Length: 10,
  Width:  5,
}

boxAreaA := boxA.Length * boxA.Width
boxAreaB := boxB.Length * boxB.Width

fmt.Println(boxAreaA) //Print 25
fmt.Println(boxAreaB) //Print 50

// Embedding Type
type NestedBox struct {
  Box    Box
  Length int
  Width  int
}

nestedBox := NestedBox{
  Box{5, 5},
  6,
  6,
}

fmt.Println(nestedBox) //Print {{5,5}, 6, 6}
fmt.Println(nestedBox.Box) //Print {5, 5}

//Return main.Box - main because it's the main package
fmt.Println(reflect.TypeOf(nestedBox.Box))// Embedding Type
type NestedBox struct {
  Box    Box
  Length int
  Width  int
}

nestedBox := NestedBox{
  Box{5, 5},
  6,
  6,
}
```

```
type Student struct {
    Name string
    Age  int
}

Way to initialise a struct
// Using new keywork to initialise
var student0pa *Student
student0 = new(Student)
student0.Name = "Alice"
fmt.Println(student0) //Print &{ 0} - Age initialised to 0

// Struct Literal
student1 := Student{Name: "John", Age: 16}
fmt.Println(student1) //Print {John 16}

dy", 15}
fmt.Println(student2) //Print {Wendy 15}

student3 := StudentAlice{}
fmt.Println(student3) //Print { 0}
student2 := Student{"Wen
student4 := &Student{Name: "Matt", Age: 14}
fmt.Println(student4) //Print &{Matt 14}
```

#### Functions
Functions in Go are first class citizens. They can be assigned to variables, passed as an argument, immediately invoked or deferred for last execution

```
func add(a int, b int) int {
  return a + b
}

// Shorthand for when arguments are all of the same type
func add(a, b int) int {
  return a + b
}

// Multi return values
func addminus(a, b int) (int, int) {
  return a + b, a - b
}

addition, subtraction := addminus(5, 3)
fmt.Println(addition, subtraction) //Print 8 2

// Ignore return with underscore
addition, _ := addminus(5, 3)
fmt.Println(addition) //Print 8

// Named return - explicitly mention return variables in the function definition itself
func addMulti(a int, b int) (addition int, multiply int) {
  addition = a + b
  multiply = a * b
  
  return
}
addition, multiple := addMulti(5, 3)
fmt.Println(addition, multiple) //Print 8 15
```

- ##### Deferred function
A **defer statement** pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

```
func CopyFile(destination, source string) (written int64, err error) {
    src, err := os.Open(source)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(destination)
    if err != nil {
        return
    }

    defer dst.Close()
    return io.Copy(dst, src)
}
```


- ##### Variadic function
```
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

// Variadic functions can be called in the usual way
// with individual arguments.
sum(1, 2) //Print [1 2] 3
sum(1, 2, 3) //Print [1 2 3] 6

   // With slices, you just pass it with argument...
nums := []int{1, 2, 3, 4}
sum(nums...) //Print [1 2 3 4] 10
}

```

#### Private and Public Visibility
Exports in Go are controlled by naming convention. A capital letter means something will be exported, and a lowercase letter means it will not be exported, and this convention is applicable to structs and other data, functions, and methods.

```
type Menu struct { // The struct itself is exported
	hamburger      string // This is not exported
	ChickenNuggets string // This is exported
}
```

It is perfectly valid to have un-exported structs with exported fields. Declaring a struct outside the package will result in an runtime error, but there is a valid use case when you need to pass an instance of the struct to a exported function in another package, such as [`encoding/json`](https://golang.org/pkg/encoding/json/) package

```
type response struct { // The struct itself is not exported
    Success bool   `json:"success"` // This is exported
    Message string `json:"message"`// This is exported
    Data    string `json:"data"`// This is exported
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    resp := &response{
        Success: true,
        Message: "OK",
        Data:    "some data",
    }
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        // Handle err
    }
}
```