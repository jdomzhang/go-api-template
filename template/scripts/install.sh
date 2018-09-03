systemctl stop {{name}}
unlink /etc/systemd/system/{{name}}.service
chmod 755 /usr/share/websites/{{name}}/{{name}}
mkdir -p /usr/share/websites/{{name}}
cp ~/upload/{{name}}/scripts/api.service /usr/share/websites/{{name}}/{{name}}.service
ln -s /usr/share/websites/{{name}}/{{name}}.service /etc/systemd/system/{{name}}.service
systemctl daemon-reload
#systemctl start {{name}}
#systemctl status {{name}}
