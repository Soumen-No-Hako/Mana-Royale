package randGen

import (
	"math/rand/v2"
)
func Generate(limit int) int {
      return rand.IntN(limit)
}
