require 'rubygems'
require 'os'

desc 'generate a release tarball'
task release: [:compile] do
  release_dir = "#{release_path}/#{APP_VERSION}"
  latest_dir  = "#{release_path}/latest"

  mkdir_p(release_dir)
  rm_rf(latest_dir)
  mkdir_p(latest_dir)

  cd(release_dir) { sh "tar czf #{APP_NAME}-#{APP_VERSION}.tgz #{src_path}" }

  Go::Build.builds.each do |gb|
    bin_path = "#{build_path}/#{gb.name}/bin/"
    tar_path = "#{APP_NAME}-#{APP_VERSION}-#{gb.name}.tgz"
    cd(release_dir) { sh "tar czf #{tar_path} #{bin_path}" }
  end

  ln Dir.glob("#{release_dir}/*"), latest_dir
end
