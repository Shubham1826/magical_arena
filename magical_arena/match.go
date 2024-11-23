package main

import "fmt"

// Match manages a battle between two players.
type Match struct {
  PlayerA *Player
  PlayerB *Player
}

// NewMatch initializes a new Match between two players.
func NewMatch(playerA, playerB *Player) *Match {
  return &Match{
    PlayerA: playerA,
    PlayerB: playerB,
  }
}

// PlayTurn simulates a turn where the attacker attacks the defender.
func (m *Match) PlayTurn(attacker, defender *Player) {
  attackRoll := attacker.RollDice()
  defenseRoll := defender.RollDice()

  attackDamage := attackRoll * attacker.Attack
  defenseStrength := defenseRoll * defender.Strength
  netDamage := attackDamage - defenseStrength

  fmt.Println("--------------------------------------------------")
  fmt.Printf("Turn: %s (Health %d) attacks %s (Health %d)\n", attacker.Name, attacker.Health, defender.Name, defender.Health)
  fmt.Printf("🎲 %s rolls an attack dice: %d\n", attacker.Name, attackRoll)
  fmt.Printf("🎲 %s rolls a defense dice: %d\n", defender.Name, defenseRoll)
  fmt.Printf("💥 Attack Damage: %d (Attack Roll %d * Attack %d)\n", attackDamage, attackRoll, attacker.Attack)
  fmt.Printf("🛡️  Defense Strength: %d (Defense Roll %d * Strength %d)\n", defenseStrength, defenseRoll, defender.Strength)

  if netDamage > 0 {
    defender.TakeDamage(netDamage)
    fmt.Printf("🔥 Net Damage Dealt: %d\n", netDamage)
    fmt.Printf("❤️ %s's Health Left: %d\n", defender.Name, defender.Health)
  } else {
    fmt.Printf("🛡️  %s defends successfully! No damage taken.\n", defender.Name)
  }
  fmt.Println("--------------------------------------------------")
}


// Start begins the match and alternates turns until one player dies.
func (m *Match) Start() {
  fmt.Println("🏟️  Welcome to the Magical Arena!")
  fmt.Printf("Players: %s vs %s\n", m.PlayerA.ToString(), m.PlayerB.ToString())
  fmt.Println("--------------------------------------------------")

  // Determine initial attacker
  var attacker, defender *Player
  if m.PlayerA.Health <= m.PlayerB.Health {
    attacker = m.PlayerA
    defender = m.PlayerB
  } else {
    attacker = m.PlayerB
    defender = m.PlayerA
  }

  for m.PlayerA.IsAlive() && m.PlayerB.IsAlive() {
    m.PlayTurn(attacker, defender)

    // Swap attacker and defender
    attacker, defender = defender, attacker
  }

  fmt.Println("🏆 The match is over!")
  if m.PlayerA.IsAlive() {
    fmt.Printf("🎉 Winner: %s\n", m.PlayerA.Name)
  } else {
    fmt.Printf("🎉 Winner: %s\n", m.PlayerB.Name)
  }
  fmt.Println("--------------------------------------------------")
}