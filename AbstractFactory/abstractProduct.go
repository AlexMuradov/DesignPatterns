package main

type CoffeeInterface interface {
	getName() string
	getPrice() float64
}

type Coffee struct {
	Name  string
	Price float64
}

func (c *Coffee) getName() string {
	return c.Name
}

func (c *Coffee) getPrice() float64 {
	return c.Price
}

type DesertInterface interface {
	getName() string
	getPrice() float64
}

type Desert struct {
	Name  string
	Price float64
}

func (d *Desert) getName() string {
	return d.Name
}

func (d *Desert) getPrice() float64 {
	return d.Price
}

type FoodInterface interface {
	getName() string
	getPrice() float64
}

type Food struct {
	Name  string
	Price float64
}

func (f *Food) getName() string {
	return f.Name
}

func (f *Food) getPrice() float64 {
	return f.Price
}
