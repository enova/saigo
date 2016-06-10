# Setting Up Your Go Environment


  1. Install [GVM](https://github.com/moovweb/gvm)

  ```bash
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
  ```

  Place this line in `~/.zshrc` or `~/.bashrc` to source the GVM directory
  ```bash
  [[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
  ```

  Reopen your terminal or run `source $HOME/.gvm/scripts/gvm`

  2. Install Go

  ```bash
  gvm install go1.4
  ```
  3. Tell GVM which Go version to use

  ```bash
  gvm use go1.4
  ```

  Read about the Go command from the Go documentation: https://golang.org/doc/articles/go_command.html

  4. Set your $GOPATH accordingly. `mygo` can be whatever name you wish to choose for your workspace

  ```bash
  $ mkdir $HOME/mygo
  $ export GOPATH=$HOME/mygo
  ```

  5. Run your first program

  ```bash
  $ mkdir -p $GOPATH/src/github.com/user
  $ mkdir $GOPATH/src/github.com/user/hello
  $ cd $GOPATH/src/github.com/user/hello
  ```

  Next, create a file named hello.go inside the hello directory, containing the following Go code.

  ```go
  package main

  import "fmt"

  func main() {
        fmt.Printf("Hello, world.\n")
  }
  ```

  Now you can build and install that program with the go tool:

  ```bash
  $ go install
  ```

  This command builds the hello command, producing an executable binary. It then installs that binary to the workspace's bin directory as hello

  ```bash
  $ $GOPATH/bin/hello
  Hello, world.
  ```

  Once you have added $GOPATH/bin to your PATH, just type the binary name:
  ```bash
  $ hello
  Hello, world.
  ```
