# GoLang

## Instalando
1. Vars de Ambiente:
  1. Verificar variáveis de ambiente: ``` $ go env ```
  1. GOROOT = aonde foi instalado (compiladores e ferramentas (SDK)) -> ``` $ go env GOROOT ```
  1. GOPATH = workspace de projetos -> ``` $ go env GOPATH ```
1. Estrutura
  1. Abrir vs: ``` $ cd ~ && code go ```
  1. pasta bin = pasta de executáveis, ao gerar um executável ele virá para esta pasta
    1. se add a pasta bin do go ao PATH do sistema, os executáveis desta estarão disponíveis como comandos
  1. pasta pkg = compilados por arquitetura de destino, sendo os ".a" os arquivos do projeto compilado
  1. pasta src = fontes, separado por mecanismos de versionamento (github, bitbucket, gitlab...)
  1. Para criar e rodar um projeto standalone __não precisa colocar no GOPATH__, quando for criar pacotes reutilizáveis é necessário criá-los dentro do GoPath
  
## ADD Visual Code ao ZSH:
1. ``` $ sudo nano ~/.zshrc ```
1. add: 
    ```zsh
        function code {
            if [[ $# = 0 ]]
            then
                open -a "Visual Studio"
            else
                local argPath="$1"
                [[ $1 = /* ]] && argPath="$1" || argPath="$PWD/${1#./}"
                open -a "Visual Studio" "$argPath"
            fi
        }
    ```
  
## Primeiro código
1. todo programa go precisa rodar dentro de um pacote
1. cada pasta com código go só pode ter uma função main
1. /fundamentos/primeiro/primeiro.go

## Comandos Go
1. ``` $ go ``` -> exibe comandos disponiveis no go
1. ``` $ go help comando ``` help de um comando especifico, ex.: ``` $ go help get ```, esse é usado para baixar pacotes do github
1. Documentação do Go offline:  ``` $ godoc -http=:6060 ```
1. ``` $ go doc cmd/vet ``` -> exibe documentação deste comando específico
1. ``` $ go build comandos.go ``` -> cria um executável: ``` $ ./comandos ```
1. ``` $ go run comandos.go ``` -> cria um executável e executa
1. Instalando dependencia (driver sql): ``` $ go get -u github.com/go-sql-driver/mysql ```, verificar se foi instalado: ``` $ ls ~/go/src/github.com/  ```
1. Executando todos arquivos .go:  
    1. acessar o pacote, exemplo: ``` $ cd fundamentos/funcoes ```
    1. rodar run go all: ``` $ go run *.go ```