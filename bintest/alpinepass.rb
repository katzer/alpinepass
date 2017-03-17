require 'open3'
require 'test/unit'

BIN  = ARGV.fetch(0).freeze
PATH = { 'PATH' => "#{File.expand_path('tools', __dir__)}:#{ENV['PATH']}"  }

class TestAlpinepass < Test::Unit::TestCase
	def test_run
		output, error, status = Open3.capture3(PATH, BIN)
		assert_true status.success?, 'Process did not exit cleanly'
		check_error(output, error, 'run')
	end
end

def check_error(output, error, test_name)
  return if error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error}"
end

def check_no_error(output, error, test_name)
  return unless error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error.inspect}"
end
