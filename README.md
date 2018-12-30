## Development
[![Build Status](https://travis-ci.com/blucard/blucard.svg?branch=master)](https://travis-ci.com/blucard/blucard)

Make sure [Go](https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/) is installed and setup correctly

In your Go workspace (`$GOPATH/github.com/blucard/`):
```bash
$ git clone git@github.com:blucard/blucard.git
```

### Initial Setup
```bash
$ make setup      # download dependencies
```

#### LocalStack Setup
```bash
$ make localstack # run local AWS stack (run this command in a different window/tab)
$ make dynamo     # creates local dynamo tables (only need to run if the table is not already created)
```

#### Start Server
```bash
$ make localstack # run in separate window/tab
$ make dev        # start local server
```
