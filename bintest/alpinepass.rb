require 'open3'
require 'test/unit'

BIN  = ARGV.fetch(0).freeze
PATH = { 'PATH' => "#{File.expand_path('tools', __dir__)}:#{ENV['PATH']}"  }
LOCATION = BIN.gsub('alpinepass', '') #The path used for input and output files.

class TestAlpinepass < Test::Unit::TestCase
  def test_show_the_version
		output, error, status = Open3.capture3(PATH, BIN, '-v')
		assert_true status.success?, 'Process did not exit cleanly!'
		expect_no_error(output, error, 'Show the version.')
    assert_include output, 'alpinepass version'
	end

  #def test_run_input_invalid
	#	output, error, status = Open3.capture3(PATH, BIN, "-i=#{LOCATION}invalid.yml")
	#	assert_false status.success?, 'Process should fail!'
	#	expect_error(output, error, 'Use invalid.yml. There should be a validation error.')
	#end

  def test_run_input_valid
		output, error, status = Open3.capture3(PATH, BIN, "-i=#{LOCATION}valid.yml")
		assert_true status.success?, 'Process did not exit cleanly!'
		expect_no_error(output, error, 'Use valid.yml.')
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
