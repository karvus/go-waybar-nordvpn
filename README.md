# go-waybar-nordvpn

Simple wrapper around the Linux CLI for [NordVPN](https://nordvpn.com/
"NordVPN"), suitable for use in a [Waybar](https://github.com/Alexays/Waybar
"Waybar") custom module.

For example, in `~/.config/waybar/config`

```json
{
 ...
     "custom/nordvpn": {
        "interval": 5,
        "format": "VPN{}",
        "exec": "~/.config/waybar/modules/go-waybar-nordvpn exec",
        "on-click": "~/.config/waybar/modules/go-waybar-nordvpn toggle",
        "return-type": "json",
    },
 ...
}
```

