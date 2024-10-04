package listing

type Tuple struct {
	Name  string
	Index int
}

func of(name string, index int) Tuple {
	return Tuple{name, index}
}

var KNIFES = []Tuple{
	of("Bayonet", 500),
	of("Classic Knife", 503),
	of("Flip Knife", 505),
	of("Gut Knife", 506),
	of("Karambit", 507),
	of("M9 Bayonet", 508),
	of("Huntsman", 509),
	of("Falchion", 512),
	of("Bowie", 514),
	of("Butterfly", 515),
	of("Shadow Daggers", 516),
	of("PARACORD", 517),
	of("Survival", 518),
	of("Ursus", 519),
	of("Navaja", 520),
	of("Nomad", 521),
	of("Stiletto", 522),
	of("Talon Knife", 523),
	of("Skeleton", 525),
	of("Kukri", 526),
}

// PISTOLS
var PISTOLS = []Tuple{
	of("DEAGLE", 1),
	of("DUELIES", 2),
	of("Five-Seven", 3),
	of("Glock", 4),
	of("TEC-9", 30),
	of("P2000", 32),
	of("P250", 36),
	of("USP", 61),
	of("CZ", 63),
	of("R8", 64),
}

// RIFLES
var RIFLES = []Tuple{
	of("AK-47", 7),
	of("AUG", 8),
	of("AWP", 9),
	of("FAMAS", 10),
	of("G3SG1", 11),
	of("GALIL", 13),
	of("M4A4", 16),
	of("SCAR-20", 38),
	of("SG-553", 39),
	of("Scout", 40),
	of("M4A1-S", 60),
}

// MGs
var MGs = []Tuple{
	of("M249", 14),
	of("NEGEV", 28),
}

// SMGs
var SMGs = []Tuple{
	of("MAC-10", 17),
	of("P90", 19),
	of("MP5", 23),
	of("UMP", 24),
	of("BIZON", 26),
	of("MP7", 33),
	of("MP9", 34),
}

// PUMPS
var PUMPS = []Tuple{
	of("XM", 25),
	of("MAG-7", 27),
	of("SAWED-OFF", 29),
	of("NOVA", 35),
}

// MISC
var MISC = []Tuple{
	of("Zeus", 31),
}

// GLOVES
var GLOVES = []Tuple{
	of("Broken Fang", 4725),
	of("Bloodhound", 5027),
	of("Sport", 5030),
	of("Driver", 5031),
	of("Hand wraps", 5032),
	of("Moto", 5033),
	of("Specialist", 5034),
	of("Hydra", 5035),
}

func CreateIndices() []Tuple {
	var types []Tuple
	types = append(types, KNIFES...)
	types = append(types, PISTOLS...)
	types = append(types, RIFLES...)
	types = append(types, MGs...)
	types = append(types, SMGs...)
	types = append(types, PUMPS...)
	types = append(types, MISC...)
	types = append(types, GLOVES...)
	return types
}
