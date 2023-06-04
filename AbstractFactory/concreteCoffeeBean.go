package main

type CoffeeBean struct {
}

func (c *CoffeeBean) SellCoffee() CoffeeInterface {
	return &CoffeeBeanCoffee{
		Coffee: Coffee{
			Name:  "Coffee Bean Coffee",
			Price: 5.0,
		},
	}
}

func (c *CoffeeBean) SellDesert() DesertInterface {
	return &CoffeeBeanDesert{
		Desert: Desert{
			Name:  "Coffee Bean Desert",
			Price: 7.0,
		},
	}
}

func (c *CoffeeBean) SellFood() FoodInterface {
	return &CoffeeBeanFood{
		Food: Food{
			Name:  "Coffee Bean Food",
			Price: 15.0,
		},
	}
}
