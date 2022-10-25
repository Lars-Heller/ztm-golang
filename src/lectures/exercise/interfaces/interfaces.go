//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lift int

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

func (l Lift) String() string {
	lifts := []string{"small lift", "standard lift", "large lift"}
	return lifts[l]
}

type Directer interface {
	directToLift() Lift
}

type MotorCycle string
type Car string
type Truck string

func (m MotorCycle) String() string {
	return fmt.Sprintf("Motorcycle %s", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car %s", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck %s", string(t))
}
func (m MotorCycle) directToLift() Lift {
	return SmallLift
}
func (c Car) directToLift() Lift {
	return StandardLift
}
func (t Truck) directToLift() Lift {
	return LargeLift
}

func directVehicles(vehicles []Directer) {
	for _, vehicle := range vehicles {
		fmt.Printf("%v gets to the %v\n", vehicle, vehicle.directToLift())
	}
}

func main() {
	vehicles := []Directer{MotorCycle("ZX-9R"), Truck("MAN"), Car("Prius")}
	directVehicles(vehicles)
}
