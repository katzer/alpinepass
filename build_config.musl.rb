require 'rubygems'
require 'os'
require 'go/build'

Go::Build.new('x86_64-pc-linux-busybox') do
  os :linux
  arch :amd64
  appname :ski
  bintest_if OS.linux? && OS.bits == 64
end
