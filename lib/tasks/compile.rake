require 'os'

desc 'compile binary'
task :compile do
  Go::Build.builds.each do |gb|
    bin_path = "#{build_path}/#{gb.name}/bin"

    mkdir_p bin_path
    cd(src_path) { sh gb.go_build(bin_path) }
    chmod_R 'u+x', bin_path
  end
end
