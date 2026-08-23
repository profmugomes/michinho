# Michinho

[![License](https://img.shields.io/badge/license-PolyForm%20Strict%201.0.0-5351FB)](LICENSE.md)

Um jogo da minhoca desenvolvido em Go com o toolkit Fyne.

## Descrição

Michinho é um jogo clássico da minhoca onde você controla uma minhoca que deve comer bolinhas para crescer. Cuidado para não bater no próprio corpo ou nas barreiras!

## Funcionalidades

- Movimentação suave com setas do teclado
- Sistema de níveis (2 níveis disponíveis)
- Barreiras no nível 2
- Sistema de save/load automático
- Interface gráfica com fonte Chilanka
- Efeitos visuais com cores personalizadas

## Requisitos

- Go 1.22 ou superior
- Sistema operacional com suporte a GUI (Linux)
- Graphics: GPU with OpenGL support (required for 3D acceleration)

## Controles

- **Setas do teclado**: Movimenta a minhoca
- **ESC**: Abre o menu de pausa
- **Mouse**: Interage com os botões do menu

## Estrutura do Projeto

```
Michinho/
├── main.go                    # Ponto de entrada
├── go.mod / go.sum            # Módulos Go
├── instrucoes.md              # Documento de requisitos
├── README.md                  # Esta documentação
├── LICENSE.md                 # Licença do projeto
├── internal/
│   ├── constants/             # Constantes do jogo
│   ├── game/                  # Lógica do jogo
│   ├── state/                 # Gerenciamento de estado
│   ├── ui/                    # Interface do usuário
│   └── theme/                 # Tema customizado
└── resources/
    └── Chilanka-Regular.ttf   # Fonte do jogo
```

## Níveis

1. **Nível 1**: Sem barreiras, velocidade padrão
2. **Nível 2**: Com barreiras, velocidade aumentada

## Recursos Utilizados de Terceiros

- [Fyne](https://fyne.io/) - Toolkit GUI para Go
- [Chilanka](https://fonts.google.com/specimen/Chilanka) - Fonte utilizada no jogo

## 👤 Author

**Murilo Gomes Julio**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2026 Murilo Gomes Julio. All Rights Reserved.

This project is licensed under the PolyForm Strict License 1.0.0.

### Summary

This software is available for noncommercial use only.

You may:
- ✔ Use the software for noncommercial purposes.
- ✔ Inspect and study the source code.

You may not:
- ✖ Use the software for commercial purposes.
- ✖ Modify the software.
- ✖ Redistribute the software.

See the full license terms at [LICENSE.md](LICENSE.md).

This summary is provided for convenience only.
