  # Fundamentos GO
  
  Repositório de estudos dos fundamentos da linguagem **Go**, organizado em
  módulos temáticos numerados.
  
  ## Tecnologias
  
  - **Go 1.23.0**
  - [`badoux/checkmail`](https://github.com/badoux/checkmail) — validação de
  e-mail (usado no módulo de pacotes)

  ## Estrutura

  Fundamentos_GO/
  ├── 1 - pacotes/         # Pacotes locais e externos
  │   └── auxiliar/
  ├── 2 - variaveis/       # Declaração e atribuição de variáveis
  ├── 3 - tipos de dados/  # Tipos primitivos do Go
  ├── 4 - funcoes/         # Funções
  ├── 5 - Operadores/      # Operadores aritméticos, relacionais e de atribuição
  └── 6 - struct/          # Structs e tipos customizados

  ## Tópicos abordados

  | Módulo | Conceitos |
  |--------|-----------|
  | 1 - pacotes | Pacotes locais, importações externas, funções exportadas |
  | 2 - variaveis | `var`, `:=`, atribuição múltipla |
  | 3 - tipos de dados | `int`, `uint`, `float`, `string`, `bool`, `rune`,
  `byte`, `error` |
  | 4 - funcoes | Declaração e uso de funções |
  | 5 - Operadores | Aritmética (`+`, `-`, `*`, `/`, `%`), relacionais (`==`,
  `!=`, `>`, `<`) |
  | 6 - struct | Definição de structs, structs aninhadas, inicialização por
  campo |

  ## Como executar

  Cada diretório é um programa Go independente. Navegue até o módulo desejado e
  execute:

  ```bash
  # Módulo com dependências externas
  cd "1 - pacotes"
  go mod tidy
  go run main.go

  # Demais módulos (arquivo único)
  cd "2 - variaveis" && go run variaveis.go
  cd "3 - tipos de dados" && go run tipos_de_dados.go
  cd "5 - Operadores" && go run operadores.go
  cd "6 - struct" && go run struck.go

  Nota: o módulo `4 - funcoes` está configurado como submódulo git e pode
  precisar de `git submodule update --init` para ser carregado.
