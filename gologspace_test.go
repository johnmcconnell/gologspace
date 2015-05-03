package gologspace

import(
	"testing"
)

func TestInterface(t *testing.T) {
	spaces := []Space{LogSpace{}, NumSpace{}}
	x :=  50.0
	y := 25.0
	for _, space := range spaces {
		x = space.Enter(x)
		y = space.Enter(y)
	}
}
