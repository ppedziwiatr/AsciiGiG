package common

type EffectType string

const (
	EffectDamageOverTime  EffectType = "damage_over_time"
	EffectHealOverTime    EffectType = "heal_over_time"
	EffectBuffAttribute   EffectType = "buff_attribute"
	EffectDebuffAttribute EffectType = "debuff_attribute"
)
