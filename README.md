# Navi: AI-Powered Shell Guide

"Navi" means "guide" in Hindi, and that's exactly what this tool aims to be. Navi is a command-line tool that utilizes the Gemini API to leverage artificial intelligence in generating computer code based on your prompts. It simplifies your workflow by understanding your intent and providing the necessary commands to achieve your tasks.

https://github.com/nitintf/navi/assets/55453926/407c41cb-0eb7-4362-a7db-65b224feb5ea

## Usage

```shell
Navi - AI-powered Programming Guide

Usage:
	navi prompt-string
  navi [flags]
  navi [command]

Examples:
navi "List all files in the current directory."

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for navi
```

Example:

```shell
$ navi "list all files`

$ ls -a
```

## Installation

You can install Navi using the `go install` command:

```shell
go install github.com/nitintf/navi
```

After installation, don't forget to export your Gemini API key:
```shell
export GEMINI_API_KEY="..."
```
or put it into a `.env`:
```shell
echo "GEMINI_API_KEY=..." > .env
```

You can obtain your Gemini API key from the [Gemini API Management page](https://aistudio.google.com/app/apikey).

## Caution

While Navi uses AI to generate shell commands, it's important to understand that AI isn't perfect. Always review the generated commands before executing them, especially if you're working in a production environment or dealing with sensitive data. Navi is a tool designed to assist you, but it doesn't replace good judgment and understanding of shell commands.

## Contribution

Contributions to Navi are very welcome! If you have a feature request, bug report, or proposal for code refactoring, please feel free to open an issue or submit a pull request.

This thing only supports Gemini for the moment, but it could be easily updated to support other models using Replicate, Ollama, etc. [Pull requests welcome](https://github.com/nitintf/navi/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc)!


## License

MIT
