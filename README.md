[![MainWorkflow](https://github.com/viallikavoo/hello-world/actions/workflows/main-workflow.yml/badge.svg?branch=main)](https://github.com/viallikavoo/hello-world/actions/workflows/main-workflow.yml)

### Running locally

#### Requirements

1. Cmake [link](https://cmake.org/install/)
2. docker
3. go
 
#### Commands

`make clean` will clean the binaries generated </br>
`make generate-dockerfile APP_NAME=hello_world` will print the generated dockerfile with a comment on top of the current timestamp, and generate a Dockerfile as well.</br>
`make test` will build and run the docker image and test if it outputs "Hello World".</br>

#### Github Workflow

`main-workflow` will be invoked on push to master, which will perform a 
1. `make generate-dockerfile APP_NAME=hello_world`
2. `make test`
3. Build, and push the image to `ghcr.io`. The image will be tagged using the branchname and shorthash

### Flow
1. The `template_generator.go` is the go program that will use the dockerfile template located at `https://raw.githubusercontent.com/viallikavoo/docker-templates/main/templates/go/Dockerfile.template` and generate a Dockerfile with the timestamp as a comment on the top.