# Magical Arena Game

## Overview

Magical Arena is a turn-based combat game where two players battle until one player's health reaches zero. Players attack and defend using dice rolls, and the damage dealt is calculated based on their attributes (health, strength, and attack). The player with lower health attacks first, and the game continues until a player's health is depleted.

## Game Rules

1. **Player Attributes**:
   - Each player has three attributes:
     - `Health`: Represents the player's health. If it reaches zero, the player loses.
     - `Strength`: Represents the player's defense ability.
     - `Attack`: Represents the player's offensive ability.

2. **Dice Roll**:
   - Both players roll a 6-sided die for their attack and defense in each turn.
   - The attacker multiplies their `Attack` attribute by the die roll to calculate their damage.
   - The defender multiplies their `Strength` attribute by their die roll to calculate the defense strength.
   - The difference between the attack damage and defense strength is subtracted from the defender's health.

3. **Turn Order**:
   - The player with the lower health attacks first.

4. **Game End**:
   - The game ends when one player's health reaches zero.

## Setup

1. **Start the Game**:
   - Navigate to the magical arena directory and then run the command given below. 
   ```bash
   go run .

2. **Run Unit Test cases**
   - Navigate to the magical arena directory and then run the given command.
  ```bash
   go test
