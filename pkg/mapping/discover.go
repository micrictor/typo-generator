package mapping

import "fmt"

type Mapping struct {
	KeyboardLayout Layout
	LayoutName     string
}

// Create a Mapping object for a given layout name
func New(layoutName string) (*Mapping, error) {
	layout, ok := keyboardLayouts[layoutName]
	if !ok {
		return nil, fmt.Errorf("failed to find keyboard layout type %s", layoutName)
	}

	return &Mapping{
		KeyboardLayout: layout,
		LayoutName:     layoutName,
	}, nil
}

// Given a character, find its position on the initialized keyboard layout
func (m *Mapping) GetPosition(key rune) (KeyPosition, error) {
	for row_number, row := range m.KeyboardLayout {
		for col_number, keyValue := range row {
			if keyValue == key {
				return KeyPosition{row_number, col_number}, nil
			}
		}
	}

	return KeyPosition{}, fmt.Errorf("failed to find key %c in layout %s", key, m.LayoutName)
}

// Given a character, generate likely typos for it based on the initialized key map
func (m *Mapping) FindTypos(key rune) ([]rune, error) {
	var typos []rune
	pos, err := m.GetPosition(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get typos: %v", err)
	}

	// First go up and down
	if pos.x > 0 {
		typos = append(typos, m.KeyboardLayout[pos.x-1][pos.y])
	}
	if pos.x < len(m.KeyboardLayout)-1 {
		typos = append(typos, m.KeyboardLayout[pos.x+1][pos.y])
	}

	// Second go left and right
	if pos.y > 0 {
		typos = append(typos, m.KeyboardLayout[pos.x][pos.y-1])
	}
	if pos.y < len(m.KeyboardLayout[pos.x])-1 {
		typos = append(typos, m.KeyboardLayout[pos.x][pos.y+1])
	}

	return typos, nil
}
