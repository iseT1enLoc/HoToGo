package main

import "fmt"

type Vehicle interface {
	Moving()
}

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

type Airplane struct {
	brand string
	seats int
}

func (f Airplane) Moving() {
	fmt.Printf("We are flying on the skyyyyyyyyy......\n")
	fmt.Printf("We comes from %v having %v seats\n\n", f.brand, f.seats)

}
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
