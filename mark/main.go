// Command mark writes the go-sciio mark, svg/color/go-sciio.svg:
//
//	cd mark && go run . > ../svg/color/go-sciio.svg
//
// go-sciio is scientific file formats, and the mark is what I/O means: data
// going out along one line and coming back along another, two opposite arrows.
//
// Nothing is placed by hand. The constants below give the length of the shafts,
// the size of the heads and the distance between the two lines; the program
// derives every coordinate from them and centres the drawing on the square.
// Change a size and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#8FA3BF" // the gradient's light stop
	dark  = "#3E516B" // its dark stop

	stroke = 13.0 // the family's line width

	shaft = 148.0 // the length of each arrow's shaft, between stroke centres
	head  = 30.0  // how far the head's wings reach back from the tip, and to each side
	apart = 72.0  // the distance between the two shafts
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

// arrow draws a shaft at height y from x0 to x1 and a head at x1; the head
// points the way the shaft goes.
func arrow(x0, x1, y float64) string {
	back := -head
	if x1 < x0 {
		back = head
	}
	return fmt.Sprintf(`<path d="M%s %sH%sM%s %sL%s %sL%s %s"/>`,
		num(x0), num(y), num(x1),
		num(x1+back), num(y-head), num(x1), num(y), num(x1+back), num(y+head))
}

func main() {
	left, right := 128-shaft/2, 128+shaft/2
	upper, lower := 128-apart/2, 128+apart/2

	paths := []string{
		arrow(left, right, upper), // out
		arrow(right, left, lower), // and back
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-sciio">
  <defs><linearGradient id="scgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#scgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(paths, "\n    "))
}
