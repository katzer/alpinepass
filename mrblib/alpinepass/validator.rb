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

module AlpinePass
  # Ensures that all required fields are present and
  # all identifiers are uniq accross all passed planets.
  class Validator
    # Fields that have to be unique accross all planets
    UNIQUE   = %w[id name].freeze
    # Fields that have to be present per planet type
    REQUIRED = {
      'server'.freeze => %w[id name user url].freeze,
      'db'.freeze     => %w[id name pqdb].freeze,
      'web'.freeze    => %w[id name].freeze,
      'log'.freeze    => %w[id name].freeze,
      nil             => %w[type].freeze
    }.freeze

    # Initializes the validator.
    #
    # @param [ Boolean ] check Flag indicates if the validator should validate.
    #
    # @return [ Void ]
    def initialize(check)
      @check = check
    end

    # Test if the fields describe a valid planet.
    #
    # @param [ Hash ] planet The planet to validate
    #
    # @return [ Boolean ] true if valid
    def validate(planet)
      return true unless @check

      type   = planet['type']
      fields = REQUIRED[type]

      return puts("unknown type \"#{type}\"") && false unless fields

      unknown = fields.reject { |name| planet[name] }

      return true if unknown.empty?

      puts "missing values for #{unknown.join(', ')} in #{planet}"

      false
    end

    # Test if references are all valid accross the planets.
    #
    # @param [ Hash ]        planet The planets to validate.
    # @param [ Array<Hash> ] planets The planets to validate against.
    #
    # @return [ Boolean ] true if valid
    def cross_validate(planet, planets)
      return true unless @check

      UNIQUE.select do |name, val = planet[name]|
        if planets.find_all { |p| p[name] == val }.size > 1
          puts "found duplicate for #{name}:#{val}"
          true
        else
          false
        end
      end.empty?
    end
  end
end
