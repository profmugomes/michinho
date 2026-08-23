// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package state

import (
	"sync"

	"profmugomes/michinho/internal/game"
)

// Screen representa as telas do jogo
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenPlaying
	ScreenPaused
	ScreenGameOver
	ScreenVictory
	ScreenAbout
)

// AppState representa o estado global da aplicação
type AppState struct {
	mu            sync.RWMutex
	CurrentScreen Screen
	CurrentLevel  int
	Score         int
	Engine        *game.Engine
}

// NewAppState cria um novo estado de aplicação
func NewAppState() *AppState {
	return &AppState{
		CurrentScreen: ScreenMenu,
		CurrentLevel:  1,
		Score:         0,
	}
}

// Transition troca a tela atual (thread-safe)
func (s *AppState) Transition(to Screen) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CurrentScreen = to
}

// GetCurrentScreen retorna a tela atual (thread-safe)
func (s *AppState) GetCurrentScreen() Screen {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CurrentScreen
}

// SetLevel define o nível atual
func (s *AppState) SetLevel(level int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CurrentLevel = level
}

// GetLevel retorna o nível atual
func (s *AppState) GetLevel() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CurrentLevel
}

// SetScore define a pontuação atual
func (s *AppState) SetScore(score int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Score = score
}

// GetScore retorna a pontuação atual
func (s *AppState) GetScore() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Score
}

// LoadSavedLevel carrega o nível salvo do arquivo JSON e o define como
// nível atual da aplicação. Deve ser chamada na inicialização para que
// ações como "Salvar Jogo" e "Sair" não sobrescrevam o progresso anterior
// com o nível padrão.
func (s *AppState) LoadSavedLevel() int {
	level := LoadGame()
	s.SetLevel(level)
	return level
}

// StartGame inicia um novo jogo no nível especificado usando o engine fornecido.
// O engine deve ser criado previamente (com callbacks configurados) e é o mesmo
// objeto usado pela UI — assim o estado da aplicação e o widget compartilham
// exatamente o mesmo jogo.
func (s *AppState) StartGame(level int, engine *game.Engine) {
	s.SetLevel(level)
	s.SetScore(0)
	s.Engine = engine
	engine.Start()
	s.Transition(ScreenPlaying)
}

// GameOver trata o fim de jogo
func (s *AppState) GameOver(score int) {
	s.SetScore(score)
	s.Transition(ScreenGameOver)
}

// Victory trata a vitória
func (s *AppState) Victory(score int) {
	s.SetScore(score)
	s.Transition(ScreenVictory)
}
