var shell = require('shelljs');
process.env["http_proxy"]  = "http://127.0.0.1:7777"
shell.echo('=========================================================');
shell.echo('=======*starting dev with fiddler proxy*=========');
shell.echo('=========================================================');
shell.exec('fresh')
