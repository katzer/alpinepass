# alpinepass

[![Travis CI Build Status](https://travis-ci.org/appPlant/alpinepass.svg?branch=master)](https://travis-ci.org/appPlant/alpinepass)
[![Appveyor Build status](https://ci.appveyor.com/api/projects/status/8uavdsappvon3gjh?svg=true)](https://ci.appveyor.com/project/EightBitBoy/alpinepass)
[![Code Climate](https://codeclimate.com/github/appPlant/alpinepass/badges/gpa.svg)](https://codeclimate.com/github/appPlant/alpinepass)

A tool for managing system environments.

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

```alpinepass``` Executing the application without any options reads the file *input.yaml* and creates the file *output.json*. The input file must exist in the same directory as the application, the output file is created in the same directory. The output does not contain any system passwords. No filters are applied to the input.

### Options

* ```-i [filepath]``` input file
* ```-o [filepath]``` output file
* ```-p``` passwords
* ```-f [filter|filter:filteroptions]``` filters