package distance

import (
	"math"
	"sort"
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
	if by[i].distance == by[j].distance {
		return by[i].name < by[j].name
	}
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

func minkowski(rating1, rating2 map[string]float64, r float64) (distance float64) {
	commonRatings := false
	for k, v := range rating1 {
		if v2, ok := rating2[k]; ok {
			distance += math.Pow(math.Abs(v-v2), r)
			commonRatings = true
		}
	}

	if commonRatings {
		return math.Pow(distance, 1/r)
	}

	return 0
}

// first find nearest neighbor
func Recommend(username string, users map[string]map[string]float64) (recommedation []nameWithDist) {
	nearest := computeNearestNeighbor(username, users)[0].name
	neighborRatings := users[nearest]
	userRatings := users[username]

	for name, dist := range neighborRatings {
		if _, ok := userRatings[name]; !ok {
			recommedation = append(recommedation, nameWithDist{name: name, distance: dist})
		}
	}
	sort.Stable(sort.Reverse(byDist(recommedation)))
	return
}
