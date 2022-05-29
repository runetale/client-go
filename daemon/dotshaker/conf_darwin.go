package daemon

const SystemConfig = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>

  <key>Label</key>
  <string>com.dotshake.dotshaker</string>

  <key>ProgramArguments</key>
  <array>
    <string>/usr/local/bin/dotshaker</string>
    <string>up</string>
  </array>

  <key>RunAtLoad</key>
  <true/>

  <key>StartInterval</key>
    <integer>5</integer>

  <key>StandardErrorPath</key>
  <string>/usr/local/var/log/dotshaker.err</string>
  <key>StandardOutPath</key>
  <string>/usr/local/var/log/dotshaker.log</string>

</dict>
</plist>
`

const DaemonFilePath = "/Library/LaunchDaemons/com.dotshake.dotshaker.plist"
const BinPath = "/usr/local/bin/dotshaker"
const ServiceName = "com.dotshake.dotshaker"
