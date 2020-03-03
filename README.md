# Dir Cleaner example

This repo contains a CLI tool to scan and remove `node_modules` from your system.

It was used as a demo on our internal Golang program. It contains different versions with improvements.


## Usage

	NAME:
	   dir-cleaner - remove some unused files on your system

	USAGE:
	   dir-cleaner [--path <path>] [--depth <num>] [--dry-run]

	VERSION:
	   dev

	AUTHOR:
	   Gustavo Marin <gustavo.marin@intelygenz.com>

	COMMANDS:
	   help, h  Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --path value, -p value       path where to start the search (default: "$PWD")
	   --dry-run                    just check without deleting data (default: false)
	   --max-depth value, -d value  how many levels to check (use 0 for no max depth) (default: 0)
	   --bytes                      count bytes instead of default blocks of 4K to match 'du' reports (default: false)
	   --verbose                    print more info into the console (default: false)
	   --help, -h                   show help (default: false)
	   --version, -v                print the version (default: false)



## Versions

### Version 0

Just the CLI skeleton to see how to use `urfave/cli` and capture parameters. 


### Version 1

First implementation, no code separation, just a simple file.


### Version 2

First implementation, no code separation, just a simple file.


### Version 3

This was a mistake... never do a version 3. (note: one extra copy/paste error)


### Version 4

This version is the most complete version. It is full documented, with examples and includes parameters for the tool [GoReleaser](https://goreleaser.com/).


## TODO

- [ ] Make the search pattern a flag (currently only search for `node_modules`)
- [ ] Automate release to Github with Actions (see `version3/.github`)
- [ ] Make it auto-installable (see `version3/install.sh`)
