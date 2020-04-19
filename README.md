# go-cookiecutter

A command-line utility that creates projects from templates for any language
The templating language is similar to [handlebars](https://handlebarsjs.com/) or [mustache](http://mustache.github.io/) templates. 


## Installation
- Download the latest binary for your OS release from the [releases page](https://github.com/arithran/go-cookiecutter/releases)
- Rename the file to `cookiecutter`
- Make the file executable (`chmod +x cookiecutter`)
- Check out the help menu for usage instructions `cookiecutter help`
- (Optional Step) Move it to a folder in your PATH variable. (`mv cookiecutter /bin`)


## Features

- [x] Basic find and replacement of variables
- [ ] Ignore parsing files in .gitignore
- [ ] Find & replace directory names that are templates

## Usage
```bash
cookiecutter run [directory] [find1 replace1 find2 replace2 ...]
```

**Example:**

This is a handler file in a project directory

```go
package handler

import (
	"fmt"
	"net/http"
)

var (
	SERVICE = "{{SERVICE}}"
	VERSION = "{{VERSION}}"
)

func Info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "service name is %s and version %s\n", SERVICE, VERSION)
}

```

Running the following command in the directory which holds this file will replace SERVICE & VERSION with login and v1 respectively
```bash
cookiecutter run . SERVICE login VERSION v1

```

Example templates can be found in the [examples folder](./examples/go-microservice-template)

## Contribute
Pull requests are welcome!
