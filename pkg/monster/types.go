package monster

type Attributes struct {
	Agility      int // Affects turn order, dodge chance
	Strength     int // Affects HP, melee damage
	Charisma     int // Affects persuasion resistance
	Intelligence int // Affects magic power, mana
}

type MonsterType string

const (
	MonsterTypeDemon     MonsterType = "demon"
	MonsterTypeUndead    MonsterType = "undead"
	MonsterTypeAnimal    MonsterType = "animal"
	MonsterTypeElement   MonsterType = "elemental"
	MonsterTypeAberation MonsterType = "aberration"
)

type MonsterBehavior int

const (
	BehaviorAggressive MonsterBehavior = iota
	BehaviorPatrolling
	BehaviorCowardly
	BehaviorSmart
)

type Ability struct {
	Name        string
	Description string
	ManaCost    int
	CoolDown    int
	CurrentCD   int
	Type        AbilityType
	Power       int
	Effect      *Effect
}

type AbilityType string

const (
	AbilityTypeAttack  AbilityType = "attack"
	AbilityTypeHeal    AbilityType = "heal"
	AbilityTypeBuff    AbilityType = "buff"
	AbilityTypeDebuff  AbilityType = "debuff"
	AbilityTypeSpecial AbilityType = "special"
)

type Effect struct {
	Name      string
	Duration  int
	Magnitude int
	Type      EffectType
}

type EffectType string

const (
	EffectDamageOverTime  EffectType = "damage_over_time"
	EffectHealOverTime    EffectType = "heal_over_time"
	EffectBuffAttribute   EffectType = "buff_attribute"
	EffectDebuffAttribute EffectType = "debuff_attribute"
)

type MonsterTemplate struct {
	Name           string
	Type           MonsterType
	Symbol         rune
	BaseAttributes Attributes
	Abilities      []Ability
	ExpValue       int
	Level          int
}

func (m *Monster) GetPosition() (int, int) {
	return m.Position.X, m.Position.Y
}

func (m *Monster) SetPosition(x, y int) {
	m.Position.X = x
	m.Position.Y = y
}

func (m *Monster) GetSymbol() rune {
	return m.Symbol
}

func (m *Monster) GetName() string {
	return m.Name
}

func GetBasicAttack(level int) Ability {
	return Ability{
		Name:        "Attack",
		Description: "Basic attack",
		ManaCost:    0,
		CoolDown:    0,
		Type:        AbilityTypeAttack,
		Power:       3 + level/2,
	}
}
