package settlements

import (
	"fmt"
	"math/rand"
)

type Settlement struct {
	Name       string
	Size       int
	Government string
	Amenities  []string
	Atmosphere string
	Geography  string
	Character  string
}
type Settlements []*Settlement

var SettlementsList Settlements
var placingSettlements bool = false

func Toggle() bool {
	placingSettlements = !placingSettlements
	return placingSettlements
}

func (s *Settlements) NewSettlement() *Settlement {
	rand.Shuffle(len(SettAmeniities), func(i, j int) {
		SettAmeniities[i], SettAmeniities[j] = SettAmeniities[j], SettAmeniities[i]
	})
	size := SettSizes[rand.Intn(len(SettSizes))]
	rand.Shuffle(len(SettNames), func(i, j int) {
		SettNames[i], SettNames[j] = SettNames[j], SettNames[i]
	})
	i := 0
	name := SettNames[i]
	for i < len(SettNames)-1 {
		i++
		for _, s := range SettlementsList {
			if s.Name == name {
				name = SettNames[i]
			}
		}
	}
	if i > len(SettNames) {
		name = "New Settlement"
	}
	switch {
	case size <= 1000:
		SettAmeniities = SettAmeniities[:2]
	case size <= 6000:
		SettAmeniities = SettAmeniities[:3]
	case size <= 25000:
		SettAmeniities = SettAmeniities[:4]
	case size <= 100000:
		SettAmeniities = SettAmeniities[:5]
	default:
		SettAmeniities = SettAmeniities[:6]
	}
	new := Settlement{
		Name:       name,
		Size:       SettSizes[rand.Intn(len(SettSizes))],
		Atmosphere: SettAtmospheres[rand.Intn(len(SettAtmospheres))],
		Government: SettGovenments[rand.Intn(len(SettGovenments))],
		Amenities:  SettAmeniities[:3],
		Character:  SettCharacters[rand.Intn(len(SettCharacters))],
		Geography:  SettGeographies[rand.Intn(len(SettGeographies))],
	}
	*s = append(*s, &new)
	return &new
}

func (s *Settlement) SizeType() string {
	switch {
	case s.Size <= 1000:
		return "Village"
	case s.Size <= 6000:
		return "Town"
	case s.Size <= 25000:
		return "City"
	case s.Size <= 100000:
		return "Metropolis"
	default:
		return "M<egalopolis"
	}
}

func (s *Settlement) AddAmenity(amenity string) {
	for _, a := range s.Amenities {
		if a == amenity {
			return
		}
	}
	s.Amenities = append(s.Amenities, amenity)
}

func (s *Settlement) RemoveAmenity(amenity string) {
	for i, a := range s.Amenities {
		if a == amenity {
			s.Amenities = append(s.Amenities[:i], s.Amenities[i+1:]...)
			return
		}
	}
}

func (s *Settlement) String() string {
	return fmt.Sprintf("%s %s", s.Name, s.SizeType())
}
