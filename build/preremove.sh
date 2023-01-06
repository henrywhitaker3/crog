systemctl stop test-foo
systemctl disable test-foo
rm /etc/systemd/system/test-foo.service
systemctl daemon-reload
