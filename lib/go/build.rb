module Go
  class Build
    attr_reader :name

    def self.builds
      @builds ||= []
    end

    def initialize(name, &block)
      @name = name
      instance_exec(&block)
      self.class.builds << self
    end

    def os(name = nil)
      @os = name if name
      @os
    end

    def arch(name = nil)
      @arch = name if name
      @arch
    end

    def appname(name = nil)
      @appname = name if name
      @appname
    end

    def bintest_if(enabled)
      @test = enabled
      @test
    end

    def bintest?
      @test == true
    end

    def go_build(binpath)
      if OS.windows?
        "set GOOS=#{os}&&set GOARCH=#{arch}&&go get -v &&go build -i -ldflags=\"-s -X main.version=$(go run #{version_path}/*.go)\" -o #{binpath}/#{appname}"
      else
        "GOOS=#{os} GOARCH=#{arch} go get -v && go build -i -ldflags=\"-s -X main.version=$(go run #{version_path}/*.go)\" -o #{binpath}/#{appname};"
      end
    end
  end
end
