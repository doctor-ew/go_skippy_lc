# GPT-Powered Translator with LangChain

This project demonstrates how to build a GPT-powered translator using LangChain, a library that simplifies the process of building applications with OpenAI's models. The application allows users to interact with the GPT model and request translations in various languages.

## Overview

The code in this repository is based on:

1. [The official LangChain Quickstart documentation](https://python.langchain.com/docs/get_started/quickstart#environment-setup).
2. [Build a GPT-powered translator with LangChain](https://levelup.gitconnected.com/build-a-gpt-powered-translator-with-langchain-3e6915914daf).
3. [LangChain Go Client](https://github.com/tmc/langchaingo).

## Prerequisites

- Go (Golang) installed on your machine.
- An OpenAI API key.
- The `.env` file in the root directory of this project should contain your OpenAI API key as `OPENAI_API_KEY=YOUR_API_KEY`.

## How to Run

1. Clone this repository:

```bash
git clone [repository_url]
cd [repository_directory]
```

2. Load the environment variables:

```bash
source .env
```

3. Run the Go application:

```bash
go run main.go
```

4. Follow the on-screen instructions to interact with the GPT model and request translations.

## Contributing

Feel free to fork this repository and submit pull requests for any enhancements or fixes. If you encounter any issues, please open an issue on GitHub.

## License

This project is open source and available under the [MIT License](LICENSE).
