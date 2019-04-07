---
title: "Tshark Decryption"
date: 2019-03-12T12:44:45Z
author: Ross Jacobs
desc: "Tshark Decryption for Kerberos, TLS, and 802.11"
tags:
  - networking
  - tshark
  - cryptography
image: https://allabouttesting.org/wp-content/uploads/2018/06/tshark-count.jpg

draft: true
---

<!-- Draft Until
* [ ] Asciinema of TLS 1.2 Decryption
* [ ] tshark Usage of WPA2-PSK decryption
* [ ] WPA2-PSK Asciinema
-->

# Decryption

There are many protocols that can be decrypted in Wireshark:

## [Kerberos](https://wiki.wireshark.org/Kerberos)

- Use [this guide](https://docs.axway.com/bundle/APIGateway_762_IntegrationKerberos_allOS_en_HTML5/page/Content/KerberosIntegration/Wireshark/wireshark_tracing_for_spnego_kerberos_auth_between.htm)
to generate a keytab file.  
- To use this keytab file for decryption, use
`tshark -r /path/to/file -K /path/to/keytab`.

## TLS 1.2 Decryption

It is possible to decrypt the data on the client side if SSL logging is
enabled. Chrome and firefox will check whether the $SSLKEYLOGFILE
environmental variable exists, and if it does, will send keys to the file.
Using tshark and firefox, we will be able to extract the html file. 

### 1. Add the SSLKEYLOGFILE environment variable

```bash
echo "export SSLKEYLOGFILE=/tmp/sslkey.log" >> ~/.bashrc
source ~/.bashrc
```

### 2. Capture traffic going to a website

Let's use
https://ss64.com as it uses TLSv1.2 and is [designed to be
lightweight](https://ss64.com/docs/site.html). They have an article on netcat, which seems apropos to use: `https://ss64.com/bash/nc.html`.

```bash
cd /tmp
url='https://ss64.com/bash/nc.html'
tshark -Q -w /tmp/myfile.pcapng & tpid=$!
firefox --headless --private $url & ffpid=$!
sleep 10 && kill -9 $tpid $ffpid
```

### 3. Export http objects to `obj/`

```bash
mkdir -p /tmp/obj
# Equivalent to Wireshark > File > Export Objects > HTTP
tshark --export-objects http,/tmp/obj -r /tmp/myfile.pcapng \
  -o tls.keylog_file:$SSLKEYLOGFILE
```

### 4. Verify that HTML extraction was successful

The two relevant files that we receive from ss64.com are `nc.html` and
`main.css`. This HTML file references its css file as "../main.css", so
create a symbolic link for verification purposes and then open it.

```bash
ln -s obj/main.css main.css
firefox --browser obj/nc.html
```

If all is well, your local version of nc's manpage looks exactly the same
as ss64's version.

## TLS 1.3 Decryption

TLS 1.3 is the next iteration after industry standard 1.2, with 1.3 adopted
by [most browsers](https://caniuse.com/#feat=tls1-3) at this point. TLS
decryption is currently broken ([bug
15537](https://bugs.wireshark.org/bugzilla/show_bug.cgi?id=15537)) when
certificate message spans multiple records. In my testing, some javascript
files (and other small files) get decrypted, but no html or css files.

## WPA2 Decryption

_Note that WPA3 decryption is [still in development](https://seclists.org/wireshark/2019/Mar/79)._

Use [Rasika Nayanajith's guide] for WPA2-PSK decryption.