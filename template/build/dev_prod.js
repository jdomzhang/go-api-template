var shell = require('shelljs');
process.env.port  = {{prodport}}
process.env["GIN_MODE"] = "release"
shell.echo('=======*starting prod*=========');
shell.exec('fresh')
