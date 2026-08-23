// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package game

import (
	"testing"
)

func TestNewWorm(t *testing.T) {
	worm := NewWorm(20, 15)

	if worm == nil {
		t.Fatal("NewWorm returned nil")
	}

	if len(worm.Body) != 3 {
		t.Errorf("Expected 3 body segments, got %d", len(worm.Body))
	}

	// Verifica posição inicial
	if worm.Body[0].X != 20 || worm.Body[0].Y != 15 {
		t.Errorf("Head position: expected (20,15), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}

	// Verifica direção inicial
	if worm.Direction != DirRight {
		t.Errorf("Expected direction DirRight, got %d", worm.Direction)
	}
}

func TestWormMove(t *testing.T) {
	worm := NewWorm(20, 15)

	// Move para direita
	worm.Move(40, 30)

	// Cabeça deve ter se movido
	if worm.Body[0].X != 21 || worm.Body[0].Y != 15 {
		t.Errorf("After move right: expected head at (21,15), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}
}

// TestWormWrapAround verifica que a minhoca atravessa as bordas do grid:
// conforme instrucoes.md ela percorre todo o espaço da tela e não morre por parede.
func TestWormWrapAround(t *testing.T) {
	// Borda direita → aparece na esquerda
	worm := NewWorm(39, 15)
	worm.Move(40, 30)
	if worm.Body[0].X != 0 || worm.Body[0].Y != 15 {
		t.Errorf("Wrap right edge: expected head at (0,15), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}

	// Borda esquerda → aparece na direita
	// (define a direção diretamente: SetDirection bloquearia a reversão
	// de 180° a partir de DirRight, que é o comportamento correto do jogo)
	worm = NewWorm(20, 15)
	worm.Direction = DirLeft
	worm.NextDir = DirLeft
	worm.Move(40, 30) // vai para 19
	for i := 0; i < 19; i++ {
		worm.Move(40, 30)
	}
	if worm.Body[0].X != 0 {
		t.Fatalf("Expected head X=0 before left wrap, got %d", worm.Body[0].X)
	}
	worm.Move(40, 30)
	if worm.Body[0].X != 39 || worm.Body[0].Y != 15 {
		t.Errorf("Wrap left edge: expected head at (39,15), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}

	// Borda superior → aparece embaixo
	worm = NewWorm(20, 0)
	worm.SetDirection(DirUp)
	worm.Move(40, 30)
	if worm.Body[0].X != 20 || worm.Body[0].Y != 29 {
		t.Errorf("Wrap top edge: expected head at (20,29), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}

	// Borda inferior → aparece em cima
	worm = NewWorm(20, 29)
	worm.SetDirection(DirDown)
	worm.Move(40, 30)
	if worm.Body[0].X != 20 || worm.Body[0].Y != 0 {
		t.Errorf("Wrap bottom edge: expected head at (20,0), got (%d,%d)", worm.Body[0].X, worm.Body[0].Y)
	}
}

func TestWormSetDirection(t *testing.T) {
	worm := NewWorm(20, 15)
	worm.Direction = DirRight

	// Não deve permitir reversão de 180°
	worm.SetDirection(DirLeft)
	if worm.NextDir != DirRight {
		t.Error("Should not allow 180° reversal")
	}

	// Deve permitir mudança válida
	worm.SetDirection(DirUp)
	if worm.NextDir != DirUp {
		t.Error("Should allow valid direction change")
	}
}

func TestWormGrow(t *testing.T) {
	worm := NewWorm(20, 15)
	initialLen := len(worm.Body)

	worm.Grow(3)
	worm.Move(40, 30) // Cresce ao mover

	if len(worm.Body) != initialLen+1 {
		t.Errorf("Expected body length %d after grow, got %d", initialLen+1, len(worm.Body))
	}
}

func TestWormCollidesWithSelf(t *testing.T) {
	worm := NewWorm(20, 15)

	// Cria uma minhoca que colide consigo mesma
	worm.Body = []Point{
		{X: 20, Y: 15},
		{X: 21, Y: 15},
		{X: 21, Y: 16},
		{X: 20, Y: 16},
		{X: 20, Y: 15}, // Colisão com a cabeça
	}

	if !worm.CollidesWithSelf() {
		t.Error("Expected self collision")
	}
}

func TestCheckCollisionWall(t *testing.T) {
	worm := NewWorm(20, 15)

	// Testa colisão com parede esquerda
	if CheckCollision(Point{X: -1, Y: 15}, nil, worm, 40, 30) != CollisionWall {
		t.Error("Expected wall collision at X=-1")
	}

	// Testa colisão com parede direita
	if CheckCollision(Point{X: 40, Y: 15}, nil, worm, 40, 30) != CollisionWall {
		t.Error("Expected wall collision at X=40")
	}

	// Testa colisão com parede superior
	if CheckCollision(Point{X: 20, Y: -1}, nil, worm, 40, 30) != CollisionWall {
		t.Error("Expected wall collision at Y=-1")
	}

	// Testa colisão com parede inferior
	if CheckCollision(Point{X: 20, Y: 30}, nil, worm, 40, 30) != CollisionWall {
		t.Error("Expected wall collision at Y=30")
	}
}

func TestCheckCollisionFood(t *testing.T) {
	head := Point{X: 20, Y: 15}
	food := Point{X: 20, Y: 15}

	if !CheckFoodCollision(head, food) {
		t.Error("Expected food collision")
	}

	food2 := Point{X: 21, Y: 15}
	if CheckFoodCollision(head, food2) {
		t.Error("Should not detect food collision at different position")
	}
}

func TestSpawnFood(t *testing.T) {
	worm := NewWorm(20, 15)
	food := SpawnFood(worm, nil, 40, 30)

	// Comida não deve estar no corpo da minhoca
	if worm.Contains(food.Position) {
		t.Error("Food spawned on worm body")
	}

	// Comida deve estar dentro do grid
	if food.Position.X < 0 || food.Position.X >= 40 || food.Position.Y < 0 || food.Position.Y >= 30 {
		t.Errorf("Food outside grid: %v", food.Position)
	}
}

func TestEngine(t *testing.T) {
	config := GetLevel(1)
	engine := NewEngine(config)

	if engine == nil {
		t.Fatal("NewEngine returned nil")
	}

	// Engine deve estar parado inicialmente
	if engine.Running {
		t.Error("Engine should not be running initially")
	}

	// Inicia o jogo
	engine.Start()
	if !engine.Running {
		t.Error("Engine should be running after Start()")
	}

	// Tick deve funcionar
	engine.Tick()

	// Minhoca deve ter se movido (posição inicial era centro do grid)
	startX := config.GridWidth / 2
	if engine.Worm.Body[0].X != startX+1 {
		t.Errorf("Expected head at X=%d after tick, got %d", startX+1, engine.Worm.Body[0].X)
	}
}

func TestEnginePauseResume(t *testing.T) {
	config := GetLevel(1)
	engine := NewEngine(config)
	engine.Start()

	// Pausa
	engine.Pause()
	if !engine.Paused {
		t.Error("Engine should be paused")
	}

	// Tick não deve fazer nada quando pausado
	initialX := engine.Worm.Body[0].X
	engine.Tick()
	if engine.Worm.Body[0].X != initialX {
		t.Error("Worm should not move when paused")
	}

	// Resume
	engine.Resume()
	if engine.Paused {
		t.Error("Engine should not be paused after resume")
	}
}

func TestLevelConfig(t *testing.T) {
	// Testa nível 1
	level1 := GetLevel(1)
	if level1.Number != 1 {
		t.Errorf("Expected level 1, got %d", level1.Number)
	}
	if len(level1.Barriers) != 0 {
		t.Error("Level 1 should have no barriers")
	}

	// Testa nível 2
	level2 := GetLevel(2)
	if level2.Number != 2 {
		t.Errorf("Expected level 2, got %d", level2.Number)
	}
	if len(level2.Barriers) == 0 {
		t.Error("Level 2 should have barriers")
	}

	// Testa nível inexistente (deve retornar nível 1)
	level99 := GetLevel(99)
	if level99.Number != 1 {
		t.Errorf("Invalid level should fallback to 1, got %d", level99.Number)
	}
}

func TestMaxLevel(t *testing.T) {
	max := MaxLevel()
	if max != 2 {
		t.Errorf("Expected max level 2, got %d", max)
	}
}

// TestEngineEatsFood é um teste de regressão: quando a cabeça da minhoca
// alcança a posição da comida, ela deve comer (score++), crescer e uma nova
// comida deve surgir. Antes, o Tick nunca avaliava a comida e a minhoca
// passava por cima dela sem comê-la.
func TestEngineEatsFood(t *testing.T) {
	config := GetLevel(1)
	engine := NewEngine(config)
	engine.Start()

	// Posiciona a comida exatamente à frente da cabeça (minhoca inicia indo para a direita)
	head := engine.Worm.Head()
	engine.Food = Food{Position: Point{X: head.X + 1, Y: head.Y}}

	initialScore := engine.Score
	initialLen := len(engine.Worm.Body)

	// Um tick deve mover a cabeça para cima da comida e comê-la
	engine.Tick()

	if engine.Score != initialScore+1 {
		t.Errorf("Expected score %d after eating, got %d", initialScore+1, engine.Score)
	}

	// A minhoca ainda não cresceu visivelmente neste tick (cresce no próximo),
	// mas Growing deve estar marcado
	if engine.Worm.Growing != 1 {
		t.Errorf("Expected Growing=1 after eating, got %d", engine.Worm.Growing)
	}

	// Nova comida deve ter sido gerada em posição válida
	newFood := engine.Food.Position
	if newFood.X < 0 || newFood.X >= config.GridWidth || newFood.Y < 0 || newFood.Y >= config.GridHeight {
		t.Errorf("New food outside grid: %v", newFood)
	}
	if engine.Worm.Contains(newFood) {
		t.Error("New food spawned on worm body")
	}

	// Após o próximo tick, o corpo deve ter crescido em 1 segmento
	engine.Tick()
	if len(engine.Worm.Body) != initialLen+1 {
		t.Errorf("Expected body length %d after growth, got %d", initialLen+1, len(engine.Worm.Body))
	}
}

// TestCheckCollisionNeverReturnsFood garante que CheckCollision trata apenas
// colisões fatais; a comida é avaliada separadamente via CheckFoodCollision.
func TestCheckCollisionNeverReturnsFood(t *testing.T) {
	worm := NewWorm(20, 15)

	// Mesmo com a cabeça sobre "algo", CheckCollision não conhece comida
	result := CheckCollision(Point{X: 20, Y: 15}, nil, worm, 40, 30)
	if result == CollisionFood {
		t.Error("CheckCollision should never return CollisionFood")
	}

	// CheckFoodCollision é quem detecta a comida
	if !CheckFoodCollision(Point{X: 5, Y: 5}, Point{X: 5, Y: 5}) {
		t.Error("CheckFoodCollision should detect matching positions")
	}
}
