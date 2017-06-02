# TODO

## Development
* Show improved error message when input file is a directory! Currently the error is about a file which cannot be read.
* Adjust validation rules.
* Change to Version 1.42.
* Update changelog.
* Make sure "-v" output matches the other application's output!

## Links
* https://github.com/golang/go/wiki/GithubCodeLayout
* https://github.com/tcnksm-sample/travis-golang/blob/master/.travis.yml
* [awesome-go](https://github.com/avelino/awesome-go)
* [waffle.io](https://waffle.io/)

## Libraries
* [spew](https://github.com/davecgh/go-spew)
* [logrus](github.com/Sirupsen/logrus)

## Tools
* [gox](github.com/mitchellh/gox)

## Other
* Plugin-System: https://awesome-go.com/#embeddable-scripting-languages
* Add a "debug" flag for showing stack traces!
* binary stripping?
* https://onsi.github.io/ginkgo/
* http://agouti.org/
* http://onsi.github.io/gomega/
* !!! https://github.com/stretchr/testify

## Plugin
* https://en.wikipedia.org/wiki/Foreign_function_interface
* https://appliedgo.net/plugins/

### External process, RPC
* https://npf.io/2015/05/pie/, Go
* https://github.com/hashicorp/go-plugin, Go
* Plugin does not crash app
* Any language
* Must be compiled, depending on language

### Compile-time plugins
* Must be compiled
* Must compile whole application

### Embedded scripting language
* https://github.com/yuin/gopher-lua, Lua
* https://github.com/Shopify/go-lua, Lua
* https://github.com/robertkrimen/otto, JS
* https://github.com/sbinet/go-python, Python
* In-process
* Performance overhead, negligible
