var shell = require('shelljs');
process.env.port  = {{prodport}}
process.env["GIN_MODE"] = "release"
shell.echo('=========================================================');
shell.echo('=======*starting prod*=========');
shell.echo('=========================================================');
shell.exec('fresh')
