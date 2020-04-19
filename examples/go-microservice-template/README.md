## This is the micro service skeleton of company XYZ.

### Getting started

copy directory to the new directory for the repo (ideally the service name)
1) Copy the directory to a new directory
`cp -R go-microservice-template login-service`

2) cd into the new directory and run `cookiecutter` command to instantiate your service
`cookiecutter run . SERVICE login VERSION v1`

3) Confirm that this is now login service by calling the info endpoint
```bash
go run main.go
curl localhost:8080/info
```
