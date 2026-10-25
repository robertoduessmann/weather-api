🌦️ Weather API

<p align="center">
  <strong>A REST API for retrieving weather information by city.</strong>
</p><p align="center">
  API REST desenvolvida em Go para consultar informações meteorológicas, incluindo temperatura, velocidade do vento, descrição das condições climáticas e previsão do tempo.
</p><p align="center">
  <a href="https://github.com/robertoduessmann/weather-api">
    <img src="https://img.shields.io/badge/Go-API-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go API" />
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge" alt="MIT License" />
  </a>
</p><p align="center">
  <a href="https://goreportcard.com/report/github.com/robertoduessmann/weather-api">
    <img src="https://goreportcard.com/badge/github.com/robertoduessmann/weather-api" alt="Go Report Card" />
  </a>
  <a href="https://godoc.org/github.com/robertoduessmann/weather-api">
    <img src="https://godoc.org/github.com/robertoduessmann/weather-api?status.svg" alt="GoDoc" />
  </a>
</p><p align="center">
  <a href="#-visão-geral">Visão geral</a> •
  <a href="#-endpoints">Endpoints</a> •
  <a href="#-instalação-e-execução">Instalação</a> •
  <a href="#-utilização">Utilização</a> •
  <a href="#-contribuição">Contribuição</a>
</p>---

📑 Sumário

- "Visão geral" (#-visão-geral)
- "Funcionalidades" (#-funcionalidades)
- "Tecnologias" (#-tecnologias)
- "Endpoints" (#-endpoints)
- "Pré-requisitos" (#-pré-requisitos)
- "Instalação e execução" (#-instalação-e-execução)
- "Utilização" (#-utilização)
- "Exemplo de resposta" (#-exemplo-de-resposta)
- "Tratamento de erros" (#-tratamento-de-erros)
- "Clientes web" (#️-clientes-web)
- "Contribuição" (#-contribuição)
- "Licença" (#-licença)

🌍 Visão geral

O Weather API é um projeto desenvolvido em Go que disponibiliza endpoints HTTP para consultar informações meteorológicas por cidade.

A API apresenta informações como temperatura atual, velocidade do vento, descrição das condições climáticas e dados de previsão.

O projeto pode ser utilizado como referência para estudos de desenvolvimento de APIs REST, integração entre aplicações e consumo de serviços meteorológicos.

Repositório: "robertoduessmann/weather-api" (https://github.com/robertoduessmann/weather-api)

«Nota: a disponibilidade do serviço público, o formato atual das respostas e a compatibilidade dos clientes relacionados precisam ser confirmados antes de serem considerados garantidos.»

✨ Funcionalidades

- 🌡️ Consulta de temperatura por cidade.
- 💨 Informações sobre a velocidade do vento.
- 🌤️ Descrição das condições meteorológicas.
- 📅 Dados de previsão do tempo.
- 🔌 Endpoints HTTP para integração com clientes.
- 🖥️ Possibilidade de integração com aplicações web.

As funcionalidades descritas refletem a documentação apresentada neste README; o comportamento efetivo depende da implementação e do serviço meteorológico utilizado.

🧰 Tecnologias

- Go: linguagem de programação utilizada no projeto.
- HTTP/REST: interface para consultas meteorológicas.
- JSON: formato de dados apresentado nos exemplos.
- curl: ferramenta para testar requisições HTTP.

🔗 Endpoints

API pública

A documentação do projeto apresenta o seguinte endpoint público:

GET https://goweather.xyz/v2/weather/{city}

Exemplo para Berlim:

GET https://goweather.xyz/v2/weather/Berlin

Substitua "{city}" pelo nome da cidade desejada.

Exemplo para Curitiba:

curl "https://goweather.xyz/v2/weather/Curitiba"

Importante: o endpoint público é apresentado conforme a documentação original. Sua disponibilidade e seu contrato de resposta atuais não estão confirmados.

API local

Com a aplicação executando localmente, a documentação indica:

GET http://localhost:3000/weather/{city}

Exemplo:

curl "http://localhost:3000/weather/Curitiba"

A porta "3000" e a rota "/weather/{city}" devem corresponder à configuração real da aplicação.

🛠️ Pré-requisitos

Antes de executar o projeto, instale:

- "Go" (https://go.dev/dl/)
- "Git" (https://git-scm.com/downloads)
- "curl" ou outra ferramenta para realizar requisições HTTP.

Confira as instalações:

go version
git --version
curl --version

🚀 Instalação e execução

1. Clone o repositório

git clone https://github.com/robertoduessmann/weather-api.git

Entre no diretório:

cd weather-api

2. Compile o projeto

go build

Se o módulo estiver configurado para gerar um executável com o nome do repositório, o binário resultante poderá ser "weather-api". O nome exato depende da configuração do projeto.

3. Execute a aplicação

Linux ou macOS

./weather-api

Windows PowerShell

.\weather-api.exe

Se a aplicação utilizar outra configuração de compilação ou nome de executável, ajuste o comando conforme necessário.

4. Teste a API

Com a aplicação em execução, abra outro terminal e execute:

curl "http://localhost:3000/weather/Curitiba"

Se a requisição falhar, confira a porta configurada, os logs da aplicação e as instruções do projeto.

📡 Utilização

Consultar uma cidade

O parâmetro "city" representa a cidade que você deseja consultar.

Exemplo de requisição:

GET /weather/Curitiba

Requisição usando "curl":

curl -i "http://localhost:3000/weather/Curitiba"

O argumento "-i" permite visualizar os cabeçalhos HTTP juntamente com a resposta.

Consultar o serviço público

curl -i "https://goweather.xyz/v2/weather/Curitiba"

O resultado depende da disponibilidade do serviço externo e da cidade solicitada.

📦 Exemplo de resposta

Abaixo está um exemplo ilustrativo de resposta JSON:

{
  "temperature": "29 °C",
  "wind": "20 km/h",
  "description": "Partly cloudy",
  "forecast": [
    {
      "day": "1",
      "temperature": "27 °C",
      "wind": "12 km/h"
    },
    {
      "day": "2",
      "temperature": "22 °C",
      "wind": "8 km/h"
    }
  ]
}

«Os valores são fictícios e servem apenas para ilustrar o formato dos dados. Não representam condições meteorológicas em tempo real.»

Descrição dos campos

Campo| Tipo ilustrativo| Descrição
"temperature"| String| Temperatura atual.
"wind"| String| Velocidade do vento.
"description"| String| Descrição das condições meteorológicas.
"forecast"| Array| Lista de previsões.
"forecast[].day"| String| Identificador do dia previsto.
"forecast[].temperature"| String| Temperatura prevista.
"forecast[].wind"| String| Velocidade do vento prevista.

Os tipos e as unidades acima refletem o exemplo documentado. Consulte a implementação para confirmar o contrato real da API.

⚠️ Tratamento de erros

Ao integrar a API em uma aplicação, considere tratar os seguintes cenários:

- 400 Bad Request: requisição ou parâmetro inválido, caso esse status seja utilizado pelo serviço.
- 404 Not Found: cidade ou recurso não encontrado, caso previsto pela implementação.
- 429 Too Many Requests: limite de requisições excedido, se houver limitação de uso.
- 500 Internal Server Error: falha interna do serviço.
- 503 Service Unavailable: serviço temporariamente indisponível.
- Erro de conexão: servidor inacessível, timeout ou falha de rede.

Essa lista é uma recomendação para clientes HTTP, não uma declaração de que a API implementa todos esses status.

🖥️ Clientes web

A documentação original relaciona os seguintes projetos de interface que podem servir como referências para aplicações de clima.

A compatibilidade atual com o endpoint desta API não foi verificada.

1. React Tempo

- Aplicação: https://reacttempo.netlify.app/
- Repositório: https://github.com/GabrielCampos99/appTempo

2. Weather App

- Aplicação: https://emaniaditya.github.io/weather-app
- Repositório: https://github.com/emaniaditya/weather-app/

3. Weather React TSX

- Aplicação: https://weather-react-tsx.netlify.app/
- Repositório: https://github.com/keissiant/weather-api-react.tsx

Esses projetos são referências externas; não são necessariamente mantidos pelos responsáveis pelo Weather API.

🤝 Contribuição

Contribuições para o código e a documentação são bem-vindas.

Como contribuir

1. Faça um fork do repositório.
2. Clone seu fork.
3. Crie uma branch para sua alteração.
4. Implemente a melhoria ou correção.
5. Execute os testes disponíveis.
6. Atualize a documentação quando necessário.
7. Envie suas alterações e abra um pull request.

Exemplo de fluxo

Crie uma branch:

git checkout -b feature/minha-melhoria

Adicione e registre as alterações:

git add .
git commit -m "Improve weather API documentation"

Envie a branch:

git push origin feature/minha-melhoria

Abra um pull request e descreva o objetivo da alteração, as mudanças realizadas e os testes executados.

Boas práticas

- Mantenha as alterações focadas.
- Siga as convenções existentes no repositório.
- Evite incluir credenciais e informações sensíveis.
- Documente mudanças nos endpoints e nos formatos de resposta.
- Inclua testes relevantes para novas funcionalidades e correções.

📄 Licença

Este projeto é apresentado como distribuído sob a licença MIT.

Consulte o arquivo "LICENSE" (LICENSE) no repositório para verificar os termos completos.

🔗 Links úteis

- Código-fonte: https://github.com/robertoduessmann/weather-api
- Issues: https://github.com/robertoduessmann/weather-api/issues
- Releases: https://github.com/robertoduessmann/weather-api/releases
- Go: https://go.dev/
- Documentação oficial do Go: https://go.dev/doc/

---

<p align="center">
  Desenvolvido com Go para facilitar o acesso a informações meteorológicas.
</p>