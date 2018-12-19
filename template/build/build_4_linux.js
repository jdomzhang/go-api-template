var shell = require('shelljs');
shell.config.verbose = true
process.env.port  = 6001
process.env['GOOS'] = 'linux'
shell.echo('=======* starting build 4 linux *=========');
shell.rm('-rf', './dist')
shell.mkdir('dist')
shell.cp('-rf', './static', './dist/')
shell.cp('-rf', './config', './dist/')
shell.exec('go build -o ./dist/{{name}} ./src')
shell.exec('rm -rf ./tmp')
shell.mkdir('tmp')
shell.exec('tar -cvf ./tmp/{{name}}.tar -C ./dist .')
shell.exec('gzip -v ./tmp/{{name}}.tar')
