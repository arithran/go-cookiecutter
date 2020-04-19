# go-cookiecutter

A command-line utility that creates projects from templates

## Features

- [x] Basic find and replacement of variables
- [ ] Ignore parsing files in .gitignore

## Usage
- cookiecutter run [directory] [find1 replace1 find2 replace2...]

```bash
cookiecutter run ./examples/go-microservice-template SERVICE login VERSION v1
```

Example templates, can be found in the [examples folder](./examples/go-microservice-template)