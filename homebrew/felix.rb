class Felix < Formula
  homepage "https://github.com/felix-cli/felix"

  # How to update:
  # Step 1) increment the version number below:
  version "0.0.1"

  # resource "felix" do
  #   url "https://github.com/felix-cli/felix/archive/refs/tags/v0.1.3.tar.gz"
  #   sha256 "dcb5796bb109dea418bdde04abf80faf681a39c0"
  # end

  # Step 2) pin the SHA you just built
  stable do
    url "https://github.com/felix-cli/felix.git", revision: "dcb5796bb109dea418bdde04abf80faf681a39c0"
  end

  # Step 3) brew uninstall felix; brew install --build-bottle ./felix.rb

  # See steps below for how to update the bottle block
  # bottle do
  #   root_url ""
  #   sha256 catalina: ""
  #   sha256 big_sur:  ""
  # end

  depends_on "go" => :build

  def install
    resource("felix").stage { }
    system "go", "build", "-ldflags", "-X main.Version=#{version}", "-o", bin/"felix", "cmd/felix/main.go"
  end
end
