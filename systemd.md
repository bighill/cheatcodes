# Daemonizing services with `systemd`

Create service `/etc/systemd/system/foo.service`

_User & Group are optional. Default user is root._

```bash
[Unit]
Description=FooService

[Service]
ExecStart=/path/to/foo/scripts/run.sh
User=foouser
Group=foouser

[Install]
WantedBy=multi-user.target
```

Enable the service so that it can now start and will start on reboot

```bash
sudo systemctl enable foo.service
```

Manage the service

```bash
sudo systemctl start foo.service
sudo systemctl stop foo.service
sudo systemctl restart foo.service

# or

sudo service foo start
sudo service foo stop
sudo service foo restart
```
