// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"profmugomes/michinho/internal/constants"
	"profmugomes/michinho/internal/game"
	"profmugomes/michinho/internal/state"
	"profmugomes/michinho/internal/theme"
	"profmugomes/michinho/internal/ui"
)

func main() {
	// Cria a aplicação
	a := app.NewWithID("com.mugomes.michinho")

	// Configura o tema customizado
	a.Settings().SetTheme(&theme.MichinhoTheme{})

	// Cria a janela
	w := a.NewWindow(constants.GameTitle)
	w.Resize(fyne.NewSize(constants.WindowWidth, constants.WindowHeight))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	// Estado da aplicação
	appState := state.NewAppState()

	// Carrega o nível salvo na inicialização: sem isso, "Salvar Jogo" e
	// "Sair" sobrescreveriam o progresso anterior com o nível padrão (1)
	appState.LoadSavedLevel()

	// Variável para o game widget atual
	var currentGameWidget *ui.GameWidget

	// Variáveis para as funções (declaradas primeiro, atribuídas depois)
	var showMenu func()
	var showGame func(level int)
	var showGameOver func(score int)
	var showVictory func(score int, currentLevel int)
	var showAbout func()

	// Função para mostrar o menu
	showMenu = func() {
		// Para o game loop se estiver rodando
		if currentGameWidget != nil {
			currentGameWidget.StopLoop()
			currentGameWidget = nil
		}

		w.SetContent(ui.BuildMenuScreen(func(action ui.MenuAction) {
			switch action {
			case ui.ActionContinue:
				level := state.LoadGame()
				showGame(level)
			case ui.ActionNew:
				showGame(1)
			case ui.ActionSave:
				level := appState.GetLevel()
				err := state.SaveGame(level)
				if err != nil {
					dialog.ShowError(err, w)
				} else {
					dialog.ShowInformation("Sucesso", "Jogo salvo no nível "+string(rune(level+'0')), w)
				}
			case ui.ActionAbout:
				showAbout()
			case ui.ActionExit:
				// Salva o jogo antes de sair
				level := appState.GetLevel()
				_ = state.SaveGame(level)
				w.Close()
			}
		}))
	}

	// Função para mostrar game over
	showGameOver = func(score int) {
		w.SetContent(ui.BuildGameOverScreen(score, showMenu))
	}

	// Função para mostrar vitória
	showVictory = func(score int, currentLevel int) {
		onNextLevel := func() {
			nextLevel := currentLevel + 1
			if nextLevel > game.MaxLevel() {
				nextLevel = 1
			}
			showGame(nextLevel)
		}
		w.SetContent(ui.BuildVictoryScreen(score, currentLevel, onNextLevel, showMenu))
	}

	// Função para mostrar sobre
	showAbout = func() {
		w.SetContent(ui.BuildAboutScreen(showMenu))
	}

	// Função para mostrar o jogo
	showGame = func(level int) {
		// Para o game loop anterior se existir
		if currentGameWidget != nil {
			currentGameWidget.StopLoop()
		}

		// Configura os callbacks
		onGameOver := func(score int) {
			appState.GameOver(score)
			fyne.Do(func() {
				showGameOver(score)
			})
		}

		onVictory := func(score int) {
			appState.Victory(score)
			fyne.Do(func() {
				showVictory(score, level)
			})
		}

		onESC := func() {
			if currentGameWidget != nil {
				currentGameWidget.Pause()
			}
			dialog.ShowConfirm(
				"Voltar ao Menu",
				"Deseja voltar para o menu?",
				func(ok bool) {
					if ok {
						// Salva o jogo antes de voltar ao menu
						_ = state.SaveGame(appState.GetLevel())
						showMenu()
					} else {
						// Retoma o jogo
						if currentGameWidget != nil {
							currentGameWidget.Resume()
						}
					}
				},
				w,
			)
		}

		// Cria o engine do jogo (único, compartilhado entre widget e estado)
		config := game.GetLevel(level)
		engine := game.NewEngine(config)
		engine.OnGameOver = onGameOver
		engine.OnVictory = onVictory

		// Cria o widget do jogo
		currentGameWidget = ui.NewGameWidget(engine, onESC, onGameOver, onVictory)

		// Inicia o MESMO engine usado pelo widget (antes criava um segundo
		// engine e o do widget ficava parado com Running=false)
		appState.StartGame(level, engine)
		w.SetContent(currentGameWidget)

		// Dá foco ao widget para receber eventos de teclado
		w.Canvas().Focus(currentGameWidget)

		currentGameWidget.StartLoop()
	}

	// Fecha a janela corretamente
	w.SetOnClosed(func() {
		level := appState.GetLevel()
		_ = state.SaveGame(level)
	})

	// Mostra o menu inicial
	showMenu()

	// Executa a aplicação
	w.ShowAndRun()
}
