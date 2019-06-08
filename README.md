# alpinepass [![GitHub release](https://img.shields.io/github/release/appplant/alpinepass.svg)](https://github.com/appplant/alpinepass/releases) [![Build Status](https://travis-ci.com/appPlant/alpinepass.svg?branch=master)](https://travis-ci.com/appPlant/alpinepass) [![Build status](https://ci.appveyor.com/api/projects/status/76ytl3ycyqs0va8j/branch/master?svg=true)](https://ci.appveyor.com/project/katzer/alpinepass/branch/master) [![Maintainability](https://api.codeclimate.com/v1/badges/f5d906346be0d06ae4fc/maintainability)](https://codeclimate.com/github/appPlant/alpinepass/maintainability)

The tool that exports your Orbit KeePass database into various file formats.

    $ alpinepass -h

    Usage: alpinepass [options...] -i input_file [-o output_file] matchers...
    Options:
    -i, --input     Path to the input file
    -o, --output    Path to the output file
    -f, --format    Format of the output file
                    Defaults to: fifa
    -c, --check     Check the content of the input file
    -p, --pretty    Pretty print output
    -s, --secrets   Export secrets like passwords
    -h, --help      This help text
    -v, --version   Show version number

## Installation

Download the latest version from the [release page][releases] and add the executable to your `PATH`.

## Usage

Transform the _KeePass_ file by using the `$ORBIT_HOME/config/orbit.xsl` stylesheet:

    KPScript -c:Export "%ORBIT_HOME%\config\orbit.kdbx" -pw:MyPw -Format:"Transform using XSL Stylesheet" -XslFile:"%ORBIT_HOME%\config\orbit.xsl" -OutFile:"%ORBIT_HOME%\config\orbit.export"

Then convert the exported data into a valid knowledge database for _fifa_:

    $ alpinepass -i keepass.export -f fifa -o orbit.json

To create a knowledge database for _fifa_ containing production databases only:

    $ alpinepass -i keepass.export -f fifa -o orbit.json type=db@env=prod

__Note:__ See [here][keepass] for how to use KPScript with single command operations to perform simple database operations. Of course you can also use the GUI to perform the export.

## Development

Clone the repo:

    $ git clone https://github.com/appplant/alpinepass.git && cd alpinepass/

And then execute:

    $ rake compile

To compile the sources locally for the host machine only:

    $ MRUBY_CLI_LOCAL=1 rake compile

You'll be able to find the binaries in the following directories:

- Linux (64-bit Musl): `mruby/build/x86_64-alpine-linux-musl/bin/alpinepass`
- Linux (64-bit GNU): `mruby/build/x86_64-pc-linux-gnu/bin/alpinepass`
- Linux (64-bit, for old distros): `mruby/build/x86_64-pc-linux-gnu-glibc-2.12/bin/alpinepass`
- OS X (64-bit): `mruby/build/x86_64-apple-darwin17/bin/alpinepass`
- Windows (64-bit): `mruby/build/x86_64-w64-mingw32/bin/alpinepass`
- Host: `mruby/build/host2/bin/alpinepass`

For the complete list of build tasks:

    $ rake -T

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/appplant/alpinepass.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

The code is available as open source under the terms of the [Apache 2.0 License][license].

Made with :yum: in Leipzig

© 2017 [appPlant GmbH][appplant]

[releases]: https://github.com/appplant/alpinepass/releases
[keepass]: https://keepass.info/help/v2_dev/scr_sc_index.html#export
[license]: http://opensource.org/licenses/Apache-2.0
[appplant]: www.appplant.de
