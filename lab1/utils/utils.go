package utils

func MakeAllSets(variablesQuantity uint8) [][]uint8 {
	sets := make([][]uint8, 1<<variablesQuantity)
	for i := range sets {
		set := fromNumberToBinaryArray(i, variablesQuantity)
		sets[i] = set
	}
	return sets
}

func fromNumberToBinaryArray(number int, arrayLength uint8) []uint8 {
	set := make([]uint8, arrayLength)
	for i := range set {
		set[len(set)-1-i] = uint8(number % 2)
		number /= 2
	}
	return set
}

func HammingDistance(sample []uint8, target []uint8) uint8 {
	var distance uint8
	for i := range sample {
		distance += sample[i] ^ target[i]
	}
	return distance
}
