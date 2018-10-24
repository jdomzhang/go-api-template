var shell = require('shelljs')

shell.config.verbose = true

// 0
shell.exec('echo step 0............')
shell.exec('yarn build')

// 1
shell.exec('echo step 1............')
shell.exec('ssh root@dongfutech.com " mkdir -p ~/upload/{{name}}/scripts "')
shell.exec('ssh root@youxin-tea.edesoft.com " mkdir -p /usr/share/websites/youxin-teaapi "')

shell.exec('scp -r ./scripts/*.* root@dongfutech.com:~/upload/{{name}}/scripts')

// 2
shell.exec('echo step 2............')
shell.exec(`ssh root@dongfutech.com \" sed -i 's/\\r//g' ~/upload/{{name}}/scripts/*.sh & sed -i 's/\\r//g' ~/upload/{{name}}/scripts/*.service \"`)

// 3
shell.exec('echo step 3............')
shell.exec(`ssh root@dongfutech.com \" sh ~/upload/{{name}}/scripts/install.sh \"`)

// 4
shell.exec('echo step 4............')
shell.exec('yarn upload')
