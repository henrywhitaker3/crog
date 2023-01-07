systemctl stop healthcheck
systemctl disable healthcheck
rm /etc/systemd/system/healthcheck.service
systemctl daemon-reload
