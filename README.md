# Consulta CEP API  

### 🇧🇷 Projeto simples em Go para consultar CEP usando múltiplas APIs  

### 🇺🇸 Simple Go project to query Brazilian ZIP codes using multiple APIs

# \---
 
## 🇧🇷 Sobre o projeto

# Este é um projeto básico desenvolvido para fins de estudo, criado por mim (Gisele).
O objetivo principal é aprender Go na prática, entendendo:
- Como trabalhar com múltiplos providers (APIs diferentes)
- Como implementar interfaces
- Como lidar com requisições HTTP
- Como padronizar respostas vindas de APIs distintas
- Como estruturar um projeto Go de forma organizada

# O programa consulta um CEP usando \*\*três APIs diferentes\*\*:

1. \*\*ViaCEP\*\*
2. \*\*BrasilAPI\*\*
3. \*\*WideNet (API CEP)\*\*

Ele tenta cada provider na ordem definida.  
Se um falhar, ele automaticamente tenta o próximo — até encontrar um resultado válido.

# \---

## 🇺🇸 About this project
# This is a basic project created for learning purposes, developed by me (Gisele).
The main goal is to learn Go by building something practical and understanding:
- How to work with multiple providers (different APIs)
- How to implement interfaces
- How to make HTTP requests
- How to normalize responses from different APIs
- How to structure a Go project in an organized way


# The program queries a Brazilian ZIP code (CEP) using \*\*three different APIs\*\*:

1. \*\*ViaCEP\*\*
2. \*\*BrasilAPI\*\*
3. \*\*WideNet (API CEP)\*\*

# It tries each provider in the defined order.  
# If one fails, it automatically falls back to the next until a valid result is found.

# \---

## 📁 Estrutura do projeto / Project structure

```
consulta-cep-api-go/
│
├── cmd/
│   └── main.go          # Ponto de entrada da aplicação / Application entry point
│
├── internal/
│   ├── service.go       # Lógica de fallback entre providers / Provider fallback logic
│   └── providers/
│       ├── provider.go  # Interface Provider e modelo CEPResponse
│       ├── via\_cep.go
│       ├── brasil\_api.go
│       └── wide\_net.go
│
└── go.mod
```

## ▶️ Como executar / How to run

### 🇧🇷 Em português
1. Abra o terminal na raiz do projeto
2. Execute:
```
go run ./cmd
Digite o CEP:01001000
CEP encontrado: &{CEP:01001-000 Logradouro:Praça da Sé Bairro:Sé Cidade:São Paulo Estado:SP}
```

### 🇺🇸 In English
1. Open your terminal at the project root  
2. Run:
```
go run ./cmd
Digite o CEP:01001000
CEP encontrado: &{CEP:01001-000 Logradouro:Praça da Sé Bairro:Sé Cidade:São Paulo Estado:SP}
```
# \---

## 🧠 O que eu aprendi / What I learned

### 🇧🇷 Português
- Como criar e usar interfaces em Go  
- Como fazer requisições HTTP  
- Como decodificar JSON com `encoding/json`  
- Como organizar um projeto usando `internal/`  
- Como implementar fallback entre múltiplas APIs  
- Como trabalhar com ponteiros e structs  

# \---

### 🇺🇸 English version
- How to create and use interfaces in Go  
- How to make HTTP requests  
- How to decode JSON using `encoding/json`  
- How to structure a project using `internal/`  
- How to implement fallback between multiple APIs  
- How to work with pointers and structs  

# \---

## 📌 Observações / Notes

### 🇧🇷 Português
- Este projeto é apenas para estudo e aprendizado.  
- As APIs usadas são públicas e gratuitas.  
- O código pode ser expandido facilmente para virar uma API HTTP real.

### 🇺🇸 English
- This project is for study and learning purposes only.  
- The APIs used are public and free.  
- The code can be easily expanded into a real HTTP API.

# \---

## 📜 Licença / License

### 🇧🇷 Português
Este projeto é apenas educacional e não possui licença formal.  
Sinta-se à vontade para modificar, estudar e evoluir.

### 🇺🇸 English
This project is educational and does not have a formal license.  
Feel free to modify, study, and improve it.
