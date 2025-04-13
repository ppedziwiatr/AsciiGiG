package game

import (
	"fmt"
	"math/rand"

	"github.com/ppedziwiatr/ascii-gig/pkg/character"
	"github.com/ppedziwiatr/ascii-gig/pkg/item"
	"github.com/ppedziwiatr/ascii-gig/pkg/monster"
)

func (g *Game) StartCombat(target *monster.Monster) {
	g.Mode = ModeCombat
	g.CurrentTarget = target
	g.AddMessage(fmt.Sprintf("You engage in combat with %s!", target.Name))
}

func (g *Game) EndCombat() {
	g.Mode = ModeExploring
	g.CurrentTarget = nil
}

func (g *Game) UseAbility(abilityIndex int) bool {
	player := g.Player
	target := g.CurrentTarget

	if abilityIndex < 0 || abilityIndex >= len(player.Abilities) {
		g.AddMessage("Invalid ability!")
		return false
	}

	ability := player.Abilities[abilityIndex]

	if ability.CurrentCD > 0 {
		g.AddMessage(fmt.Sprintf("%s is still on cooldown for %d turns!", ability.Name, ability.CurrentCD))
		return false
	}

	if player.Mana < ability.ManaCost {
		g.AddMessage("Not enough mana!")
		return false
	}

	player.Mana -= ability.ManaCost

	player.Abilities[abilityIndex].CurrentCD = ability.CoolDown

	power := ability.Power

	switch ability.Type {
	case character.AbilityTypeAttack:
		if isAbilityMagical(ability) {
			power += player.Attributes.Intelligence / 3
		} else if isAbilityRanged(ability) {
			power += player.Attributes.Agility / 3
		} else {
			power += player.Attributes.Strength / 3
		}
	case character.AbilityTypeHeal:
		power += player.Attributes.Intelligence / 4
	}

	switch ability.Type {
	case character.AbilityTypeAttack:

		hitChance := 0.8 + float64(player.Attributes.Agility-target.Attributes.Agility)*0.02
		hitChance = max(0.5, min(0.95, hitChance)) // Clamp between 50% and 95%

		if rand.Float64() <= hitChance {

			target.Health -= power
			g.AddMessage(fmt.Sprintf("You use %s and deal %d damage to %s!",
				ability.Name, power, target.Name))

			if ability.Effect != nil {

			}

			if target.Health <= 0 {
				g.handleMonsterDefeat(target)
				return true
			}
		} else {

			g.AddMessage(fmt.Sprintf("You use %s but miss %s!", ability.Name, target.Name))
		}
	case character.AbilityTypeHeal:
		player.Health += power
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		g.AddMessage(fmt.Sprintf("You use %s and heal for %d health!", ability.Name, power))
	case character.AbilityTypeBuff:

		g.AddMessage(fmt.Sprintf("You use %s and feel stronger!", ability.Name))
	case character.AbilityTypeDebuff:

		g.AddMessage(fmt.Sprintf("You use %s and weaken %s!", ability.Name, target.Name))
	}

	g.monsterAttack()

	return true
}

func (g *Game) AttemptToFlee() bool {

	fleeChance := 0.4 + float64(g.Player.Attributes.Agility-g.CurrentTarget.Attributes.Agility)*0.05
	fleeChance = max(0.2, min(0.8, fleeChance)) // Clamp between 20% and 80%

	if rand.Float64() < fleeChance {
		g.AddMessage("You escaped from combat!")
		g.EndCombat()
		return true
	} else {
		g.AddMessage("You failed to escape!")

		g.monsterAttack()
		return false
	}
}

func (g *Game) AttemptToPersuade() bool {

	if g.Mode != ModeCombat {
		return false
	}

	player := g.Player
	monster := g.CurrentTarget

	persuasionChance := float64(player.Attributes.Charisma-monster.Attributes.Charisma) * 0.03
	persuasionChance = max(0.0, min(0.6, persuasionChance)) // Clamp between 0% and 60%

	if rand.Float64() < persuasionChance {

		extraGold := rand.Intn(monster.Gold/2) + 1
		player.Gold += extraGold
		g.AddMessage(fmt.Sprintf("You persuade %s to give you %d gold!", monster.Name, extraGold))

		if rand.Float64() < 0.2 && monster.DropRate > 0 {
			newItem := item.GenerateItem(g.CurrentLevel + 1)
			if player.PickUpItem(newItem) {
				g.AddMessage(fmt.Sprintf("The %s also gives you %s!", monster.Name, newItem.Name))
			} else {
				g.AddMessage("Your inventory is full, so you can't take the offered item.")

				newItem.SetPosition(monster.Position.X, monster.Position.Y)
				g.AddItem(newItem)
			}
		}

		g.RemoveMonster(monster)
		g.EndCombat()
		return true
	} else {
		g.AddMessage(fmt.Sprintf("You try to persuade %s, but they're not interested!", monster.Name))

		g.monsterAttack()
		return false
	}
}

func (g *Game) monsterAttack() {
	player := g.Player
	currentTarget := g.CurrentTarget

	ability := currentTarget.ChooseAbility()

	power := ability.Power

	switch ability.Type {
	case monster.AbilityTypeAttack:
		if isMonsterAbilityMagical(ability) {
			power += currentTarget.Attributes.Intelligence / 3
		} else if isMonsterAbilityRanged(ability) {
			power += currentTarget.Attributes.Agility / 3
		} else {
			power += currentTarget.Attributes.Strength / 3
		}
	}

	hitChance := 0.8 + float64(currentTarget.Attributes.Agility-player.Attributes.Agility)*0.02
	hitChance = max(0.5, min(0.95, hitChance)) // Clamp between 50% and 95%

	if rand.Float64() <= hitChance {

		player.Health -= power
		g.AddMessage(fmt.Sprintf("%s uses %s and deals %d damage to you!",
			currentTarget.Name, ability.Name, power))

		if ability.Effect != nil {

		}

		if player.Health <= 0 {
			g.GameOver = true
			g.AddMessage("You have been defeated!")
		}
	} else {

		g.AddMessage(fmt.Sprintf("%s uses %s but misses you!", currentTarget.Name, ability.Name))
	}
}

func (g *Game) handleMonsterDefeat(monster *monster.Monster) {
	player := g.Player

	g.AddMessage(fmt.Sprintf("You defeated %s! Gained %d experience and %d gold.",
		monster.Name, monster.ExpValue, monster.Gold))

	leveledUp := player.GainExperience(monster.ExpValue)
	player.Gold += monster.Gold

	if leveledUp {
		g.AddMessage(fmt.Sprintf("Level up! You are now level %d. Health and mana fully restored.", player.Level))
	}

	if rand.Float64() <= monster.DropRate {

		newItem := item.GenerateItem(g.CurrentLevel + 1)
		newItem.SetPosition(monster.Position.X, monster.Position.Y)
		g.AddItem(newItem)
		g.AddMessage(fmt.Sprintf("The %s dropped %s!", monster.Name, newItem.Name))
	}

	g.RemoveMonster(monster)

	g.EndCombat()
}

func isAbilityMagical(ability character.Ability) bool {
	return ability.Name == "Magic Missile" ||
		ability.Name == "Fireball" ||
		ability.Name == "Frost Nova" ||
		ability.Name == "Arcane Explosion" ||
		ability.Name == "Ice Bolt" ||
		ability.Name == "Lightning Strike"
}

func isAbilityRanged(ability character.Ability) bool {
	return ability.Name == "Shoot" ||
		ability.Name == "Quick Shot" ||
		ability.Name == "Rain of Arrows" ||
		ability.Name == "Magic Missile" ||
		ability.Name == "Fireball"
}

func isMonsterAbilityMagical(ability monster.Ability) bool {
	return ability.Name == "Fireball" ||
		ability.Name == "Hellfire" ||
		ability.Name == "Mind Blast" ||
		ability.Name == "Necrotic Blast" ||
		ability.Name == "Freeze" ||
		ability.Name == "Lightning Strike" ||
		ability.Name == "Eye Ray" ||
		ability.Name == "Cosmic Horror" ||
		ability.Name == "Elemental Surge"
}

func isMonsterAbilityRanged(ability monster.Ability) bool {
	return ability.Name == "Toxic Spit" ||
		ability.Name == "Fireball" ||
		ability.Name == "Lightning Strike" ||
		ability.Name == "Eye Ray"
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
