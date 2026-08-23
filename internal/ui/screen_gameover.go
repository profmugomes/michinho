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

// BuildGameOverScreen constrói a tela de fim de jogo
func BuildGameOverScreen(score int, onBack func()) fyne.CanvasObject {
	// Título
	title := canvas.NewText("Game Over", theme.TextColor)
	title.TextSize = 48
	title.Alignment = fyne.TextAlignCenter

	// Mensagem de resultado (widget.Label suporta múltiplas linhas)
	message := widget.NewLabel(fmt.Sprintf("Bolinhas comidas: %d", score))
	message.Alignment = fyne.TextAlignCenter

	// Botão voltar
	btnBack := widget.NewButton("Voltar ao Menu", onBack)

	// Layout
	return container.NewCenter(
		container.NewVBox(
			title,
			widget.NewLabel(""), // espaçamento
			message,
			widget.NewLabel(""), // espaçamento
			btnBack,
		),
	)
}
