var shell = require('shelljs');
shell.config.verbose = true

shell.exec('git pull')
shell.exec('git submodule update --init --recursive')
