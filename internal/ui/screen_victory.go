// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"profmugomes/michinho/internal/theme"
)

// BuildVictoryScreen constrói a tela de vitória
func BuildVictoryScreen(score int, currentLevel int, onNextLevel func(), onMenu func()) fyne.CanvasObject {
	// Título
	title := canvas.NewText("Vitória!", theme.TextColor)
	title.TextSize = 48
	title.Alignment = fyne.TextAlignCenter

	// Mensagem de resultado (widget.Label suporta múltiplas linhas)
	message := widget.NewLabel(
		fmt.Sprintf("Nível %d concluído!\nBolinhas comidas: %d", currentLevel, score))
	message.Alignment = fyne.TextAlignCenter

	// Botões
	var btnNext *widget.Button
	if currentLevel < 2 { // Máximo 2 níveis por enquanto
		btnNext = widget.NewButton("Próximo Nível", onNextLevel)
	}

	btnMenu := widget.NewButton("Voltar ao Menu", onMenu)

	// Layout
	items := []fyne.CanvasObject{
		title,
		widget.NewLabel(""), // espaçamento
		message,
		widget.NewLabel(""), // espaçamento
	}

	if btnNext != nil {
		items = append(items, btnNext)
	}
	items = append(items, btnMenu)

	return container.NewCenter(
		container.NewVBox(items...),
	)
}
