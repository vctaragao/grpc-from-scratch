# Projeto para estudar GRPC com a linguagem GO

## Instalação

Seguir Quickstart do site de documentação do grpc: [prerequisitos](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

## Setup
1. Após a instalação dos prérequisitos clonar o projeto

```bash
git clone git@github.com:vctaragao/grpc-from-scratch.git
```

2. Entrar na pasta

```bash
cd grpc-from-scratch
```

## Rodar servidor e client

1. Rodar Servidor

```bash
go run server/main.go
```

Deve se obter essa resposta

![image](https://user-images.githubusercontent.com/26884793/193913839-3cf3a472-6ace-4cca-b988-334eb5b9d1b2.png)

2. Rodar Cliente

```bash
go run client/main.go
```
Deve se obter essa resposta

![image](https://user-images.githubusercontent.com/26884793/193913920-e3c343d2-b82f-4624-8b4b-7c21540b6f6d.png)

Ao voltar para o terminal do servidor, se terá essa resposta

![image](https://user-images.githubusercontent.com/26884793/193914172-eb2220de-ef5d-4de1-97ca-306aff02c363.png)

### Implementação

- Client em GO
- Server em GO
