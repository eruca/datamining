package distance

var Bangs = []string{0: "Blue Traveler", 1: "Broken Bells", 2: "Norah Jones",
	3: "Phoenix", 4: "Slightly Stoopid", 5: "The Strokes",
	6: "Vampire Weekend", 7: "Deadmau5"}

var (
	users = map[string]map[string]float64{
		"Angelica": {
			Bangs[0]: 3.5,
			Bangs[1]: 2.0,
			Bangs[2]: 4.5,
			Bangs[3]: 5.0,
			Bangs[4]: 1.5,
			Bangs[5]: 2.5,
			Bangs[6]: 2.0,
		},
		"Bill": {
			Bangs[0]: 2.0,
			Bangs[1]: 3.5,
			Bangs[7]: 4.0,
			Bangs[3]: 2.0,
			Bangs[4]: 3.5,
			Bangs[6]: 3.0,
		},
		"Chan": {
			Bangs[0]: 5.0,
			Bangs[1]: 1.0,
			Bangs[7]: 1.0,
			Bangs[2]: 3.0,
			Bangs[3]: 5.0,
			Bangs[4]: 1.0,
		},
		"Dan": {
			Bangs[0]: 3.0,
			Bangs[1]: 4.0,
			Bangs[7]: 4.5,
			Bangs[3]: 3.0,
			Bangs[4]: 4.5,
			Bangs[5]: 4.0,
			Bangs[6]: 2.0,
		},
		"Hailey": {
			Bangs[1]: 4.0,
			Bangs[7]: 1.0,
			Bangs[2]: 4.0,
			Bangs[5]: 4.0,
			Bangs[6]: 1.0,
		},
		"Jordyn": {
			Bangs[1]: 4.5,
			Bangs[7]: 4.0,
			Bangs[2]: 5.0,
			Bangs[3]: 5.0,
			Bangs[4]: 4.5,
			Bangs[5]: 4.0,
			Bangs[6]: 4.0,
		},
		"Sam": {
			Bangs[0]: 5.0,
			Bangs[1]: 2.0,
			Bangs[2]: 3.0,
			Bangs[3]: 5.0,
			Bangs[4]: 4.0,
			Bangs[5]: 5.0,
		},
		"Veronica": {
			Bangs[0]: 3.0,
			Bangs[2]: 5.0,
			Bangs[3]: 4.0,
			Bangs[4]: 2.5,
			Bangs[5]: 3.0,
		},
	}
)
