package daemon

var SystemConfig = `[Unit]
Description=dotshaker daemon
Requires=NetworkManager.service
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
User=root
Type=simple
ExecStart=/bin/dotshaker
Restart=on-failure
RestartSec=15s

[Install]
WantedBy=multi-user.target
`

const DaemonFilePath = "/etc/systemd/system/dotshaker.service"
const BinPath = "/bin/dotshaker"
const ServiceName = "dotshaker"
