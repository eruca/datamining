package distance

import (
	"fmt"
	"testing"
)

func TestComputeNearestNeighbor(t *testing.T) {
	result := computeNearestNeighbor("Hailey", users)

	fmt.Println(result)
}
