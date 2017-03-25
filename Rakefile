lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require_relative ENV.fetch('BUILD_CONFIG', 'build_config.glibc-2.14.rb')
require 'fileutils'

APP_NAME = 'alpinepass'.freeze
APP_ROOT = Dir.pwd.freeze

def release_path
  "#{APP_ROOT}/releases"
end

def build_path
  "#{APP_ROOT}/build"
end

def src_path
  "#{APP_ROOT}/src"
end

def orbit_path
  "#{APP_ROOT}/bintest/orbit"
end

def version_path
  "#{src_path}/version"
end

puts ENV["PATH"]

def get_version
  if ENV["DOCKER"]
    puts "###--- DOCKER"
    return "d1.0"
  elsif ENV["TRAVIS"]
    puts "###--- TRAVIS"
    return "t1.0"
  else
    puts "###--- DEV"
    return `go run #{version_path}/*.go`
  end
  #return `cp -r ./src $AP_HOME/src && go run #{version_path}/*.go -v`
end

APP_VERSION = get_version().freeze

Dir.chdir('lib') { Dir['tasks/*.rake'].each { |file| load file } }

desc 'print version'
task :version do
  puts APP_VERSION
end

desc 'removes artifacts and folders not becessary for a distribution.'
task :clean do
  rm_rf "#{orbit_path}/bin"
  rm_rf build_path
end
