package entity

type SatelliteElem struct {
	*Position
	Message []string `json:"message"`
}

type SatelliteElems []*SatelliteElem

func (s SatelliteElems) Positions() []*Position {
	positions := make([]*Position, len(s))
	for index, elem := range s {
		positions[index] = elem.Position
	}
	return positions
}

func (s SatelliteElems) Messages() [][]string {
	messages := make([][]string, len(s))
	for index, elem := range s {
		messages[index] = elem.Message
	}
	return messages
}

type Satellites struct {
	SatelliteElems `json:"satellites"`
}
