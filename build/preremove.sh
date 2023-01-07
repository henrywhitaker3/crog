systemctl stop crog
systemctl disable crog
rm /etc/systemd/system/crog.service
systemctl daemon-reload
