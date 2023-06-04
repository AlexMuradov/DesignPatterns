package main

type GloriaJean struct {
}

func (g *GloriaJean) SellCoffee() CoffeeInterface {
	return &GloriaJeanCoffee{
		Coffee: Coffee{
			Name:  "Gloria Jean Coffee",
			Price: 5.0,
		},
	}
}

func (g *GloriaJean) SellDesert() DesertInterface {
	return &GloriaJeanDesert{
		Desert: Desert{
			Name:  "Gloria Jean Desert",
			Price: 7.0,
		},
	}
}

func (g *GloriaJean) SellFood() FoodInterface {
	return &GloriaJeanFood{
		Food: Food{
			Name:  "Gloria Jean Food",
			Price: 15.0,
		},
	}
}
