type Dir int

const (
	East Dir = iota
	North
	West
	South
)

var delta = [4][2]int{
	East:  {1, 0},
	North: {0, 1},
	West:  {-1, 0},
	South: {0, -1},
}

type Robot struct {
	w, h int
	x, y int
	dir  Dir
}

func Constructor(width int, height int) Robot {
	return Robot{w: width, h: height, x: 0, y: 0, dir: East}
}

func (r *Robot) Step(num int) {
	if r.w == 1 && r.h == 1 {
		return
	}

	perim := 2*(r.w+r.h) - 4
	num %= perim

	if num == 0 {
		num = perim
	}

	for num > 0 {
		d := delta[r.dir]
		nx := r.x + d[0]
		ny := r.y + d[1]

		if nx < 0 || nx >= r.w || ny < 0 || ny >= r.h {
			r.dir = (r.dir + 1) % 4
			continue
		}

		r.x = nx
		r.y = ny
		num--
	}
}

func (r *Robot) GetPos() []int {
	return []int{r.x, r.y}
}

func (r *Robot) GetDir() string {
	switch r.dir {
	case East:
		return "East"
	case North:
		return "North"
	case West:
		return "West"
	case South:
		return "South"
	default:
		return ""
	}
}