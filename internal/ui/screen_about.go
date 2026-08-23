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

// BuildAboutScreen constrói a tela sobre o jogo.
// Usa widget.Label multilinha com alinhamento central: linhas curtas ficam
// centradas em relação à linha mais larga, evitando o aspecto desalinhado
// de vários canvas.Text empilhados.
func BuildAboutScreen(onBack func()) fyne.CanvasObject {
	// Título principal
	title := canvas.NewText("Michinho", theme.TextColor)
	title.TextSize = 44
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Subtítulo
	subtitle := canvas.NewText("Sobre o Jogo", theme.TextColor)
	subtitle.TextSize = 18
	subtitle.Alignment = fyne.TextAlignCenter

	// Bloco de informações (versão, desenvolvedor, licença)
	info := widget.NewLabel("Versão: 1.0.0\nDesenvolvedor: Mugomes\nLicença: PolyForm Strict 1.0.0")
	info.Alignment = fyne.TextAlignCenter

	// Descrição do jogo
	description := widget.NewLabel(
		"Um jogo da minhoca onde você controla uma minhoca\n" +
			"que deve comer as bolinhas para crescer.\n" +
			"Cuidado para não bater no próprio corpo!\n" +
			"No nível 2, barreiras surgem em seu caminho.")
	description.Alignment = fyne.TextAlignCenter
	description.Wrapping = fyne.TextWrapWord

	// Botão voltar
	btnBack := widget.NewButton("Voltar", onBack)

	// Layout: bloco único centralizado vertical e horizontalmente,
	// com espaçadores entre as seções
	return container.NewCenter(
		container.NewVBox(
			title,
			subtitle,
			widget.NewLabel(""), // espaçamento
			info,
			widget.NewLabel(""), // espaçamento
			description,
			widget.NewLabel(""), // espaçamento
			btnBack,
		),
	)
}
