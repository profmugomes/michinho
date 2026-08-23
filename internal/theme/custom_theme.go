// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// MichinhoTheme é o tema customizado para o jogo Michinho
type MichinhoTheme struct{}

// Color retorna a cor para o nome especificado
func (m *MichinhoTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Usa cores padrão do Fyne para a maioria dos elementos
	// O fundo preto e cores específicas são aplicados diretamente nos widgets
	return theme.DefaultTheme().Color(name, variant)
}

// Font retorna a fonte Chilanka para todos os estilos de texto
func (m *MichinhoTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceChilankaRegularTtf
}

// Icon retorna o ícone padrão para o nome especificado
func (m *MichinhoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size retorna o tamanho para o nome especificado
func (m *MichinhoTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
