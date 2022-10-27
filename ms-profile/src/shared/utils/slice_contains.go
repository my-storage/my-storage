package utils

func SliceContains[SliceType comparable](target SliceType, source []SliceType) (contains bool, index int) {
	for index, value := range source {
		if value == target {
			return true, index
		}
	}

	return false, -1
}
