package item

type Shield struct {
}

// MaxCount always returns 1.
func (Shield) MaxCount() int {
	return 1
}

// DurabilityInfo ...
func (s Shield) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: 337,
		BrokenItem:    simpleItem(Stack{}),
	}
}

// EncodeItem ...
func (Shield) EncodeItem() (name string, meta int16) {
	return "minecraft:shield", 0
}
