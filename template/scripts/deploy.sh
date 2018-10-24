mkdir -p /usr/share/websites/{{name}}
systemctl stop {{name}}
tar --overwrite --overwrite-dir -C /usr/share/websites/{{name}} -xvzf ./{{name}}.tar.gz
# rm /usr/share/websites/{{name}}/*.tar
chmod 755 /usr/share/websites/{{name}}/{{name}}
#PORT=3009 /usr/share/websites/{{name}}/{{name}}
systemctl start {{name}}
systemctl status {{name}}