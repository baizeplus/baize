package arrayUtils

func IsInArray[T comparable](val T, array []T) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}
