package currency

const ratio = 10000

func FromCoins(coins int) float64 {
	return float64(coins) / ratio
}

func ToCoins(value float64) int {
	return int(value * ratio)
}
