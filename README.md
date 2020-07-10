<h1 align="center">Quick Git Repo</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

Quick Git Repo (qgr) is a command line utility to aid in the creation of a git repository.
Currently the utility puts the [MIT License](https://choosealicense.com/licenses/mit/)
in `LICENSE`, a fancy readme in `README.md`, and with the flag `-i`, can initialize a
git repository in your current directory.

## Installation

### Scoop
Do you have [scoop](https://github.com/lukesampson/scoop) installed? 
Just use my scoop [bucket](https://github.com/GreatGodApollo/trough) to install QGR.
```bash
$ scoop bucket add trough https://github.com/GreatGodApollo/trough.git

$ scoop install qgr
```

### Prepackaged Binaries
Didn't find a method of installation that suits you above? Guess it's time for you to head
on over to the [releases](https://github.com/GreatGodApollo/qgr/releases) page, where you can download the right executable for your system.
Of course you may want to add it to your path for ease of use ;)


## Built With
- [spf13/cobra](https://github.com/spf13/cobra)
- [go-cmd/cmd](https://github.com/go-cmd/cmd)

## Usage

```bash
$ qgr --help
This command line utility gives you a simple place
to start with your git repository. It starts you out
with a fancy shmancy README.md and the MIT LICENSE.

Usage:
  qgr [flags]

Flags:
  -a, --author string           author's name
  -u, --authorUsername string   author's username
  -d, --description string      project's description
  -h, --help                    help for qgr
  -i, --init                    initialize git repo
  -n, --name string             project's name
  -v, --version                 version for qgr
```

## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)
