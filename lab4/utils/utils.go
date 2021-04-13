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

func HammingDistance(sample []uint8, target []uint8) uint64 {
	var distance uint64
	for i := range sample {
		distance += uint64(sample[i] ^ target[i])
	}
	return distance
}

func GetCSet(target []uint8) []uint64 {
	zeroCandidate := make([]uint64, 0, len(target))
	oneCandidate := make([]uint64, 0, len(target))

	for i, value := range target {
		if value == 1 {
			oneCandidate = append(oneCandidate, uint64(i))
		} else {
			zeroCandidate = append(zeroCandidate, uint64(i))
		}
	}

	if len(zeroCandidate) < len(oneCandidate) {
		return zeroCandidate
	} else {
		return oneCandidate
	}
}
