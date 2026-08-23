// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"profmugomes/michinho/internal/theme"
)

// MenuAction representa as ações do menu
type MenuAction string

const (
	ActionContinue MenuAction = "continue"
	ActionNew      MenuAction = "new"
	ActionSave     MenuAction = "save"
	ActionAbout    MenuAction = "about"
	ActionExit     MenuAction = "exit"
)

// BuildMenuScreen constrói a tela do menu principal
func BuildMenuScreen(onAction func(MenuAction)) fyne.CanvasObject {
	// Título do jogo
	title := canvas.NewText("Michinho", theme.TextColor)
	title.TextSize = 48
	title.Alignment = fyne.TextAlignCenter

	// Botões do menu
	btnContinue := widget.NewButton("Continuar Jogo", func() {
		onAction(ActionContinue)
	})

	btnNew := widget.NewButton("Novo Jogo", func() {
		onAction(ActionNew)
	})

	btnSave := widget.NewButton("Salvar Jogo", func() {
		onAction(ActionSave)
	})

	btnAbout := widget.NewButton("Sobre o Jogo", func() {
		onAction(ActionAbout)
	})

	btnExit := widget.NewButton("Sair", func() {
		onAction(ActionExit)
	})

	// Layout dos botões
	buttons := container.NewVBox(
		btnContinue,
		btnNew,
		btnSave,
		btnAbout,
		btnExit,
	)

	// Centraliza tudo
	return container.NewCenter(
		container.NewVBox(
			title,
			widget.NewLabel(""), // Espaçamento
			buttons,
		),
	)
}
