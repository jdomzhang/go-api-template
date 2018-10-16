var shell = require('shelljs')
shell.config.verbose = true

shell.exec('yarn build')
shell.exec('yarn upload')
