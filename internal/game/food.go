// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

import (
	"math/rand"
)

// SpawnFood gera comida em uma posição aleatória não ocupada
func SpawnFood(worm *Worm, barriers []Barrier, gridW, gridH int) Food {
	// Lista de posições válidas
	var validPositions []Point

	for x := 0; x < gridW; x++ {
		for y := 0; y < gridH; y++ {
			p := Point{X: x, Y: y}
			if !IsOccupied(p, worm, barriers) {
				validPositions = append(validPositions, p)
			}
		}
	}

	// Se não há posições válidas (jogo ganho), retorna posição inválida
	if len(validPositions) == 0 {
		return Food{Position: Point{X: -1, Y: -1}}
	}

	// Escolhe uma posição aleatória
	idx := rand.Intn(len(validPositions))
	return Food{Position: validPositions[idx]}
}
