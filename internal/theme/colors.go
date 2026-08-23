// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package theme

import "image/color"

// Cores do jogo Michinho
var (
	// WormColor é a cor da minhoca (#52c002)
	WormColor = color.RGBA{R: 82, G: 192, B: 2, A: 255}

	// BgColor é a cor de fundo (#000000)
	BgColor = color.Black

	// TextColor é a cor do texto (#FFFFFF)
	TextColor = color.White

	// FoodColor é a cor da comida (verde mais escuro para contraste)
	FoodColor = color.RGBA{R: 0, G: 200, B: 0, A: 255}

	// BarrierColor é a cor das barreiras (cinza)
	BarrierColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}
)
