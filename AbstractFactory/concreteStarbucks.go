package main

type Starbucks struct {
}

func (s *Starbucks) SellCoffee() CoffeeInterface {
	return &StarbucksCoffee{
		Coffee: Coffee{
			Name:  "Starbucks Coffee",
			Price: 5.0,
		},
	}
}

func (s *Starbucks) SellDesert() DesertInterface {
	return &StarbucksDesert{
		Desert: Desert{
			Name:  "Starbucks Desert",
			Price: 7.0,
		},
	}
}

func (s *Starbucks) SellFood() FoodInterface {
	return &StarbucksFood{
		Food: Food{
			Name:  "Starbucks Food",
			Price: 15.0,
		},
	}
}
