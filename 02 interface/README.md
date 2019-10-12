# Interface

An **interface** is a **programming** structure/syntax that allows the computer to enforce certain properties on an object (class).

In Go, the primary job of an **interface** is to provide method signatures consisting of the method name, input arguments and return types. It is up to a Type (e.g. struct type) to declare methods and implement them.

In Go, a user defined struct type is also known as the **Concrete Type**, and **interface** types provide **contracts** to concrete types

## Basic Implementation

[Playground example](https://play.golang.org/p/GiTa0V4I5PY)
```
type (
  Shape interface {
    Area() int
    Perimeter() int
  }

  Rectangle struct { // Rectangle struct is the concrete type to Shape interface
    width  int
    height int
  }
)

// Area method is implemented by type of Rectangle
func (r Rectangle ) Area() int {
  return r.width * r.height
}

// Perimeter method is implemented by type of Rectangle
func (r Rectangle ) Perimeter() int {
  return 2 *(r.width * r.height)
}

func main() {
  var rectangle Shape = Rectangle{5, 4}

  // Rectangle implement Shape interface,
  // because the methods are defined by Shape interface 
  area := rectangle.Area()
  perimeter := rectangle.Perimeter()

  fmt.Println(area)
  fmt.Println(perimeter)
}
```

## Multiple Interfaces
A type can implement multiple interfaces

[Playground example](https://play.golang.org/p/iKEL2vZjv40)
```
type (
  Shape interface {
    Area() float64
  }

  Object interface {
    Volume() float64
  }

  Cube struct {
    side float64
  }
)

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
    cube := Cube{3}
  
    // The DYNAMIC values of the following 2 var decalaration are cube
    // and the DYNAMIC type are Cube struct
    var shape Shape = cube // The static type of shape is Shape
    var object Object = cube // The static type of object is Object

    fmt.Println("Volume of shape of interface type Shape is", shape.Area())
    fmt.Println("Area of object of interface type Object is", object.Volume())
}
```

## Empty Interface
When an interface has zero methods, it is called an **empty interface**. This is represented by `interface{}`. Since empty interface has zero methods, all types implement this interface.

[Playground example](https://play.golang.org/p/N3tHmSYh_9J)
```
type HelloWorld string

type Rect struct {
  width  float64
  height float64
}

// explain function take in argument of all types because it's typed as interface{}
// which accept all type
func explain(i interface{}) {
  fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
  helloWorld := HelloWorld("Hello World!")
  rectangle := Rect{5.5, 4.5}
  explain(helloWorld)
  explain(rectangle)
}
```

## Type Assertion
We can find out the underlying dynamic value of an interface using the syntax `i.(Type)` where `i` is an interface and `Type` is a type that implements the interface `i`. **Go will check if dynamic type of** `i` **is identical to** `Type`.

[Playground example](https://play.golang.org/p/yeq27clLCg3)
```
type (
  Shape interface {
    Area() float64
  }

  Object interface {
    Volume() float64
  }

  // Skin interface does not implement Area() or Volume()
  Skin interface {
    Color() float64 // Volume method is NOT implemented by type of Cube
  }

  Cube struct {
    side float64
  }
)

// Area method is implemented by type of Cube
func (c Cube) Area() float64 {
  return 6 * (c.side * c.side)
}

// Volume method is implemented by type of Cube
func (c Cube) Volume() float64 {
  return c.side * c.side * c.side
}

func main() {
  // Declare shape as interface type Shape which also implement Object implicitly
  var shape Shape = Cube{3}

  value1, ok1 := shape.(Object)
  fmt.Printf("Dynamic value of Shape 's' with value %v implements interface Object? %v\n", value1, ok1)

  value2, ok2 := shape.(Skin) // Color() is not implemented for Cube
  fmt.Printf("Dynamic value of Shape 's' with value %v implements interface Skin? %v\n", value2, ok2)
}
```

## Type Switching
**Type switch**. The syntax for type switch is similar to type assertion and it is `i.(type)`where `i` is interface and `type` is a fixed keyword. Using this we can get the concrete type of the interface instead of value. **But this syntax will only work in** **`switch`** **statement**.

[Playground example](https://play.golang.org/p/nvOral9xrwm)
```
func explain(i interface{}) {
  switch i.(type) {
    case string:
      fmt.Println("i stored string ", strings.ToUpper(i.(string)))
    case int:
      fmt.Println("i stored int", i)
    default:
      fmt.Println("i stored something else", i)
    }
}

func main() {
  explain("Hello World")
  explain(52)
  explain(true)
}
```

## Embedding Composition
In Go, an interface cannot implement other interfaces or extend them, but we can create new interface by merging two or more interfaces.

[Playground example](https://play.golang.org/p/6F3ZSsFCsyV)
```
type (
  Shape interface {
    Area() float64
  }

  Object interface {
    Volume() float64
  }

  Material interface {
    Shape
    Object
  }

  Cube struct {
    side float64
  }
)

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
  c := Cube{3}

  var m Material = c
  var s Shape = c
  var o Object = c
	
  // Dynamic type and values of m, s and o are exactly the same,
  // because the methods defined in the interfaces are all implemented by Cube
  fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
  fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
  fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}
```

## Pointer vs Value Receiver
In struct,  a method with pointer receiver will work on both pointer or value, But in case of interfaces, if a method has a pointer receiver, then the **interface will have a pointer of dynamic type rather than the value of dynamic type**.

[Playground example (with assigned pointer)](https://play.golang.org/p/I3l5R1tsCwG)<br />
[Playground example (with assigned value)](https://play.golang.org/p/IacrF3PcGCL)
```
type (
  Shape interface {
    Area() float64
    Perimeter() float64
  }

  Rect struct {
    width  float64
    height float64
  }
)

// Receiver is a pointer
func (r *Rect) Area() float64 {
  return r.width * r.height
}

// Receiver is a value
func (r Rect) Perimeter() float64 {
  return 2 * (r.width + r.height)
}

func main() {
  r := Rect{5.0, 4.0}
  // r.Area() and r.Perimeter will work because, a method with
  // pointer receiver will work on both pointer or value for struct

  var s Shape = &r // assigned pointer
  area := s.Area()
  perimeter := s.Perimeter()

  fmt.Println("area of rectangle is", area)
  fmt.Println("perimeter of rectangle is", perimeter)
}
```
