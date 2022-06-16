# go-gin-template

This is a template for go project which is running with the gin for HTTP Client service with a clean architecture

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.

## Using `go-gin-template` project

To use `go-gin-template` project, follow these steps:

```bash
# Navigate into the project
cd ./go-gin-template

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
help                           Display this help screen
```

## Folder Structure
This project design by using clean architecture and hexagonal architecture so folder of project will organize base on 
clean architecture below 

Ref: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

Here below is how folder map to layer and component in clean architecture

- domain -> Entity
- usecase -> Usecase 
- repository -> Repository
- api -> Handler
- driver -> remote call

## Commit rules

### Conventional Commits

- We are using Conventional Commits rule to add readable meaning to commit messages
- We are following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) rule

#### Format for commit

[type][optional scope]: [optional REFERENCE-1234] [description]

- ex. build(husky): [BO-000] add husky and commitlint

- List of commit type
  [
  'build',
  'ci',
  'chore',
  'docs',
  'feat',
  'fix',
  'perf',
  'refactor',
  'revert',
  'style',
  'test'
  ]

## Contributing to `go-gin-template`

To contribute to `go-gin-template`, follow these steps:'

1. Clone this repository.
2. Create a feature branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <branch_name>`
5. Create the pull request against `master`.
