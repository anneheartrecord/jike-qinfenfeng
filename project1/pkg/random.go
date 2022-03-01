package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return rndCode
}
