var shell = require('shelljs')
shell.config.verbose = true

shell.exec('scp -r ./tmp/*.gz {{root}}:~/upload/{{name}}/')

shell.exec('ssh {{root}} "sh ~/upload/{{name}}/scripts/deploy.sh"')
