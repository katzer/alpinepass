# alpinepass

[![Travis CI Build Status](https://travis-ci.org/appPlant/alpinepass.svg?branch=master)](https://travis-ci.org/appPlant/alpinepass)
[![Code Climate](https://codeclimate.com/github/appPlant/alpinepass/badges/gpa.svg)](https://codeclimate.com/github/appPlant/alpinepass)
[![codebeat badge](https://codebeat.co/badges/35499bc6-3480-4e21-9e9e-29bb7d394bab)](https://codebeat.co/projects/github-com-appplant-alpinepass)
[![Go Report Card](https://goreportcard.com/badge/github.com/appPlant/alpinepass)](https://goreportcard.com/report/github.com/appPlant/alpinepass)

A tool for managing system environments.

```
$ alpinepass -h
NAME:
   alpinepass - Manage system environment information.

USAGE:
   alpinepass [global options] command [command options] [arguments...]

VERSION:
   1.4.3

AUTHOR:
   appPlant GmbH

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug                   Print the stack trace when an error occurs.
   --display, -d             Display the output in the console. An output file will not be written. Use this for previewing the generated file.
   --filter value, -f value  Filter configurations by name, type and more. Syntax: -f "[property]:[regex]"
   --input value, -i value   Specify the input file, absolute or relative path.
   --output value, -o value  Specify the output file, absolute or relative path. An existing file will be overwritten.
   --passwords, -p           Include passwords in the output.
   --readable, -r            Make the output more readable.
   --skip, -s                Skip the input validation.
   --help, -h                show help
   --version, -v             print the version
```

## Installation

Download the latest version from the release page and add the executable to your ```PATH```.

## Development setup

* install [Go](https://golang.org/)
* set the `$GOPATH`
* clone the repository into `{$GOPATH}/src/github.com/appPlant/alpinepass`
* `cd` into the repository
* run `go get -v`
* build the application with `go build`
* execute the application with `./alpinepass`

## Manual

### Overview

**alpinepass** is an application for managing information about computer system environments. It transforms information defined in a YAML file into different output formats like JSON and KeePass files. Filters allow creating output which matches certain creteria like information about databases or test systems only.

### Basic usage

Executing the application without any options reads the file ***input.yaml*** and creates the file ***output.json***. The input file must exist in the same directory as the application, the output file is created in the same directory. The output does not contain any system passwords. No filters are applied to the input.

```$ alpinepass```

Create a JSON output containing all information including passwords:

```$ alpinepass -p```

Create a JSON output containing production databases only:

```$ alpinepass -f=type:db;environment:prod```

Create a YAML output file ```/home/admin/endor_test.yml``` containing all test systems located on Endor:

```$ alpinepass -t=/home/admin -n=endor_test -o=yaml -f=location:Endor;environment:test```

## Dependencies

* [cli](https://github.com/urfave/cli)
* [testify](https://github.com/stretchr/testify)
* [yaml](https://github.com/go-yaml/yaml)
* [structs](https://github.com/fatih/structs)

## License

The code is available as open source under the terms of the [Apache 2.0 License](http://opensource.org/licenses/Apache-2.0).

Made with :heart: and :coffee: in Leipzig!

© 2017 [appPlant GmbH](http://www.appplant.de)
