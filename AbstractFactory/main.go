package main

import (
	"errors"
	"fmt"
)

type CoffeShopInterface interface {
	SellCoffee() CoffeeInterface
	SellDesert() DesertInterface
	SellFood() FoodInterface
}

func InitCoffeeShop(n string) (CoffeShopInterface, error) {

	if n == "Starbucks" {
		return &Starbucks{}, nil
	} else if n == "CoffeeBean" {
		return &CoffeeBean{}, nil
	} else if n == "GloriaJean" {
		return &GloriaJean{}, nil
	} else {
		return nil, errors.New("Invalid CoffeeShop")
	}
}

func main() {
	fmt.Println("Hello, playground")

	coffeeShop, err := InitCoffeeShop("Starbucks")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := coffeeShop.SellCoffee()
	fmt.Println(res.getName())
}
