package daemon

var SystemConfig = `[Unit]
Description=dotshaker daemon
Requires=NetworkManager.service
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
User=root
Type=simple
ExecStart=/bin/dotshake up
Restart=on-failure
RestartSec=15s

[Install]
WantedBy=multi-user.target
`

const DaemonFilePath = "/etc/systemd/system/dotshake.service"
const BinPath = "/bin/dotshake"
const ServiceName = "dotshake"
