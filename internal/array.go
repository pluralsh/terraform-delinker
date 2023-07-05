package internal

func Includes[T comparable](arr []T, e T) bool {
	for _, el := range arr {
		if el == e {
			return true
		}
	}

	return false
}

func IncludesArray[T comparable](arr []T, partial []T) bool {
	mem := make(map[T]struct{}, 0)

	for _, el := range arr {
		mem[el] = struct{}{}
	}

	for _, el := range partial {
		if _, exists := mem[el]; exists {
			return true
		}
	}

	return false
}
