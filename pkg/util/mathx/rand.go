package mathx

import "math/rand"

// Shuffle randomly shuffles a list. Convenience function.
func Shuffle[T any](list []T) {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}
