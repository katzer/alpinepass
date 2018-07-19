# Apache 2.0 License
#
# Copyright (c) 2016 Sebastian Katzer, appPlant GmbH
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

require 'open3'
require_relative '../mrblib/alpinepass/version'

BINARY = File.expand_path('../mruby/bin/alpinepass', __dir__).freeze

%w[-v --version].each do |flag|
  assert("version [#{flag}]") do
    output, status = Open3.capture2(BINARY, flag)

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, AlpinePass::VERSION
  end
end

%w[-h --help].each do |flag|
  assert("usage [#{flag}]") do
    output, status = Open3.capture2(BINARY, flag)

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage'
  end
end

%w[-i --input].each do |flag|
  assert("wrong path [#{flag}]") do
    _, output, status = Open3.capture3(BINARY, flag, 'bad/file/path.ext')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include output, 'IOError'
  end
end

assert('input file [-i]') do
  output, status = Open3.capture2(BINARY, '-i', 'bad/file/path.ext')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'IOError'
end

assert('no flag') do
  output, status = Open3.capture2(BINARY)

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, 'usage'
end

assert('unknown flag') do
  _, output, status = Open3.capture3(BINARY, '-unknown')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'unknown option'
end
