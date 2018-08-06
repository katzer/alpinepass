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

require 'json'
require 'open3'
require 'tempfile'
require_relative '../mrblib/alpinepass/version'

BINARY  = File.expand_path('../mruby/bin/alpinepass', __dir__).freeze
VALID   = File.expand_path('../test/fixtures/valid.export', __dir__).freeze
INVALID = File.expand_path('../test/fixtures/invalid.export', __dir__).freeze

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
    assert_include output, 'Usage'
  end
end

%w[-i --input].each do |flag|
  assert("wrong path [#{flag}]") do
    _, output, status = Open3.capture3(BINARY, flag, 'bad/file/path.ext')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include output, 'IOError'
  end
end

%w[-i --input].each do |flag|
  assert("bad format [#{flag}]") do
    _, output, status = Open3.capture3(BINARY, flag, __FILE__)

    assert_false status.success?, 'Process did exit cleanly'
    assert_include output, 'ParserError'
  end
end

[['-i'], ['-p', '-i'], ['-s', '-i']].each do |flags|
  assert("input file [#{flags.inspect}]") do
    output, status = Open3.capture2(BINARY, *flags, VALID)

    assert_true status.success?, 'Process did not exit cleanly'

    assert_nothing_raised do
      planets = JSON.parse(output)

      assert_kind_of Array, planets
      assert_equal 3, planets.count

      planets.each do |planet|
        assert_kind_of Hash, planet
        assert_false planet.empty?
        assert_equal flags.include?('-s'), planet.include?('password')
      end
    end
  end
end

assert('matcher [type=db]') do
  output, status = Open3.capture2(BINARY, '-i', VALID, 'type=db')

  assert_true status.success?, 'Process did not exit cleanly'

  assert_nothing_raised do
    planets = JSON.parse(output)

    assert_kind_of Array, planets
    assert_equal 1, planets.count
    assert_equal 'db', planets.first['type']
  end
end

%w[-f --format].each do |flag|
  assert("unknown format [#{flag}]") do
    _, output, status = Open3.capture3(BINARY, '-i', VALID, flag, ':x')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include output, 'unsupported format'
  end
end

%w[-o --output].each do |flag|
  assert("wrong path [#{flag}]") do
    _, output, status = Open3.capture3(BINARY, '-i', VALID, flag, 'b/a.d')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include output, 'IOError'
  end
end

%w[-c --check].each do |flag|
  assert("invalid content [#{flag}]") do
    file           = Tempfile.new('orbit.json')
    output, status = Open3.capture2 BINARY, flag, '-i', INVALID, '-o', file.path

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'missing values for name, user, url'
    assert_include output, 'missing values for type'
    assert_include output, 'found duplicate for id:duplicated-planet'
    assert_include output, 'found duplicate for name:Duplicate'
    assert_include output, 'unknown type'

    assert_nothing_raised do
      assert_equal 1, JSON.parse(file.read).count
    end
  ensure
    file.close
  end
end

[['-i'], ['-p', '-i'], ['-s', '-i']].each do |flags|
  assert("output file #{flags}") do
    file           = Tempfile.new('orbit.json')
    output, status = Open3.capture2 BINARY, *flags, VALID, '-o', file.path

    assert_true status.success?, 'Process did not exit cleanly'
    assert_raise(JSON::ParserError) { JSON.parse(output) }

    assert_nothing_raised do
      planets = JSON.parse(file.read)

      assert_kind_of Array, planets
      assert_equal 3, planets.count

      planets.each do |planet|
        assert_kind_of Hash, planet
        assert_false planet.empty?
        assert_equal flags.include?('-s'), planet.include?('password')
      end
    end
  ensure
    file.close
  end
end

assert('no flag') do
  output, status = Open3.capture2(BINARY)

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, 'Usage'
end

assert('unknown flag') do
  _, output, status = Open3.capture3(BINARY, '-unknown')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'unknown option'
end
