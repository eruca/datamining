package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeNearestNeighbor(t *testing.T) {
	result := computeNearestNeighbor("Hailey", users)

	assert.Equal(t, result, []nameWithDist{{"Veronica", 2.0}, {"Chan", 4.0},
		{"Sam", 4.0}, {"Dan", 4.5}, {"Angelica", 5.0}, {"Bill", 5.5}, {"Jordyn", 7.5}})
}

func TestRecommend(t *testing.T) {
	result := Recommend("Hailey", users)
	assert.Equal(t, result, []nameWithDist{{Bangs[3], 4.0}, {Bangs[0], 3.0}, {Bangs[4], 2.5}})

	result = Recommend("Chan", users)
	assert.Equal(t, result, []nameWithDist{{Bangs[5], 4.0}, {Bangs[6], 1.0}})

	result = Recommend("Sam", users)
	assert.Equal(t, result, []nameWithDist{{Bangs[7], 1.0}})

	result = Recommend("Angelica", users)
	assert.Nil(t, result)
}
