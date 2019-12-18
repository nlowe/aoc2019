package day17

type scaffold struct {
	x int
	y int

	s *scaffolding
}

func (s *scaffold) alignment() int {
	return s.x * s.y
}

type scaffolding struct {
	robotX    int
	robotY    int
	robotFace rune

	m map[int]map[int]*scaffold
}

func (s *scaffolding) get(x, y int) *scaffold {
	if col, ok := s.m[x]; ok {
		if result, ok := col[y]; ok {
			return result
		}
	}

	return nil
}

func (s *scaffolding) set(x, y int) {
	if _, ok := s.m[x]; !ok {
		s.m[x] = map[int]*scaffold{}
	}

	s.m[x][y] = &scaffold{x, y, s}
}

func (s *scaffolding) clear(x, y int) {
	delete(s.m[x], y)
}

func (s *scaffolding) neighbors(x, y int) (result []*scaffold) {
	for _, delta := range []struct {
		X int
		Y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if v := s.get(x+delta.X, y+delta.Y); v != nil {
			result = append(result, v)
		}
	}

	return
}

func (s *scaffolding) isIntersection(x, y int) bool {
	if s.get(x, y) == nil {
		return false
	}

	return len(s.neighbors(x, y)) == 4
}
