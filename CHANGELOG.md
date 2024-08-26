# Release Notes: _alpinepass_

The tool that exports your Orbit KeePass database into various file formats.

## 1.5.2

Released at:

1. Upgraded to mruby 3.2.0

2. Compiled binary for OSX build with MacOSX11.3 SDK

3. Added binary for `arm64-apple-darwin19` target

[Full Changelog](https://github.com/katzer/alpinepass/compare/1.5.1...master)

## 1.5.1

Released at: 18.03.2020

<details><summary>Releasenotes</summary>
<p>

1. Compiled binary for OSX build with MacOSX10.15 SDK

2. Upgraded to mruby 2.1.0

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/1.5.0...1.5.1)
</details>

## 1.5.0

Released at: 13.08.2019

<details><summary>Releasenotes</summary>
<p>

1. Compiled with `MRB_WITHOUT_FLOAT`

2. Compiled binary for OSX build with MacOSX10.13 SDK (Darwin17)

3. Upgraded to mruby 2.0.1

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/1.4.7...1.5.0)
</details>

## 1.4.7

Released at: 02.01.2019

<details><summary>Releasenotes</summary>
<p>

1. Print errors and warning to SDTOUT with colors.

2. Handle unknown types as generic instead of invalid.

3. Dropped compatibility with orbit v1.4.6 due to breaking changes in _fifa_.

4. Removed LVAR section for non test builds.

5. Upgraded to mruby 2.0.0

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/1.4.6...1.4.7)
</details>

## 1.4.6

Released at: 16.08.2018

<details><summary>Releasenotes</summary>
<p>

Tool has been fully reworked.

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

Transform the _KeePass_ file by using the `$ORBIT_HOME/config/orbit.xsl` stylesheet:

    KPScript -c:Export "%ORBIT_HOME%\config\orbit.kdbx" -pw:MyPw -Format:"Transform using XSL Stylesheet" -XslFile:"%ORBIT_HOME%\config\orbit.xsl" -OutFile:"%ORBIT_HOME%\config\orbit.export"

Then convert the exported data into a valid knowledge database for _fifa_:

    $ alpinepass -i keepass.export -f fifa -o orbit.json

To create a knowledge database for _fifa_ containing production databases only:

    $ alpinepass -i keepass.export -f fifa -o orbit.json type=db@env=prod

See [here](https://keepass.info/help/v2_dev/scr_sc_index.html#export) for how to use KPScript with single command operations to perform simple database operations. Of course you can also use the GUI to perform the export.

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/1.4.3...1.4.6)
</details>

## 1.4.3

Released at: 06.10.2017

<details><summary>Releasenotes</summary>
<p>

1. Support multiple users for each system configuration.

2. Improve error messages, do not show the help text when an error occurs.

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/567d876072c20ade96c2de50de5c58707658311e...0224af8cbe1869f2bf9151e62ea75f827570fc04)
</details>

## 1.4.2

Released at: 29.09.2017

<details><summary>Releasenotes</summary>
<p>

1. Adjust the release script.

2. Introduce the "tool" configuration type.

3. Improved error messages. When an error occurs an error message is shown which indicates the error reason followed by the help text.

4. Add the "--debug" flag which prints the stacktrace when an error occurs. No help text is shown.

5. New filter logic:
    * "Exact" filters "key=value" match when the key's content matches exactly the filter value. "Earth" matches "Earth" but not "Earths".
    * "Contains" filters "key:value" match then the filter value is contained anywhere in the key's content. "Bar" matches "FooBar" and "BarFoo" but not "BazFoo".

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/0224af8cbe1869f2bf9151e62ea75f827570fc04...3051f84b40aecbb59355c22fd4dd7a79c6833c51)
</details>

## 1.0.0

Released at: 20.03.2017

<details><summary>Releasenotes</summary>
<p>

1. Basic functionality. Read "input.yml" and write "output.json".
    ```
    $ ls
    input.yml
    $ alpinepass
    $ ls
    input.yml output.json
    ```

2. New "input" flag for specifying the input file.
    ```
    $ alpinepass -i /path/to/input.yml
    ```

3. New "output" flag for specifying the output file.
    ```
    $ alpinepass -o /path/to/output.json
    ```

4. New "display" flag for previewing the output in the console. An output file will not be written.
    ```
    $ alpinepass -d
    [{"id": "B01.prod.server.PROD-App","title": "PROD App","location": "B01","environment": "prod","user": "prodUserB01"}]
    ```

5. New "readable" flag for formatting the output. It works with both output file and console output!
    ```
    $ alpinepass -r
    $ cat output.yml
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01"
        }
    ]
    $ alpinepass -d -r
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01"
        }
    ]
    ```

6. New "passwords" flag for including passwords in the output.
    ```
    $ alpinepass -d -p
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
            "password": "prod_pw"
        }
    ]
    ```

7. New "filter" flag for filtering the input.
    ```
    $ alpinepass -d -p
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
            "password": "prod_pw"
        },
        {
            "id": "D02.prod.server.PROD-App",
            "title": "PROD App",
            "location": "D02",
            "environment": "prod",
            "user": "prodUserD02",
            "password": "prod_pw"
        }
    ]
    $ alpinepass -d -p -f location:D02
    [
        {
            "id": "D02.prod.server.PROD-App",
            "title": "PROD App",
            "location": "D02",
            "environment": "prod",
            "user": "prodUserD02",
            "password": "prod_pw"
        }
    ]
    ```

8. Input verification checks that certain properties are present for the different configuration types.

9. New "skip" flag for disabling input verification.
    ```
    $ alpinepass -d -r -s
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD DB",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
            "host": "prodDatabase.B01"
        }
    ]
    $ alpinepass -d -r
    The host "prodDatabase.B01" does not match the naming convention "[environment]Db.[location]".
    ```

</p>

[Full Changelog](https://github.com/katzer/alpinepass/compare/ebf72cfa34e40dc418ec45b42c9a902a1abe5a20...3051f84b40aecbb59355c22fd4dd7a79c6833c51)
</details>