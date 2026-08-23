// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package state

import (
	"os"
	"path/filepath"
	"testing"

	"profmugomes/michinho/internal/constants"
	"profmugomes/michinho/internal/game"
)

func TestNewAppState(t *testing.T) {
	appState := NewAppState()

	if appState == nil {
		t.Fatal("NewAppState returned nil")
	}

	if appState.CurrentScreen != ScreenMenu {
		t.Error("Initial screen should be ScreenMenu")
	}

	if appState.CurrentLevel != 1 {
		t.Error("Initial level should be 1")
	}
}

func TestTransition(t *testing.T) {
	appState := NewAppState()

	appState.Transition(ScreenPlaying)
	if appState.GetCurrentScreen() != ScreenPlaying {
		t.Error("Should transition to ScreenPlaying")
	}

	appState.Transition(ScreenGameOver)
	if appState.GetCurrentScreen() != ScreenGameOver {
		t.Error("Should transition to ScreenGameOver")
	}
}

func TestSetGetLevel(t *testing.T) {
	appState := NewAppState()

	appState.SetLevel(2)
	if appState.GetLevel() != 2 {
		t.Error("Should set and get level 2")
	}
}

func TestSetGetScore(t *testing.T) {
	appState := NewAppState()

	appState.SetScore(10)
	if appState.GetScore() != 10 {
		t.Error("Should set and get score 10")
	}
}

// TestStartGameUsesSameEngine é um teste de regressão: o StartGame deve iniciar
// o MESMO engine recebido (o mesmo usado pelo widget da UI). Antes, o StartGame
// criava um segundo engine internamente e o engine do widget ficava parado
// com Running=false, impedindo a minhoca de se mover.
func TestStartGameUsesSameEngine(t *testing.T) {
	appState := NewAppState()

	config := game.GetLevel(1)
	engine := game.NewEngine(config)

	appState.StartGame(1, engine)

	// O estado deve apontar para o MESMO engine
	if appState.Engine != engine {
		t.Fatal("AppState.Engine should be the same engine passed to StartGame")
	}

	// O engine do widget deve estar rodando (senão o game loop ignora os ticks)
	if !engine.Running {
		t.Fatal("Engine must be Running after StartGame (regression: worm would not move)")
	}

	if engine.Paused {
		t.Error("Engine must not be Paused after StartGame")
	}

	// Ticks devem mover a minhoca de fato
	startX := engine.Worm.Body[0].X
	engine.Tick()
	if engine.Worm.Body[0].X == startX {
		t.Error("Worm position should change after Tick when game started")
	}
}

func TestSaveLoadGame(t *testing.T) {
	// Salva o jogo
	err := SaveGame(2)
	if err != nil {
		t.Fatalf("SaveGame failed: %v", err)
	}

	// Verifica se o arquivo foi criado
	homeDir, _ := os.UserHomeDir()
	savePath := filepath.Join(homeDir, constants.SaveFileName)
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		t.Error("Save file was not created")
	}

	// Carrega o jogo
	level := LoadGame()
	if level != 2 {
		t.Errorf("Expected level 2, got %d", level)
	}

	// Limpa
	os.Remove(savePath)
}

func TestLoadGameNoFile(t *testing.T) {
	// Remove arquivo de save se existir
	homeDir, _ := os.UserHomeDir()
	savePath := filepath.Join(homeDir, constants.SaveFileName)
	os.Remove(savePath)

	// Deve retornar nível 1 como fallback
	level := LoadGame()
	if level != 1 {
		t.Errorf("Expected level 1 when no save file, got %d", level)
	}
}

func TestHasSaveGame(t *testing.T) {
	// Remove arquivo se existir
	homeDir, _ := os.UserHomeDir()
	savePath := filepath.Join(homeDir, constants.SaveFileName)
	os.Remove(savePath)

	if HasSaveGame() {
		t.Error("HasSaveGame should return false when no file exists")
	}

	// Cria arquivo
	SaveGame(1)
	if !HasSaveGame() {
		t.Error("HasSaveGame should return true when file exists")
	}

	// Limpa
	os.Remove(savePath)
}

// TestLoadSavedLevel é um teste de regressão: a aplicação deve carregar o
// nível salvo na inicialização. Sem isso, "Salvar Jogo" e "Sair" sobrescreviam
// o progresso do jogador com o nível padrão (1).
func TestLoadSavedLevel(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	savePath := filepath.Join(homeDir, constants.SaveFileName)
	defer os.Remove(savePath)

	// Sem save: deve carregar nível 1
	os.Remove(savePath)
	appState := NewAppState()
	if got := appState.LoadSavedLevel(); got != 1 {
		t.Errorf("Without save file, expected level 1, got %d", got)
	}
	if appState.GetLevel() != 1 {
		t.Error("AppState level should be 1 without save file")
	}

	// Com save no nível 2: novo AppState deve refletir o nível salvo
	if err := SaveGame(2); err != nil {
		t.Fatalf("SaveGame failed: %v", err)
	}
	appState2 := NewAppState()
	if got := appState2.LoadSavedLevel(); got != 2 {
		t.Errorf("With saved level 2, expected LoadSavedLevel=2, got %d", got)
	}
	if appState2.GetLevel() != 2 {
		t.Error("AppState level should reflect the saved level (regression: exit would overwrite progress)")
	}
}
