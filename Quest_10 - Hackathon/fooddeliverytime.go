package piscine_go

type snacks struct {
	cooktime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]snacks{
		"burger":  {cooktime: 15},
		"chips":   {cooktime: 10},
		"nuggets": {cooktime: 12},
	}

	item, okay := menu[order]
	if !okay {
		return 404
	}

	return item.cooktime
}
