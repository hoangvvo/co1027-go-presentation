package eat

func consumeFood(nutritionalValues int) {
	// consuming food
}

func processFood(food string, quantity int) int {
	return 1000
}

func Eat(food string, quantity int) {
	val := processFood(food, quantity)
	consumeFood(val)
	useBathroom()
}
