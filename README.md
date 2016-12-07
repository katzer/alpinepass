# alpinepass

[![Travis CI Build Status](https://travis-ci.org/appPlant/alpinepass.svg?branch=master)](https://travis-ci.org/appPlant/alpinepass)
[![Appveyor Build status](https://ci.appveyor.com/api/projects/status/8uavdsappvon3gjh?svg=true)](https://ci.appveyor.com/project/EightBitBoy/alpinepass)
[![Code Climate](https://codeclimate.com/github/appPlant/alpinepass/badges/gpa.svg)](https://codeclimate.com/github/appPlant/alpinepass)

A tool for managing system environments.

```
$ alpinepass -h

usage: alpinepass [options...]

Options:

-d Show debug messages during application execution.

-f=[filter|filter:filteroptions] Apply one or more filters. Filters are given as a list seperated by semicolons. A filter's options are separated by commas.

-h Display this help message.

-i=[filepath] Define the input file. The file content must be YAML.

-n=[filename] Specify the output's file name. This does not influence the file extension since it depends on the output format which is defined with the -t option. The default file name is output.

-o=[json|yaml|keepass] Specify the output format.

-p Include passwords in the output.

-t=[directorypath] Define the output target directory. The default directory is the application directory.

-v Print the application's version.
```

## Development setup

* install [Go](https://golang.org/)
* set your `$GOPATH`
* clone the repository into `{$GOPATH}/src/github.com/appPlant/alpinepass`
* `cd` into the repository
* run `go get -v`
* build the application with `go build`
* execute the application with `./alpinepass`

## Manual

### Overview

**alpinepass** is an application for managing information about computer system environments. It transforms information defined in a YAML file into different output formats like JSON and KeePass files. Filters allow creating output which matches certain creteria like information about databases or test systems only.

### Basic usage

```alpinepass``` Executing the application without any options reads the file ***input.yaml*** and creates the file ***output.json***. The input file must exist in the same directory as the application, the output file is created in the same directory. The output does not contain any system passwords. No filters are applied to the input.

### Options

### Input format

[TODO] Structure of input.yaml
