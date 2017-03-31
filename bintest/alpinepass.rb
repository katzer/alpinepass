require 'open3'
require 'test/unit'

BIN  = ARGV.fetch(0).freeze
PATH = { 'PATH' => "#{File.expand_path('tools', __dir__)}:#{ENV['PATH']}"  }

class TestAlpinepass < Test::Unit::TestCase
	def test_run_without_arguments_and_input
		# output, error, status = Open3.capture3(PATH, BIN, '-r', '-s', '-i foobar')

    # puts PATH
    # puts BIN

		# expect_error(output, error, 'run without arguments and input')

		# assert_true status.success?, 'Process did not exit cleanly!'
    # assert_include output, 'The file input.yml does not exist!', 'The error message is not correct!'
	end

  def test_run_input_valid
		output, error, status = Open3.capture3(PATH, BIN, '-r', '-s', '-i valid.yml')

		expect_no_error(output, error, 'run valid.yml')
		assert_true status.success?, 'Process did not exit cleanly!'
	end
end

#there should be no error
def expect_no_error(output, error, test_name)
  return if error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error}"
end

#there should be an error
def expect_error(output, error, test_name)
  return unless error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error.inspect}"
end
