package mapping

type Layout [][]rune

type KeyPosition struct {
	x int
	y int
}

var keyboardLayouts = map[string]Layout{
	"qwerty": {
		[]rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		[]rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		[]rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	},
}
