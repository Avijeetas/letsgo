package main

import "fmt"

type vehicle interface {
	color() string
	speed() string
	company() string
}

type truck struct {
}
type car struct {
}
type cycle struct {
}

func (t truck) color() string {
	return "The color of truck is blue."
}

func (t truck) speed() string {
	return "The speed of truck is 10km/h."
}

func (t truck) company() string {
	return "The company of truck is Tata."
}

func (c car) color() string {
	return "The color of car is blue."
}

func (c car) speed() string {
	return "The speed of car is 10km/h."
}

func (c car) company() string {
	return "The company of car is Nissan."
}
func (cy cycle) color() string {
	return "The color of cycle is blue."
}

func (cy cycle) speed() string {
	return "The speed of cycle is 1km/h."
}

func (cy cycle) company() string {
	return "The company of cycle is Duranto."
}

func printvalue(v vehicle) {

	fmt.Println(v.color())
	fmt.Println(v.speed())
	fmt.Println(v.company())

}
func main() {

	c := car{}
	t := truck{}
	cy := cycle{}
	printvalue(c)
	printvalue(t)
	printvalue(cy)

}
