// --Summary:
//
//	Implement receiver functions to create stat modifications
//	for a video game character.
package main

import (
	"fmt"
)

//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Health int
type Energy int
type Name string
type Player struct {
	health, maxHealth Health
	energy, maxEnergy Energy
	name              Name
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

func (player *Player) stats() {
	fmt.Printf("%s HP %d/%d M %d/%d\n", player.name, player.health, player.maxHealth, player.energy, player.maxEnergy)
}

const (
	Damage = 11
	Heal   = 9
	Cost   = 7
)

func (player *Player) heal() {
	fmt.Println(player.name, "casting heal")
	if player.energy < Cost {
		return
	}
	player.energy -= Cost
	player.health += Heal
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	player.stats()
}

func (player *Player) castMagicMissile(other *Player) {
	fmt.Println(player.name, "casting magic missile upon", other.name)
	if player.energy < Cost {
		return
	}
	player.energy -= Cost
	other.health -= Damage
	if other.health < 0 {
		other.health = 0
	}
	player.stats()
	other.stats()
}

func (player *Player) rest() {
	fmt.Println(player.name, "taking a rest")
	player.energy = player.maxEnergy
	player.health = player.maxHealth
	player.stats()
}

func main() {
	goblin := Player{name: "Bilbo", health: 14, maxHealth: 14, energy: 23, maxEnergy: 23}
	ork := Player{name: "Ushkar", health: 72, maxHealth: 72, energy: 3, maxEnergy: 3}

	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)

	goblin.rest()
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&goblin)
	goblin.heal()

	goblin.rest()
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)
	goblin.castMagicMissile(&ork)

}
