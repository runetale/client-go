package daemon

const SystemConfig = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>

  <key>Label</key>
  <string>com.wissy.wissy</string>

  <key>ProgramArguments</key>
  <array>
    <string>/usr/local/bin/wissy</string>
    <string>up</string>
  </array>

  <key>RunAtLoad</key><true/>
  <key>AbandonProcessGroup</key><true/>
  <key>StartInterval</key>
    <integer>15</integer>

</dict>
</plist>
`

const DameonFilePath = "/Library/LaunchDaemons/com.wissy.wissy.plist"
const BinPath = "/usr/local/bin/wissy"
const ServiceName = "com.wissy.wissy"
