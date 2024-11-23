package main

import (
  "fmt"
  "math/rand"
)

// Player represents a player in the Magical Arena game.
type Player struct {
  Name     string
  Health   int
  Strength int
  Attack   int
}

// NewPlayer initializes a new Player.
func NewPlayer(name string, health, strength, attack int) *Player {
  return &Player{
    Name:     name,
    Health:   health,
    Strength: strength,
    Attack:   attack,
  }
}

// RollDice simulates rolling a 6-sided dice.
func (p *Player) RollDice() int {

  return rand.Intn(6) + 1
}

// TakeDamage reduces the player's health by the given damage amount.
func (p *Player) TakeDamage(damage int) {
  p.Health -= damage
  if p.Health < 0 {
    p.Health = 0
  }
}

// IsAlive checks if the player is still alive.
func (p *Player) IsAlive() bool {
  return p.Health > 0
}

// ToString provides a string representation of the player's stats.
func (p *Player) ToString() string {
  return fmt.Sprintf("%s [Health: %d, Strength: %d, Attack: %d]", p.Name, p.Health, p.Strength, p.Attack)
}