var shell = require('shelljs')
shell.config.verbose = false

var process = require('process')
process.chdir('./src')

shell.exec('swag init')
shell.echo('http://localhost:{{devport}}/swagger/index.html')
