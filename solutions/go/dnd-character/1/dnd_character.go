package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	modifiedScore := ( score - 10 )
    if modifiedScore % 2 != 0 {
        modifiedScore--
    }

    return modifiedScore/2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(16) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
        Strength : Ability(),
        Dexterity : Ability(),
        Constitution : Ability(),
        Intelligence : Ability(),
        Wisdom : Ability(),
        Charisma : Ability(),
    }

    c.Hitpoints = 10 + Modifier(c.Constitution)

    return c
}
