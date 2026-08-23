// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

// CheckCollision verifica colisão da cabeça da minhoca com um ponto
func CheckCollision(head Point, barriers []Barrier, worm *Worm, gridW, gridH int) CollisionType {
	// Verifica colisão com paredes
	if head.X < 0 || head.X >= gridW || head.Y < 0 || head.Y >= gridH {
		return CollisionWall
	}

	// Verifica colisão com o corpo (excluindo a cabeça)
	for i := 1; i < len(worm.Body); i++ {
		if worm.Body[i] == head {
			return CollisionSelf
		}
	}

	// Verifica colisão com barreiras
	for _, barrier := range barriers {
		for _, cell := range barrier.Cells {
			if cell == head {
				return CollisionBarrier
			}
		}
	}

	return CollisionNone
}

// CheckFoodCollision verifica se a cabeça da minhoca comeu a comida
func CheckFoodCollision(head Point, food Point) bool {
	return head == food
}

// IsOccupied verifica se um ponto está ocupado (minhoca, barreiras)
func IsOccupied(p Point, worm *Worm, barriers []Barrier) bool {
	// Verifica minhoca
	if worm.Contains(p) {
		return true
	}

	// Verifica barreiras
	for _, barrier := range barriers {
		for _, cell := range barrier.Cells {
			if cell == p {
				return true
			}
		}
	}

	return false
}
