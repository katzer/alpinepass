namespace :test do
  desc 'run integration tests'
  task bintest: [:compile] do
    config_path = "#{orbit_path}/config"
    #ssh_path    = "#{orbit_path}/config/ssh"

    mkdir_p config_path
    #mkdir_p ssh_path
    #cp '/root/.ssh/orbit.key', ssh_path

    Go::Build.builds.each do |gb|
      next unless gb.bintest?
      test_bin_path = "#{orbit_path}/bin/alpinepass"
      bin_path      = "#{build_path}/#{gb.name}/bin"

      test_bin_path << '.exe' if OS.windows?

      cp_r bin_path, orbit_path
      sh "ruby #{APP_ROOT}/bintest/alpinepass.rb #{test_bin_path}"
    end
  end
end
