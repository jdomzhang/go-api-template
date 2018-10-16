var shell = require('shelljs')
shell.config.verbose = true

shell.exec('scp -r ./tmp/*.gz root@dongfutech.com:~/upload/{{name}}/')

shell.exec('ssh root@dongfutech.com "sh ~/upload/{{name}}/scripts/deploy.sh"')
