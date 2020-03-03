/*

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

*/
package main
