# clash subscription updater
> Update the clash `config.yaml` peroidly with optional patch

## Install
### ArchLinux
Install from AUR `clash-subscription-updater-git`
```sh
yay -S clash-subscription-updater-git
```

## Usage
```shell
-d, --clash-config-dir string    config directory of clash (default "/home/fengkx/.config/clash")
-h, --help                      show this message
-i, --interval int              interval to fetch configuration (minutes) (default 60)
    --override                  override the existed config file
-v, --version                   show current version

```

It will init a config file in `$HOME/.config/clash-subscription-updater.yaml`
you can add additional clash configs in the file to patch(prepend) to the subscription.

for example
```yaml
clash-config-dir: /home/fengkx/.config/clash
help: true
interval: 60
override: true
proxies:
- name: NeteaseMusic
  port: 9726
  server: 127.0.0.1
  type: http
rules:
- DOMAIN-SUFFIX,163.com,NeteaseMusic,
subscription-url: https://clash-rule-set-flatten.vercel.app/flat?url=xxxxxxxxx
```
`proxies` and `rules` will prepend to existed field

Only `proxies` and `rules` can be patched for now.
