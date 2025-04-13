package common

type Entity interface {
	GetPosition() (int, int)
	SetPosition(x, y int)
	GetSymbol() rune
	GetName() string
}
