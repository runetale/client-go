package daemon

var SystemConfig = `[Unit]
Description=Wissy Daemon
Requires=NetworkManager.service
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
User=root
Type=simple
ExecStart=/bin/wissy up
Restart=on-failure
RestartSec=15s

[Install]
WantedBy=multi-user.target
`

const DaemonFilePath = "/etc/systemd/system/wissy.service"
const BinPath = "/bin/wissy"
const ServiceName = "wissy"
