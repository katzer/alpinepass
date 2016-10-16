
def __main__(argv)
  puts "argv:"
  puts argv
  if argv[1] == "version"
    puts "v#{Alpinepass::VERSION}"
  else
    puts "Hello World"
  end
end
