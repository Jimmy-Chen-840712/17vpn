# 17vpn

![image](https://user-images.githubusercontent.com/91862792/172811759-851153ee-8e76-4e77-a45a-a11504dce767.png)


### Pre-Installation

Follow the [confluence](https://17media.atlassian.net/wiki/spaces/H/pages/1027244286/OKTA+Pritunl+VPN) install Pritunl client and import profiles first

### Installation

```shell
git cloen git@github.com:Jimmy-Chen-840712/17vpn.git
cd 17vpn
go build -o 17vpn main.go
mv 17vpn /usr/local/bin
```

### Usage

```shell
# Initial your OTP key and Pin (first time)
# Enter ID or Server to connect/disconnect
$ 17vpn

# Connect to the profile directly
$ 17vpn c 1

# Disconnect all connections
$ 17vpn d
```
