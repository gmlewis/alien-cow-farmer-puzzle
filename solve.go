// alien-cow-farmer-puzzle solves the alien cow farmer puzzle on YouTube:
// https://www.youtube.com/watch?v=y6gioQ_1FYU
package main

import (
	"fmt"
	"log"
)

const (
	s1y       = 3
	s2y       = 11
	s1dx      = 8
	s2dx      = 8
	solvedS1x = 5
	solvedP2y = -3
)

func main() {
	p := puzState{
		p1y: 0,
		p2y: 0,
		s1x: 3,
		s2x: 3,
	}

	fmt.Printf("# Initial position: %v\n", p)

	visited := map[puzState]bool{p: true}
	solution := p.solve(nil, visited)

	for {
		results := optimize(p, solution)
		if len(results) == len(solution) {
			break
		}
		solution = results
	}
	printSolution(p, solution)

	log.Printf("Done.")
}

func printSolution(p puzState, solution []puzState) {
	last := p
	for i, pos := range solution {
		switch {
		case last.p1y != pos.p1y:
			fmt.Printf("movep1(%v)", pos.p1y)
		case last.p2y != pos.p2y:
			fmt.Printf("movep2(%v)", pos.p2y)
		case last.s1x != pos.s1x:
			fmt.Printf("moves1(%v)", pos.s1x)
		case last.s2x != pos.s2x:
			fmt.Printf("moves2(%v)", pos.s2x)
		}
		fmt.Printf("\t # Move #%v: %v\n", i+1, pos)
		last = pos
	}
}

func optimize(p puzState, solution []puzState) []puzState {
	result := make([]puzState, 0, len(solution))

	change := func(p1, p2 puzState) string {
		switch {
		case p1.p1y != p2.p1y:
			return "p1"
		case p1.p2y != p2.p2y:
			return "p2"
		case p1.s1x != p2.s1x:
			return "s1"
		case p1.s2x != p2.s2x:
			return "s2"
		}
		return "none"
	}

	for i, s := range solution {
		if i < 2 {
			result = append(result, s)
			continue
		}

		prev := result[len(result)-2]
		var prevChange string
		if i < 3 {
			prevChange = change(p, prev)
		} else {
			prevChange = change(result[len(result)-3], prev)
		}
		last := result[len(result)-1]
		lastChange := change(prev, last)

		changed := change(last, s)

		switch {
		case changed == "p1" && lastChange == "p1":
			result[len(result)-1] = s
			// log.Printf("i=%v: two p1's in a row, was %v, now %v", i, last, s)
			continue
		case changed == "p2" && lastChange == "p2":
			result[len(result)-1] = s
			// log.Printf("i=%v: two p2's in a row, was %v, now %v", i, last, s)
			continue
		case changed == "p1" && lastChange == "p2" && prevChange == "p1":
			result[len(result)-2] = puzState{p1y: s.p1y, p2y: prev.p2y, s1x: prev.s1x, s2x: prev.s2x}
			// log.Printf("i=%v: p1,p2,p1 in a row - overwrite prev, was %v, now %v", i, prev, result[len(result)-2])
			continue
		case changed == "p2" && lastChange == "p1" && prevChange == "p2":
			result[len(result)-2] = puzState{p1y: prev.p1y, p2y: s.p2y, s1x: prev.s1x, s2x: prev.s2x}
			// log.Printf("i=%v: p2,p1,p2 in a row - overwrite prev, was %v, now %v", i, prev, result[len(result)-2])
			continue
		case changed == "s1" && lastChange == "s1":
			result[len(result)-1] = s
			// log.Printf("i=%v: two s1's in a row, was %v, now %v", i, last, s)
			continue
		case changed == "s2" && lastChange == "s2":
			result[len(result)-1] = s
			// log.Printf("i=%v: two s2's in a row, was %v, now %v", i, last, s)
			continue
		}

		// changes = append(changes, changed)
		result = append(result, s)
		// log.Printf("i=%v: no optimization: %v", i, s)
		// log.Printf("optimize: i=%v: prev=%v, prevChange=%v, last=%v, lastChange=%v, s=%v, changed=%v", i, prev, prevChange, last, lastChange, s, changed)
	}

	var final []puzState
	for i, s := range result {
		if i == 0 {
			final = append(final, s)
			continue
		}
		if change(final[len(final)-1], s) == "none" {
			continue
		}
		final = append(final, s)
	}

	return final
}

func (p puzState) solve(moves []puzState, visited map[puzState]bool) []puzState {
	if p.solved() {
		return moves
	}

	pm := p.possibleMoves(visited)
	for _, m := range pm {
		if visited[m] {
			continue
		}
		visited[m] = true
		if s := m.solve(append(moves, m), visited); s != nil {
			return s
		}
	}

	return nil
}

func (p puzState) possibleMoves(visited map[puzState]bool) []puzState {
	var result []puzState

	if v, ok := p.slideTop(1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideTop(-1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideBot(1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideBot(-1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideLeft(-1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideLeft(1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideRight(-1); ok {
		result = append(result, v)
	}
	if v, ok := p.slideRight(1); ok {
		result = append(result, v)
	}

	return result
}

func (p puzState) slideTop(dx int) (puzState, bool) {
	np := puzState{p1y: p.p1y, p2y: p.p2y, s1x: p.s1x + dx, s2x: p.s2x}
	valid := p1[key{np.s1x, np.p1y + s1y}] && p2[key{np.s1x + s1dx, np.p2y + s1y}]
	return np, valid
}

func (p puzState) slideBot(dx int) (puzState, bool) {
	np := puzState{p1y: p.p1y, p2y: p.p2y, s1x: p.s1x, s2x: p.s2x + dx}
	valid := p1[key{np.s2x, np.p1y + s2y}] && p2[key{np.s2x + s2dx, np.p2y + s2y}]
	return np, valid
}

func (p puzState) slideLeft(dy int) (puzState, bool) {
	np := puzState{p1y: p.p1y + dy, p2y: p.p2y, s1x: p.s1x, s2x: p.s2x}
	valid := p1[key{np.s1x, np.p1y + s1y}] && p1[key{np.s2x, np.p1y + s2y}]
	// if valid {
	// 	log.Printf("slideLeft(%v): old=%v, new=%v, valid = p1[%v] && p1[%v] = true", dy, p, np, key{np.s1x, np.p1y + s1y}, key{np.s2x, np.p1y + s2y})
	// }
	return np, valid
}

func (p puzState) slideRight(dy int) (puzState, bool) {
	np := puzState{p1y: p.p1y, p2y: p.p2y + dy, s1x: p.s1x, s2x: p.s2x}
	valid := p2[key{np.s1x + s1dx, np.p2y + s1y}] && p2[key{np.s2x + s2dx, np.p2y + s2y}]
	return np, valid
}

func (p puzState) solved() bool {
	return p.s1x == solvedS1x && p.p2y == solvedP2y
}

func (p puzState) String() string {
	return fmt.Sprintf("[p1y=%v,p2y=%v,s1x=%v,s2x=%v]", p.p1y, p.p2y, p.s1x, p.s2x)
}

type puzState struct {
	p1y int
	p2y int
	s1x int
	s2x int
}

type key struct {
	x int
	y int
}

type panel map[key]bool

var p1 = panel{
	{1, 0}: true,
	{5, 0}: true,
	{0, 1}: true,
	{1, 1}: true,
	{2, 1}: true,
	{4, 1}: true,
	{5, 1}: true,
	{6, 1}: true,
	{0, 2}: true,
	{2, 2}: true,
	{3, 2}: true,
	{4, 2}: true,
	{6, 2}: true,
	{0, 3}: true,
	{3, 3}: true,
	{6, 3}: true,
	{0, 4}: true,
	{6, 4}: true,
	{0, 5}: true,
	{1, 5}: true,
	{5, 5}: true,
	{6, 5}: true,
	{1, 6}: true,
	{5, 6}: true,

	{0, 8}:  true,
	{5, 8}:  true,
	{0, 9}:  true,
	{4, 9}:  true,
	{5, 9}:  true,
	{0, 10}: true,
	{1, 10}: true,
	{3, 10}: true,
	{4, 10}: true,
	{1, 11}: true,
	{2, 11}: true,
	{3, 11}: true,
	{3, 12}: true,
	{3, 13}: true,
	{4, 13}: true,
	{6, 13}: true,
	{4, 14}: true,
	{5, 14}: true,
	{6, 14}: true,
}

var p2 = panel{
	{9, 0}:  true,
	{13, 0}: true,
	{8, 1}:  true,
	{9, 1}:  true,
	{12, 1}: true,
	{13, 1}: true,
	{14, 1}: true,
	{8, 2}:  true,
	{11, 2}: true,
	{12, 2}: true,
	{14, 2}: true,
	{8, 3}:  true,
	{11, 3}: true,
	{14, 3}: true,
	{8, 4}:  true,
	{10, 4}: true,
	{11, 4}: true,
	{14, 4}: true,
	{8, 5}:  true,
	{9, 5}:  true,
	{10, 5}: true,
	{13, 5}: true,
	{14, 5}: true,
	{13, 6}: true,

	{9, 8}:   true,
	{13, 8}:  true,
	{8, 9}:   true,
	{9, 9}:   true,
	{10, 9}:  true,
	{13, 9}:  true,
	{14, 9}:  true,
	{8, 10}:  true,
	{10, 10}: true,
	{11, 10}: true,
	{14, 10}: true,
	{8, 11}:  true,
	{11, 11}: true,
	{14, 11}: true,
	{8, 12}:  true,
	{11, 12}: true,
	{14, 12}: true,
	{8, 13}:  true,
	{9, 13}:  true,
	{11, 13}: true,
	{12, 13}: true,
	{14, 13}: true,
	{9, 14}:  true,
	{12, 14}: true,
	{13, 14}: true,
	{14, 14}: true,
}
