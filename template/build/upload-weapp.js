var shell = require('shelljs')
shell.config.verbose = true

shell.exec('ssh {{root}} " mkdir -p /usr/share/websites/{{name}}/userdata/weapp "')
shell.exec('scp -r ./userdata/weapp/res {{root}}:/usr/share/websites/{{name}}/userdata/weapp/')
