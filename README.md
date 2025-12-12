# 🈺 Tradutor IA — Projeto em Desenvolvimento

Bem-vindo ao repositório do **Tradutor IA**, um projeto criado para aprender, praticar e implementar tradução automática utilizando **inteligência artificial**. Este README documenta todas as etapas, funcionalidades e como rodar o projeto.

---

## 🚀 Sobre o Projeto

O objetivo deste projeto é construir um **tradutor inteligente** capaz de traduzir textos entre diferentes idiomas (como Português ↔ Japonês), utilizando APIs de IA e técnicas modernas de NLP.

Este projeto serve tanto como **experimento educacional** quanto base para aplicações mais avançadas.

---

## 📌 Funcionalidades (atual e planejadas)

### ✔️ Implementadas

* Tradução básica usando API
* Interface simples para envio de texto
* Retorno rápido com formato legível

### 🛠️ Em desenvolvimento

* Detecção automática de idioma
* Traduções contextuais (melhor qualidade)
* Fala → Texto / Texto → Fala
* Histórico de traduções
* Modo "Aprender Japonês"

---

## 🧠 Tecnologias Utilizadas

* **Golang para Backend**
* **API de Inteligência Artificial Gemini**

---

## ⚙️ Como Rodar o Projeto

### 1. Adquira sua chave API

```
https://aistudio.google.com/api-keys
```

### 2. Instale as dependências

**Go**

```
go install go.golang,org/genai@1.39.0
```

### 3. Configure a chave da API

```
setx GEMINI_API_KEY sua_key
```

### 4. Inicie o servidor

**Go**

```
cd backend
go run main.go
```

---

## 🧪 Exemplo de Uso da API

### Requisição:

```json
Digite o texto a ser traduzido: Olá, tudo bem?
```

### Resposta: (Depende de como você polir o prompt enviado ao Gemini)

```json
(No prompt presente no código):
A forma mais comum e educada de dizer "Olá, tudo bem?" em japonês é:

**こんにちは、お元気ですか？**

**Como se lê (Romaji):**
**Konnichiwa, o-genki desu ka?**
```

---

## 📜 Licença

Este projeto é de uso livre para estudo. Sinta-se à vontade para modificar e melhorar.

---

---

## 🤝 Contribuições

Contribuições são bem-vindas! Abra um **pull request** ou **issue**.

---

## 📧 Contato

**Caso queira trocar ideias ou tirar dúvidas: **

** Mateus Schmidtke — Desenvolvedor do projeto.**

**Gmail: [mateusschmidtke0@gmail.com](mailto:mateusschmidtke0@gmail.com)**
