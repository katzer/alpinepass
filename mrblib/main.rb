# Apache 2.0 License
#
# Copyright (c) 2018 Sebastian Katzer, appPlant GmbH
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

@parser = OptParser.new do |opts|
  opts.add :input,    :string
  opts.add :output,   :string
  opts.add :format,   :string, 'json'
  opts.add :pretty,   :bool, false
  opts.add :secrets,  :bool, false
  opts.add :validate, :bool, false
end

@parser.on! :help do
  <<-USAGE

#{AlpinePass::LOGO}

usage: alpinepass [options...] -i input_file -o output_file matchers...
Options:
-i, --input     Path to the input file
-o, --output    Path to the output file
-f, --format    Format of the output file
                Defaults to: json
-p, --pretty    Pretty print output
-s, --secrets   Export secrets like passwords
-v, --validate  Validate the content of the input file
-h, --help      This help text
-v, --version   Show version number
USAGE
end

@parser.on! :version do
  "alpinepass #{AlpinePass::VERSION} - #{OS.sysname} #{OS.bits(:binary)}-Bit (#{OS.machine})" # rubocop:disable LineLength
end

# Entry point of the tool.
#
# @param [ Array<String> ] args The ARGV array.
#
# @return [ Void ]
def __main__(args)
  AlpinePass::Task.new(parse(args[1..-1])).exec
end

# Parse the command-line arguments.
#
# @param [ Array<String> ] args The command-line arguments to parse.
#
# @return [ Hash<Symbol, Object> ]
def parse(args)
  opts = @parser.parse(args.empty? ? ['-h'] : args)

  tail = @parser.tail
  tail << 'id:.' if tail.empty?

  opts[:matchers] = tail.map! { |m| AlpinePass::Matcher.new(m) }

  opts
end
