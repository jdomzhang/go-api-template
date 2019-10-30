var shell = require('shelljs')
shell.config.verbose = true

shell.exec(
  'scp scripts/api.conf {{root}}:/etc/nginx/sites-enabled/{{name}}.conf'
)
