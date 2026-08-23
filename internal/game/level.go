// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

import (
	"time"

	"profmugomes/michinho/internal/constants"
)

// LevelConfig define a configuração de um nível
type LevelConfig struct {
	Number           int
	Speed            time.Duration
	GridWidth        int
	GridHeight       int
	Barriers         [][]Point
	VictoryThreshold float32
}

// Levels é o mapa de configurações de níveis disponíveis
var Levels = map[int]LevelConfig{
	1: {
		Number:           1,
		Speed:            constants.DefaultTickSpeed,
		GridWidth:        constants.GridCols,
		GridHeight:       constants.GridRows,
		Barriers:         nil, // Sem barreiras no nível 1
		VictoryThreshold: constants.VictoryThreshold,
	},
	2: {
		Number:    2,
		Speed:     150 * time.Millisecond, // Mais rápido
		GridWidth: constants.GridCols,
		GridHeight: constants.GridRows,
		Barriers: [][]Point{
			// Barreira vertical central (20px altura × 7px largura)
			// No grid: 7 colunas × 10 linhas (considerando cellSize ~20px)
			// Posicionada no centro do grid
			{
				{X: 17, Y: 10}, {X: 17, Y: 11}, {X: 17, Y: 12}, {X: 17, Y: 13}, {X: 17, Y: 14},
				{X: 17, Y: 15}, {X: 17, Y: 16}, {X: 17, Y: 17}, {X: 17, Y: 18}, {X: 17, Y: 19},
				{X: 18, Y: 10}, {X: 18, Y: 11}, {X: 18, Y: 12}, {X: 18, Y: 13}, {X: 18, Y: 14},
				{X: 18, Y: 15}, {X: 18, Y: 16}, {X: 18, Y: 17}, {X: 18, Y: 18}, {X: 18, Y: 19},
				{X: 19, Y: 10}, {X: 19, Y: 11}, {X: 19, Y: 12}, {X: 19, Y: 13}, {X: 19, Y: 14},
				{X: 19, Y: 15}, {X: 19, Y: 16}, {X: 19, Y: 17}, {X: 19, Y: 18}, {X: 19, Y: 19},
				{X: 20, Y: 10}, {X: 20, Y: 11}, {X: 20, Y: 12}, {X: 20, Y: 13}, {X: 20, Y: 14},
				{X: 20, Y: 15}, {X: 20, Y: 16}, {X: 20, Y: 17}, {X: 20, Y: 18}, {X: 20, Y: 19},
				{X: 21, Y: 10}, {X: 21, Y: 11}, {X: 21, Y: 12}, {X: 21, Y: 13}, {X: 21, Y: 14},
				{X: 21, Y: 15}, {X: 21, Y: 16}, {X: 21, Y: 17}, {X: 21, Y: 18}, {X: 21, Y: 19},
				{X: 22, Y: 10}, {X: 22, Y: 11}, {X: 22, Y: 12}, {X: 22, Y: 13}, {X: 22, Y: 14},
				{X: 22, Y: 15}, {X: 22, Y: 16}, {X: 22, Y: 17}, {X: 22, Y: 18}, {X: 22, Y: 19},
				{X: 23, Y: 10}, {X: 23, Y: 11}, {X: 23, Y: 12}, {X: 23, Y: 13}, {X: 23, Y: 14},
				{X: 23, Y: 15}, {X: 23, Y: 16}, {X: 23, Y: 17}, {X: 23, Y: 18}, {X: 23, Y: 19},
			},
		},
		VictoryThreshold: 0.85,
	},
}

// GetLevel retorna a configuração de um nível
func GetLevel(number int) LevelConfig {
	if level, ok := Levels[number]; ok {
		return level
	}
	// Retorna nível 1 como fallback
	return Levels[1]
}

// MaxLevel retorna o número do maior nível disponível
func MaxLevel() int {
	max := 1
	for k := range Levels {
		if k > max {
			max = k
		}
	}
	return max
}
