# go-cookiecutter

A command-line utility that creates projects from templates

## Installation
- Download the latest binary for your OS release from the [releases page](https://github.com/arithran/go-cookiecutter/releases)
- Rename the file to `cookiecutter`
- Make the file executable (`chmod +x cookiecutter`)
- Checkout the help menu for usage instructions `cookiecutter help`
- (Optional Step) Move it to a folder in your PATH variable. (`mv cookiecutter /bin`)


## Features

- [x] Basic find and replacement of variables
- [ ] Ignore parsing files in .gitignore

## Usage
- cookiecutter run [directory] [find1 replace1 find2 replace2...]

```bash
cookiecutter run ./examples/go-microservice-template SERVICE login VERSION v1
```

Example templates, can be found in the [examples folder](./examples/go-microservice-template)

## Contribute
Pull requests are welcome!