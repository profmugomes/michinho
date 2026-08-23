// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

// Engine é o motor principal do jogo
type Engine struct {
	Worm      Worm
	Food      Food
	Level     Level
	Score     int
	Running   bool
	Paused    bool
	OnGameOver func(score int)
	OnVictory  func(score int)
}

// NewEngine cria um novo engine de jogo para o nível especificado
func NewEngine(levelConfig LevelConfig) *Engine {
	// Posição inicial da minhoca (centro do grid)
	startX := levelConfig.GridWidth / 2
	startY := levelConfig.GridHeight / 2

	worm := NewWorm(startX, startY)

	// Converte barriers do LevelConfig para []Barrier
	var barriers []Barrier
	for _, b := range levelConfig.Barriers {
		barriers = append(barriers, Barrier{Cells: b})
	}

	level := Level{
		Number:           levelConfig.Number,
		Barriers:         barriers,
		Speed:            levelConfig.Speed,
		GridW:            levelConfig.GridWidth,
		GridH:            levelConfig.GridHeight,
		VictoryThreshold: levelConfig.VictoryThreshold,
	}

	engine := &Engine{
		Worm:  *worm,
		Level: level,
	}

	// Gera primeira comida
	engine.Food = SpawnFood(&engine.Worm, engine.Level.Barriers, engine.Level.GridW, engine.Level.GridH)

	return engine
}

// Tick avança um frame do jogo
func (e *Engine) Tick() {
	if !e.Running || e.Paused {
		return
	}

	// Move a minhoca (com wrap-around nas bordas do grid)
	e.Worm.Move(e.Level.GridW, e.Level.GridH)

	// Verifica colisões fatais da cabeça (parede, corpo, barreira)
	head := e.Worm.Head()
	collision := CheckCollision(head, e.Level.Barriers, &e.Worm, e.Level.GridW, e.Level.GridH)

	switch collision {
	case CollisionWall, CollisionSelf, CollisionBarrier:
		e.Running = false
		if e.OnGameOver != nil {
			e.OnGameOver(e.Score)
		}
		return
	}

	// Verifica se comeu a comida
	// (CheckCollision não avalia comida; a verificação é feita aqui)
	if CheckFoodCollision(head, e.Food.Position) {
		e.Score++
		e.Worm.Grow(1) // Crescimento mínimo

		// Gera nova comida
		e.Food = SpawnFood(&e.Worm, e.Level.Barriers, e.Level.GridW, e.Level.GridH)

		// Verifica vitória
		if e.IsVictory() {
			e.Running = false
			if e.OnVictory != nil {
				e.OnVictory(e.Score)
			}
			return
		}
	}
}

// SetDirection define a direção de movimento
func (e *Engine) SetDirection(d Direction) {
	e.Worm.SetDirection(d)
}

// IsVictory verifica se o jogador venceu (preencheu o threshold do grid)
func (e *Engine) IsVictory() bool {
	totalCells := e.Level.GridW * e.Level.GridH
	occupiedCells := len(e.Worm.Body)
	fillingRatio := float32(occupiedCells) / float32(totalCells)

	return fillingRatio >= e.Level.VictoryThreshold
}

// Start inicia o jogo
func (e *Engine) Start() {
	e.Running = true
	e.Paused = false
}

// Pause pausa o jogo
func (e *Engine) Pause() {
	e.Paused = true
}

// Resume retoma o jogo
func (e *Engine) Resume() {
	e.Paused = false
}
