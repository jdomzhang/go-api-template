systemctl stop {{name}}
unlink /etc/systemd/system/multi-user.target.wants/{{name}}.service
chmod 755 /usr/share/websites/{{name}}/{{name}}
mkdir -p /usr/share/websites/{{name}}
cp ~/upload/{{name}}/scripts/api.service /usr/share/websites/{{name}}/{{name}}.service
ln -s /usr/share/websites/{{name}}/{{name}}.service /etc/systemd/system/multi-user.target.wants/{{name}}.service
systemctl daemon-reload
systemctl restart nginx
