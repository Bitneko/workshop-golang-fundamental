### Interface
An **interface** is a **programming** structure/syntax that allows the computer to enforce certain properties on an object (class).

In Go, the primary job of an **interface** is to provide method signatures consisting of the method name, input arguments and return types. It is up to a Type (e.g. struct type) to declare methods and implement them.

In Go, a user defined struct type is also known as the **Concrete Type**, and **interface** types provide **contracts** to concrete types

#### Basic Implementation
```
type Shape interface {
    Area() int
    Perimeter() int
}

type Rectangle struct {
  width  int
  height int
}

// Area method is implemented by type of Rect
func (r Rectangle ) Area() int {
  return r.width * r.height
}

// Perimeter method is implemented by type of Rect
func (r Rectangle ) Perimeter() int {
  return 2 *(r.width * r.height)
}

func main() {
  rectangle := Rectangle{5, 4}
  
  // Rectangle implement Shape interface,
  // because the methods are defined by Shape interface 
  area := rectangle.Area()
  perimeter := rectangle.Perimeter()
}

```

#### Multiple Interfaces
A type can implement multiple interfaces

```
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	cube := Cube{3}
  
    // The dynamic values of the following 2 var are cube
	var shape Shape = cube // The static type of shape is Shape
	var object Object = cube // The static type of object is Object

	fmt.Println("volume of shape of interface type Shape is", shape.Area())
	fmt.Println("area of object of interface type Object is", object.Volume())
}
```

#### Empty Interface
When an interface has zero methods, it is called an **empty interface**. This is represented by `interface{}`. Since empty interface has zero methods, all types implement this interface.

```
type HelloWorld string

type Rect struct {
	width  float64
	height float64
}

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

#### Type Assertion
We can find out the underlying dynamic value of an interface using the syntax `i.(Type)` where `i` is an interface and `Type` is a type that implements the interface `i`. **Go will check if dynamic type of** `i` **is identical to** `Type`.

```
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

// Skin interface does not implement Area() or Volume()
type Skin interface {
	Color() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	var shape Shape = Cube{3}
	value1, ok1 := shape.(Object)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Object? %v\n", value1, ok1)
	value2, ok2 := shape.(Skin)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Skin? %v\n", value2, ok2)
}
```

#### Type Switching
**Type switch**. The syntax for type switch is similar to type assertion and it is `i.(type)`where `i` is interface and `type` is a fixed keyword. Using this we can get the concrete type of the interface instead of value. **But this syntax will only work in** **`switch`** **statement**.
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

#### Embedding Composition
In Go, an interface cannot implement other interfaces or extend them, but we can create new interface by merging two or more interfaces.
```
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
}

type Cube struct {
	side float64
}

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
	fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
	fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
	fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}
```

#### Pointer vs Value Receiver
In struct,  a method with pointer receiver will work on both pointer or value, But in case of interfaces, if a method has a pointer receiver, then the **interface will have a pointer of dynamic type rather than the value of dynamic type**.
```
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

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