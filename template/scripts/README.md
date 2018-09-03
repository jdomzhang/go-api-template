> ls /usr/local/{{name}}

`
{{name}}.service  deploy.sh  install.sh
`

# Important!!!
CR "\r" should be removed in *.sh and *.service files,
Otherwise linux service would not work!!!

## Use command to remove CR
> sed -i 's/\r//g' *.sh

> sed -i 's/\r//g' *.service

# install service

> sh install.sh

# deploy service

> sh deploy.sh

# setup nginx
* Manually setup
> vim /etc/nginx/sites-available/default

> systemctl daemon-reload

> systemctl restart nginx
