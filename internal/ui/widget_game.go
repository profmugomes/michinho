// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"profmugomes/michinho/internal/game"
	"profmugomes/michinho/internal/theme"
)

// GameWidget é o widget customizado para a área de jogo
type GameWidget struct {
	widget.BaseWidget
	engine      *game.Engine
	cellSize    float32
	onESC       func()
	onGameOver  func(score int)
	onVictory   func(score int)
	stopChannel chan struct{}
}

// NewGameWidget cria um novo widget de jogo
func NewGameWidget(engine *game.Engine, onESC func(), onGameOver func(int), onVictory func(int)) *GameWidget {
	gw := &GameWidget{
		engine:      engine,
		onESC:       onESC,
		onGameOver:  onGameOver,
		onVictory:   onVictory,
		stopChannel: make(chan struct{}),
	}

	engine.OnGameOver = onGameOver
	engine.OnVictory = onVictory

	gw.ExtendBaseWidget(gw)
	return gw
}

// CreateRenderer cria o renderer para o widget
func (g *GameWidget) CreateRenderer() fyne.WidgetRenderer {
	return newGameRenderer(g)
}

// StartLoop inicia o game loop
func (g *GameWidget) StartLoop() {
	go func() {
		ticker := time.NewTicker(g.engine.Level.Speed)
		defer ticker.Stop()

		for {
			select {
			case <-g.stopChannel:
				return
			case <-ticker.C:
				if !g.engine.Running || g.engine.Paused {
					continue
				}

				g.engine.Tick()

				// Usa Refresh() do BaseWidget para redesenhar
				fyne.Do(func() {
					g.Refresh()
				})
			}
		}
	}()
}

// StopLoop para o game loop
func (g *GameWidget) StopLoop() {
	select {
	case g.stopChannel <- struct{}{}:
	default:
	}
}

// Pause pausa o jogo
func (g *GameWidget) Pause() {
	if g.engine != nil {
		g.engine.Pause()
	}
}

// Resume retoma o jogo
func (g *GameWidget) Resume() {
	if g.engine != nil {
		g.engine.Resume()
	}
}

// SetDirection define a direção de movimento
func (g *GameWidget) SetDirection(d game.Direction) {
	if g.engine != nil {
		g.engine.SetDirection(d)
	}
}

// === fyne.Focusable ===

func (g *GameWidget) TypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyUp:
		g.SetDirection(game.DirUp)
	case fyne.KeyDown:
		g.SetDirection(game.DirDown)
	case fyne.KeyLeft:
		g.SetDirection(game.DirLeft)
	case fyne.KeyRight:
		g.SetDirection(game.DirRight)
	case fyne.KeyEscape:
		if g.onESC != nil {
			g.onESC()
		}
	}
}

func (g *GameWidget) TypedRune(r rune) {}
func (g *GameWidget) FocusGained()     {}
func (g *GameWidget) FocusLost()       {}

// MinSize retorna o tamanho mínimo do widget
func (g *GameWidget) MinSize() fyne.Size {
	return fyne.NewSize(800, 600)
}

// ============================================================
// Renderer: usa um container interno para os objetos dinâmicos
// (segmentos da minhoca, comida e barreiras). O container garante
// que objetos adicionados/removidos entrem na árvore do canvas,
// tornando o crescimento da minhoca visível.
// ============================================================

type gameRenderer struct {
	widget    *GameWidget
	bg        *canvas.Rectangle
	playArea  *fyne.Container // filhos dinâmicos: segmentos, comida, barreiras
	scoreText *canvas.Text
}

func newGameRenderer(w *GameWidget) *gameRenderer {
	r := &gameRenderer{
		widget: w,
	}

	// Fundo - criado uma única vez
	r.bg = canvas.NewRectangle(theme.BgColor)

	// Área de jogo com layout manual (filhos posicionados por célula)
	r.playArea = container.NewWithoutLayout()

	// Placar - criado uma única vez
	r.scoreText = canvas.NewText("", theme.TextColor)
	r.scoreText.TextSize = 16

	return r
}

func (r *gameRenderer) Destroy() {
	r.widget.StopLoop()
}

func (r *gameRenderer) Layout(size fyne.Size) {
	r.redraw(size)
}

func (r *gameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(800, 600)
}

// Objects retorna uma lista ESTÁVEL de 3 objetos.
// Os elementos dinâmicos vivem dentro do container playArea,
// que gerencia corretamente a entrada/saída de filhos no canvas.
func (r *gameRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.playArea, r.scoreText}
}

func (r *gameRenderer) Refresh() {
	r.redraw(r.widget.Size())
}

// redraw reconstrói a área de jogo refletindo o estado atual do engine
func (r *gameRenderer) redraw(size fyne.Size) {
	engine := r.widget.engine

	// Calcula o tamanho da célula
	cellW := size.Width / float32(engine.Level.GridW)
	cellH := size.Height / float32(engine.Level.GridH)
	if cellW < cellH {
		r.widget.cellSize = cellW
	} else {
		r.widget.cellSize = cellH
	}
	cs := r.widget.cellSize

	// Fundo
	r.bg.Resize(size)

	// Reconstrói os filhos da área de jogo
	r.playArea.RemoveAll()

	// Barreiras (abaixo da minhoca)
	for _, barrier := range engine.Level.Barriers {
		for _, cell := range barrier.Cells {
			rect := canvas.NewRectangle(theme.BarrierColor)
			rect.Resize(fyne.NewSize(cs-1, cs-1))
			rect.Move(fyne.NewPos(float32(cell.X)*cs, float32(cell.Y)*cs))
			r.playArea.Add(rect)
		}
	}

	// Comida
	food := engine.Food.Position
	if food.X >= 0 && food.Y >= 0 { // posição válida
		circle := canvas.NewCircle(theme.FoodColor)
		circle.Resize(fyne.NewSize(cs-2, cs-2))
		circle.Move(fyne.NewPos(float32(food.X)*cs+1, float32(food.Y)*cs+1))
		r.playArea.Add(circle)
	}

	// Segmentos da minhoca (acima de tudo)
	for _, seg := range engine.Worm.Body {
		rect := canvas.NewRectangle(theme.WormColor)
		rect.Resize(fyne.NewSize(cs-1, cs-1))
		rect.Move(fyne.NewPos(float32(seg.X)*cs, float32(seg.Y)*cs))
		r.playArea.Add(rect)
	}

	// Placar
	r.scoreText.Text = fmt.Sprintf("Score: %d | Nível: %d", engine.Score, engine.Level.Number)
	r.scoreText.Resize(fyne.NewSize(300, 30))
	r.scoreText.Move(fyne.NewPos(10, 10))

	canvas.Refresh(r.playArea)
}
