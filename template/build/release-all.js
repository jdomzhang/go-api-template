var shell = require('shelljs')

shell.config.verbose = true

// 1
shell.exec('echo step 1............')
shell.exec('yarn release')

// 2
shell.exec('echo step 2............')
shell.exec('yarn upload-weapp')

// 3
shell.exec('echo step 3............')
shell.exec('yarn deploy:nginx')

// done
shell.exec('echo done............')
