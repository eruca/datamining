package distance

import (
	"math"
	"sort"
)

var (
	users = map[string]map[string]float64{
		"Angelica": {
			"Blue Traveler":    3.5,
			"Broken Bells":     2.0,
			"Norah Jones":      4.5,
			"Phoenix":          5.0,
			"Slightly Stoopid": 1.5,
			"The Strokes":      2.5,
			"Vampire Weekend":  2.0,
		},
		"Bill": {
			"Blue Traveler":    2.0,
			"Broken Bells":     3.5,
			"Deadmau5":         4.0,
			"Phoenix":          2.0,
			"Slightly Stoopid": 3.5,
			"Vampire Weekend":  3.0,
		},
		"Chan": {
			"Blue Traveler":    5.0,
			"Broken Bells":     1.0,
			"Deadmau5":         1.0,
			"Norah Jones":      3.0,
			"Phoenix":          5.0,
			"Slightly Stoopid": 1.0,
		},
		"Dan": {
			"Blue Traveler":    3.0,
			"Broken Bells":     4.0,
			"Deadmau5":         4.5,
			"Norah Jones":      4.5,
			"Phoenix":          3.0,
			"Slightly Stoopid": 4.5,
			"The Strokes":      4.0,
			"Vampire Weekend":  2.0,
		},
		"Hailey": {
			"Broken Bells":    4.0,
			"Deadmau5":        1.0,
			"Norah Jones":     4.0,
			"The Strokes":     4.0,
			"Vampire Weekend": 1.0,
		},
		"Jordyn": {
			"Broken Bells":     4.5,
			"Deadmau5":         4.0,
			"Norah Jones":      5.0,
			"Phoenix":          5.0,
			"Slightly Stoopid": 4.5,
			"The Strokes":      4.0,
			"Vampire Weekend":  4.0,
		},
		"Sam": {
			"Blue Traveler":    5.0,
			"Broken Bells":     2.0,
			"Norah Jones":      3.0,
			"Phoenix":          5.0,
			"Slightly Stoopid": 4.0,
			"The Strokes":      5.0,
		},
		"Veronica": {
			"Blue Traveler":    3.0,
			"Norah Jones":      5.0,
			"Phoenix":          4.0,
			"Slightly Stoopid": 2.5,
			"The Strokes":      3.0,
		},
	}
)

// Computes the Manhattan distance,
// Both rating1 and rating2 are dictionaryies of the form
func manhattan(rating1, rating2 map[string]float64) (distance float64) {
	for k, v := range rating1 {
		if v2, ok := rating2[k]; ok {
			distance += math.Abs(v - v2)
		}
	}
	return distance
}

type nameWithDist struct {
	name     string
	distance float64
}

type byDist []nameWithDist

func (by byDist) Len() int {
	return len(by)
}

func (by byDist) Less(i, j int) bool {
	return by[i].distance < by[j].distance
}

func (by byDist) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

// creates a sorted list of users based on their distance to username
func computeNearestNeighbor(username string, users map[string]map[string]float64) []nameWithDist {
	result := make(byDist, len(users)-1)

	var distance float64
	index := 0
	for name, _ := range users {
		if name != username {
			distance = manhattan(users[username], users[name])
			result[index] = nameWithDist{name, distance}
			index++
		}
	}
	sort.Sort(result)
	return result
}
