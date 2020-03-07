[![Tests](https://img.shields.io/github/workflow/status/guumaster/dir-cleaner/Test)](https://github.com/guumaster/dir-cleaner/actions?query=workflow%3ATest)
[![GitHub Release](https://img.shields.io/github/release/guumaster/dir-cleaner.svg?logo=github&labelColor=262b30)](https://github.com/guumaster/dir-cleaner/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/guumaster/dir-cleaner)](https://goreportcard.com/report/github.com/guumaster/dir-cleaner)
[![License](https://img.shields.io/github/license/guumaster/dir-cleaner)](https://github.com/guumaster/dir-cleaner/LICENSE)
# Dir Cleaner example

A simple tool to scan and remove unwanted directories from your system. (mainly  `node_modules`)

It was used as a demo on our internal Golang program. It contains different versions with improvements.

## Installation

Go to [release page](https://github.com/guumaster/dir-cleaner/releases) and download the binary you need.


## Usage

	NAME:
	   dir-cleaner - remove some unused files on your system

	USAGE:
	   dir-cleaner [--path <path>] [--depth <num>] [--dry-run]

	VERSION:
	   1.0.0

	AUTHOR:
	   Guumaster <guuweb@gmail.com>

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


## TODO

- [ ] Make the search pattern a flag (currently only search for `node_modules`)
- [ ] When searching for `node_modules` match first occurrence and not inner folders.


### Dependencies & Refs
  * [dustin/go-humanize](https://github.com/dustin/go-humanize)
  * [urfave/cli](https://github.com/urfave/cli)


### LICENSE

 [MIT license](LICENSE)


### Author(s)

* [guumaster](https://github.com/guumaster)

