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
- Sistema operacional com suporte a GUI (Linux, macOS, Windows)

## Instalação

```bash
# Clonar o repositório
git clone https://github.com/profmugomes/michinho.git

# Entrar no diretório
cd michinho

# Compilar e executar
go run .
```

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

## Desenvolvedor

- **Nome**: Mugomes
- **GitHub**: [profmugomes](https://github.com/profmugomes)

## Licença

Este projeto está sob a licença PolyForm Strict 1.0.0. Veja o arquivo [LICENSE.md](LICENSE.md) para mais detalhes.

## Agradecimentos

- [Fyne](https://fyne.io/) - Toolkit GUI para Go
- [Chilanka](https://fonts.google.com/specimen/Chilanka) - Fonte utilizada no jogo
