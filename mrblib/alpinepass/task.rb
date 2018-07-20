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
  # Class is responsible for reading the input file and delegate the data to the
  # exporter class.
  class Task
    # Initialize the task for given spec.
    #
    # @param [ Hash<Symbol,String|Array>] spec The parsed command-line args.
    #
    # @return [ Void ]
    def initialize(spec)
      @spec = spec
      @spec.freeze
    end

    # Execcute the task.
    #
    # @return [ Void ]
    def exec
      write convert select parse
    end

    # Read the input file and parse JSON into planets.
    #
    # @return [ Array<Hash> ]
    def parse
      planets = JSON.parse(File.open(@spec[:input], &:read))

      planets.each   { |p| p.delete 'password' } unless @spec[:secrets]
      planets.select { |p| @spec[:matchers].any? { |m| m.match? p } }
    end

    # Validate the planets if specified.
    #
    # @param [ Array<Hash> ] planets The planets to validate.
    #
    # @return [ Array<Hash> ] Subset of valid planets.
    def select(planets)
      planets = planets.select { |p| Planet.valid? p } if @spec[:check]
      planets
    end

    # Convert the data into a string depending on the output format.
    #
    # @param [ Array<Hash> ] planets The planets to convert.
    #
    # @return [ String ]
    def convert(planets)
      case @spec[:format]
      when 'json' then JSONConverter.new(@spec[:pretty]).convert(planets)
      else raise "unsupported format: #{@spec[:format]}"
      end
    end

    # Write the content to the output file.
    #
    # @param [ String ] content The content to write.
    #
    # @return [ Void ]
    def write(content)
      if (path = @spec[:output])
        File.open(path, 'w+') { |f| f.write(content) }
      else
        puts content
      end
    end
  end
end
