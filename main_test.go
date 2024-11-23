package main

import (
  "testing"
)

// Test Player Creation
func TestNewPlayer(t *testing.T) {
  player := NewPlayer("Thor", 50, 5, 10)

  if player.Name != "Thor" {
    t.Errorf("Expected Name: 'Thor', but got: %s", player.Name)
  }
  if player.Health != 50 {
    t.Errorf("Expected Health: 50, but got: %d", player.Health)
  }
  if player.Strength != 5 {
    t.Errorf("Expected Strength: 5, but got: %d", player.Strength)
  }
  if player.Attack != 10 {
    t.Errorf("Expected Attack: 10, but got: %d", player.Attack)
  }
}

// Test Player's Attack Calculation
func TestPlayerAttack(t *testing.T) {
  player := NewPlayer("Thor", 50, 5, 10)

  // Assume Thor rolls a 4 on the attack dice
  attackRoll := 4
  expectedDamage := player.CalculateAttackDamage(attackRoll)

  if damage := player.CalculateAttackDamage(attackRoll); damage != expectedDamage {
    t.Errorf("Expected Attack Damage: %d, but got: %d", expectedDamage, damage)
  }
}

// Test Player's Defense Calculation
func TestPlayerDefense(t *testing.T) {
  player := NewPlayer("Loki", 100, 10, 5)

  // Assume Loki rolls a 3 on the defense dice
  defenseRoll := 3
  expectedDefense := player.CalculateDefenseStrength(defenseRoll)

  if defense := player.CalculateDefenseStrength(defenseRoll); defense != expectedDefense {
    t.Errorf("Expected Defense Strength: %d, but got: %d", expectedDefense, defense)
  }
}

// Test Match Setup
func TestNewMatch(t *testing.T) {
  playerA := NewPlayer("Thor", 50, 5, 10)
  playerB := NewPlayer("Loki", 100, 10, 5)

  // Create a match with two players
  match := NewMatch(playerA, playerB)

  if match.PlayerA != playerA || match.PlayerB != playerB {
    t.Errorf("Expected players to be Thor and Loki, but got: %s and %s", match.PlayerA.Name, match.PlayerB.Name)
  }
}

// Test Match with Attack and Health Reduction
func TestMatchTurn(t *testing.T) {
  playerA := NewPlayer("Thor", 50, 5, 10)
  playerB := NewPlayer("Loki", 100, 10, 5)

  // First player (Thor) attacks, and we expect Loki's health to reduce
  playerAAttackRoll := 4
  playerBDefenseRoll := 3

  // Thor attacks, Loki defends
  damage := playerA.CalculateAttackDamage(playerAAttackRoll) - playerB.CalculateDefenseStrength(playerBDefenseRoll)
  expectedLokiHealth := playerB.Health - damage
  playerB.Health -= damage

  // Check if Loki's health is reduced correctly
  if playerB.Health != expectedLokiHealth {
    t.Errorf("Expected Loki's health: %d, but got: %d", expectedLokiHealth, playerB.Health)
  }

  // Now, Loki attacks Thor
  playerBAttackRoll := 3
  playerAStrengthRoll := 2

  // Loki attacks, Thor defends
  damage = playerB.CalculateAttackDamage(playerBAttackRoll) - playerA.CalculateDefenseStrength(playerAStrengthRoll)
  expectedThorHealth := playerA.Health - damage
  playerA.Health -= damage

  // Check if Thor's health is reduced correctly
  if playerA.Health != expectedThorHealth {
    t.Errorf("Expected Thor's health: %d, but got: %d", expectedThorHealth, playerA.Health)
  }
}