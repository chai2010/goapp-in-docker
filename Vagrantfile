# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD license
# that can be found in the LICENSE file.

# vagrant up
# vagrant reload
# vagrant status
# vagrant ssh

# vagrant halt
# vagrant destroy

# vagrant reload --provision
# vagrant reload --provision-with shell

Vagrant.configure(2) do |config|
    # https://github.com/tmatilai/vagrant-proxyconf
    # vagrant plugin install vagrant-proxyconf
    if Vagrant.has_plugin?("vagrant-proxyconf")
      #config.proxy.http     = "http://127.0.0.1:2081/"
      #config.proxy.https    = "http://127.0.0.1:2081/"
      #config.proxy.no_proxy = "localhost,127.0.0.1,.example.com"
    end
  
    config.vm.box = "centos/7"
    config.vm.box_check_update = false
  
    config.vm.synced_folder ".", "/vagrant"
    config.vm.network "forwarded_port", guest: 80, host: 2010
  
    # https://www.vagrantup.com/docs/provisioning/basic_usage.html
    config.vm.provision "shell", path: "bootstrap.sh"
    config.vm.provision "shell", inline: <<-SHELL
      docker pull hello-world
      docker run  hello-world
      echo "done"
    SHELL
  end
  