package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	letters := []byte("adalksdaldlakoakkf123kaksafdmsldmlfs3546546445dhdfh")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for k, _ := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}

	return string(result)

}
