// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

import "time"

// Direction representa a direção de movimento da minhoca
type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

// Point representa uma posição no grid
type Point struct {
	X, Y int
}

// Worm representa a minhoca
type Worm struct {
	Body      []Point    // Body[0] é a cabeça
	Direction Direction  // Direção atual
	NextDir   Direction  // Direção pendente (próximo tick)
	Growing   int        // Segmentos restantes para crescer
}

// Food representa a comida no jogo
type Food struct {
	Position Point
}

// Barrier representa uma barreira no jogo
type Barrier struct {
	Cells []Point
}

// Level representa a configuração de um nível
type Level struct {
	Number           int
	Barriers         []Barrier
	Speed            time.Duration
	GridW            int
	GridH            int
	VictoryThreshold float32
}

// CollisionType representa o tipo de colisão detectada
type CollisionType int

const (
	CollisionNone CollisionType = iota
	CollisionSelf
	CollisionWall
	CollisionBarrier
	CollisionFood
)

// NewWorm cria uma minhoca nova na posição inicial
func NewWorm(startX, startY int) *Worm {
	return &Worm{
		Body: []Point{
			{X: startX, Y: startY},
			{X: startX - 1, Y: startY},
			{X: startX - 2, Y: startY},
		},
		Direction: DirRight,
		NextDir:   DirRight,
		Growing:   0,
	}
}

// Head retorna a posição da cabeça da minhoca
func (w *Worm) Head() Point {
	return w.Body[0]
}

// SetDirection define a direção de movimento (com validação de reversão)
func (w *Worm) SetDirection(d Direction) {
	// Impede reversão de 180°
	if (w.Direction == DirUp && d == DirDown) ||
		(w.Direction == DirDown && d == DirUp) ||
		(w.Direction == DirLeft && d == DirRight) ||
		(w.Direction == DirRight && d == DirLeft) {
		return
	}
	w.NextDir = d
}

// Move move a minhoca na direção atual dentro dos limites do grid.
// Ao alcançar uma borda, a minhoca atravessa para o lado oposto
// (wrap-around): conforme instrucoes.md, ela pode percorrer todo o
// espaço disponível na tela e não morre por parede.
func (w *Worm) Move(gridW, gridH int) {
	w.Direction = w.NextDir

	head := w.Head()
	var newHead Point

	switch w.Direction {
	case DirUp:
		newHead = Point{X: head.X, Y: head.Y - 1}
	case DirDown:
		newHead = Point{X: head.X, Y: head.Y + 1}
	case DirLeft:
		newHead = Point{X: head.X - 1, Y: head.Y}
	case DirRight:
		newHead = Point{X: head.X + 1, Y: head.Y}
	}

	// Wrap-around nas bordas (sempre retorna valor em [0, dimensão))
	newHead.X = ((newHead.X % gridW) + gridW) % gridW
	newHead.Y = ((newHead.Y % gridH) + gridH) % gridH

	// Adiciona nova cabeça
	w.Body = append([]Point{newHead}, w.Body...)

	// Se está crescendo, mantém a cauda; senão remove o último segmento
	if w.Growing > 0 {
		w.Growing--
	} else {
		w.Body = w.Body[:len(w.Body)-1]
	}
}

// Grow faz a minhoca crescer
func (w *Worm) Grow(amount int) {
	w.Growing += amount
}

// CollidesWithSelf verifica se a minhoca colidiu consigo mesma
func (w *Worm) CollidesWithSelf() bool {
	head := w.Head()
	for i := 1; i < len(w.Body); i++ {
		if w.Body[i] == head {
			return true
		}
	}
	return false
}

// Contains verifica se um ponto está no corpo da minhoca
func (w *Worm) Contains(p Point) bool {
	for _, seg := range w.Body {
		if seg == p {
			return true
		}
	}
	return false
}
