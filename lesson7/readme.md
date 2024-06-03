# HOTOGO-Lesson7

# Consumer Interface

- This is how we implement polymorphism in Golang

## 1.What is polymorphism?

Polymorphism is an must-knowned charact-eristic in Object Oriented Programming. It is a way that we can specifically implement the p-articular properties from the root/abstract cl-ass

![Untitled](HOTOGO-Lesson7%208eddcb8e087c4631be0d89c49b4933b3/Untitled.png)

## 2.Consumer interface

### 2.1.Template

```go
Here is the template:

type AbstractClass interface{
	FunctionName()
}

type Particular1 struct{
	attributes_1
	...
	attribute_n
	}

func (f*Particular1)FunctionName(){
//code
}
```

### 2.2.Example

It is a like-related concept in Go lang development that allows us to implement polymorphism.

Here is my code example of polymorphism.

![Untitled](HOTOGO-Lesson7%208eddcb8e087c4631be0d89c49b4933b3/Untitled%201.png)

```go
type Vehicle interface {
	Moving()
}

```

```go
type Car struct {
	wheels int
	seats  int
	brand  string
	price  string
}
func (f Car) Moving() {
	fmt.Printf("The car has %v wheels and %v seats\n", f.wheels, f.seats)
	fmt.Printf("It is assembled by %v brand and sold at %v\n\n", f.brand, f.price)
}
```

```go
type Motobike struct {
	wheels   int
	seats    int
	brand    string
	velocity string
}

func (f Motobike) Moving() {
	fmt.Printf("The motobike has %v wheels and %v seats\n", f.wheels, f.seats)
	fmt.Printf("It is assembled by %v brand with the high speed of %v\n\n", f.brand, f.velocity)
}
```

```go
type Airplane struct {
	brand string
	seats int
}

func (f Airplane) Moving() {
	fmt.Printf("We are flying on the skyyyyyyyyy......\n")
	fmt.Printf("We comes from %v having %v seats\n\n", f.brand, f.seats)

}
```

```go
func main() {
	vehicles := make([]Vehicle, 3)
	vehicles[0] = Car{
		wheels: 4,
		seats:  4,
		brand:  "Roll Roycesa",
		price:  "1.5 M dollars",
	}
	vehicles[1] = Motobike{
		wheels:   2,
		seats:    2,
		brand:    "Kawasaki",
		velocity: "300km/h",
	}
	vehicles[2] = Airplane{
		brand: "Boing",
		seats: 255,
	}
	for _, vehicle := range vehicles {
		vehicle.Moving()
	}
}
```

In this part of code, you should focus on the vehicle slice `vehicles := make([]Vehicle, 3)` and then we assign the particular implement class to vehicle 1st,2nd,3rd.

After that the result would go as below:

![Result](media\result.png)

## References

[This Will Make Everyone Understand Golang Interfaces](https://www.youtube.com/watch?v=rH0bpx7I2Dk)
