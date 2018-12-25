## Development
Make sure [Go](https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/) is installed and setup correctly

In your Go workspace (`$GOPATH/github.com/blucard/`):
```bash
$ git clone git@github.com:blucard/blucard.git
```

### Intital Setup
```bash
$ make setup      # download dependencies
```

```bash
$ make localstack           # run local AWS stack (run this command in a different window/tab)
$ pip install awscli-local  # only need to run once
$ make dynamo               # creates local dynamo tables

$ make dev        # start local server
```
