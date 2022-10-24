//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestRest(t *testing.T) {
	cases := []struct {
		health, maxHealth, expectedHealth Health
		energy, maxEnergy, expectedEnergy Energy
	}{
		{1, 2, 2, 0, 10, 10},
		{10, 10, 10, 4, 4, 4},
		{0, 4, 4, 0, 6, 6},
	}
	for _, tc := range cases {
		p := Player{
			name:      "Bob",
			health:    tc.health,
			maxHealth: tc.maxHealth,
			energy:    tc.energy,
			maxEnergy: tc.maxEnergy,
		}
		p.rest()
		if p.health != tc.expectedHealth {
			t.Errorf("should be max health after rest, got %v, wanted %v", p.health, tc.expectedHealth)
		}
		if p.energy != tc.expectedEnergy {
			t.Errorf("should be max energy after rest, got %v, wanted %v", p.energy, tc.expectedEnergy)
		}
	}
}
