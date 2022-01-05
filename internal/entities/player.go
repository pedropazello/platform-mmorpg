package entities

type Player struct {
	Id               string
	Name             string
	Experience       int
	Level            int
	Health           int
	MaxHealth        int
	EquipedWeapon    Item
	EquipedArmor     Item
	Inventory        []Item
	CurrentPositionX int
	CurrentPositionY int
}
