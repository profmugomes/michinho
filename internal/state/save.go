// Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

// Licensed under the PolyForm Strict License 1.0.0.
// See LICENSE.md for details.

package state

import (
	"encoding/json"
	"os"
	"path/filepath"

	"profmugomes/michinho/internal/constants"
	"profmugomes/michinho/internal/game"
)

// SaveData representa os dados salvos do jogo
type SaveData struct {
	Level int `json:"level"`
}

// saveFilePath retorna o caminho completo do arquivo de save
func saveFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	return filepath.Join(homeDir, constants.SaveFileName)
}

// SaveGame salva o nível atual em um arquivo JSON
func SaveGame(level int) error {
	data := SaveData{
		Level: level,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(saveFilePath(), jsonData, 0644)
}

// LoadGame carrega o nível salvo do arquivo JSON
// Retorna o nível salvo ou 1 como fallback
func LoadGame() int {
	filePath := saveFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		// Arquivo não existe ou erro de leitura
		return 1
	}

	var saveData SaveData
	if err := json.Unmarshal(data, &saveData); err != nil {
		// JSON inválido
		return 1
	}

	// Valida o nível
	if saveData.Level < 1 || saveData.Level > game.MaxLevel() {
		return 1
	}

	return saveData.Level
}

// HasSaveGame verifica se existe um arquivo de save
func HasSaveGame() bool {
	_, err := os.Stat(saveFilePath())
	return err == nil
}
