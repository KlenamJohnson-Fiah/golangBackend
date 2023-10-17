package accountnumbergenerator

import "math/rand"

func RandomString(n int) string {
	var Number = []rune("0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = Number[rand.Intn(len(Number))]
	}
	return string(b)
}
