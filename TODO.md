# TODO

## Development
* Implement "-f foo:bar" filter logic.
* Use some kind of enum for config types.
* Show improved error message when input file is a directory! Currently the error is about a file which cannot be read.
* Adjust validation rules.
* Update changelog.
* Make sure "-v" output matches the other application's output!
* Use Go 1.8! Make sure to update docker images and stuff!
* Improve validation code.

## Links
* https://github.com/golang/go/wiki/GithubCodeLayout
* https://github.com/tcnksm-sample/travis-golang/blob/master/.travis.yml
* [awesome-go](https://github.com/avelino/awesome-go)
* [waffle.io](https://waffle.io/)

## Libraries
* https://github.com/asaskevich/govalidator
* https://github.com/go-playground/validator
* https://github.com/avelino/awesome-go
* [spew](https://github.com/davecgh/go-spew)
* [logrus](github.com/Sirupsen/logrus)

## Tools
* [gox](github.com/mitchellh/gox)

## Other
* https://github.com/mattn
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
* Plugin as shared library?
* C-plugin

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
