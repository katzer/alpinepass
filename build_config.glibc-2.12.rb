require 'rubygems'
require 'os'
require 'go/build'

Go::Build.new('x86_64-pc-linux-gnu-glibc-2.12') do
  os :linux
  arch :amd64
  appname :ski
  bintest_if OS.linux? && OS.bits == 64
end

Go::Build.new('i686-pc-linux-gnu-glibc-2.12') do
  os :linux
  arch :'386'
  appname :ski
  bintest_if OS.linux? && OS.bits == 32
end
