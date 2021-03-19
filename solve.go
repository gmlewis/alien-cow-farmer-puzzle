// alien-cow-farmer-puzzle solves the alien cow farmer puzzle on YouTube:
// https://www.youtube.com/watch?v=y6gioQ_1FYU
package main

import (
	"fmt"
	"log"
)

const (
	s1y       = 3
	s2y       = 12
	s1dx      = 8
	s2dx      = 8
	solvedS1x = 5
	solvedP1y = -3
)

func main() {
	p := puzState{
		p1y: 0,
		p2y: 0,
		s1x: 3,
		s2x: 3,
	}

	fmt.Printf("Initial position: %v\n", p)

	visited := map[puzState]bool{p: true}
	solution := p.solve(nil, visited)

	for i, pos := range solution {
		fmt.Printf("Move #%v: %v\n", i+1, pos)
	}

	log.Printf("Done.")
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
	return np, valid
}

func (p puzState) slideRight(dy int) (puzState, bool) {
	np := puzState{p1y: p.p1y, p2y: p.p2y + dy, s1x: p.s1x, s2x: p.s2x}
	valid := p2[key{np.s1x + s1dx, np.p2y + s1y}] && p2[key{np.s2x + s2dx, np.p2y + s2y}]
	return np, valid
}

func (p puzState) solved() bool {
	return p.s1x == solvedS1x && p.p1y == solvedP1y
}

func (p puzState) String() string {
	return fmt.Sprintf("[%v,%v,%v,%v]", p.p1y, p.p2y, p.s1x, p.s2x)
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
