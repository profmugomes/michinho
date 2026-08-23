// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package constants

import "time"

// Dimensões da janela
const (
	WindowWidth  = 800
	WindowHeight = 600
)

// Tamanho do grid (células)
const (
	GridCols = 40
	GridRows = 30
)

// Cores do jogo
const (
	WormColorHex = "#52c002"
	BgColorHex   = "#000000"
	TextColorHex = "#FFFFFF"
)

// Velocidade do jogo
const (
	DefaultTickSpeed = 200 * time.Millisecond
)

// Limiar de vitória (fração do grid preenchido)
const (
	VictoryThreshold = 0.8
)

// Nome do arquivo de save
const (
	SaveFileName = ".michinho_save.json"
)

// Nome do jogo
const (
	GameTitle = "Michinho"
)
