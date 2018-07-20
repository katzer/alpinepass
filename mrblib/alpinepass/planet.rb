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
  # Provides access to the content of the ORBIT_FILE.
  class Planet < BasicObject
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

    # Test if the fields describe a valid planet.
    #
    # @param [ Hash ] planet The planet to validate
    #
    # @return [ Boolean ]
    def self.valid?(fields)
      error = new(fields).validate

      Kernel.puts error if error

      error.nil?
    end

    # A planet is a single entry found in the ORBIT_FILE.
    #
    # @param [ Hash ] fields Extracted fields from the file.
    #
    # @return [ Planet ]
    def initialize(fields)
      @fields = fields
    end

    # Test if the planet is valid.
    #
    # @return [ String ] true if valid otherwise the error msg
    def validate
      type = @fields['type']

      return "unknown type \"#{type}\"" unless REQUIRED.include? type

      unknown = REQUIRED[type].reject { |name| @fields[name] }

      return if unknown.empty?

      "missing values for #{unknown.join(', ')} in #{@fields}"
    end
  end
end
