pick
====
A minimal password manager for OS X and Linux.

```sh
$ go get -u github.com/bndw/pick
```

or

```sh
$ git clone https://github.com/bndw/pick && cd pick
$ make
$ make install
```

![demo](https://github.com/bndw/pick/raw/master/demo.gif)

## Usage
```
Usage:
  pick [command]

Available Commands:
  add         Add a credential
  backup      Backup safe
  cat         Cat a credential
  cp          Copy a credential to the clipboard
  export      Export decrypted credentials in JSON format
  ls          List all credentials
  rm          Remove a credential
  version     Print the version number of pick

Flags:
  -h, --help   help for pick

Use "pick [command] --help" for more information about a command.
```

## Similar software
* [pwd.sh: Unix shell, GPG-based password manager](https://github.com/drduh/pwd.sh)
* [Pass: the standard unix password manager](http://www.passwordstore.org/)
