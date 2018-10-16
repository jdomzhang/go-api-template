var shell = require('shelljs')
shell.config.verbose = true

shell.exec('go build -o ./tmp/src ./src')
