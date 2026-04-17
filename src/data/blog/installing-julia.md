---
title: "Installing Julia"
pubDatetime: 2019-03-10T23:32:38Z
description: "Install Julia locally on Windows, macOS, or Linux — globally or per-user."
tags:
  - julia
  - setup
---

![](/img/julia/install/install_julia.png)

_Install Julia locally to your system_

This article covers the local installation of Julia using packages from the [Download Page](https://julialang.org/downloads/).

Installs can either be without sudo/admin _USER_ or _GLOBAL_. Choose one and follow the instructions for that target.

## Table of contents

## Windows

1. Download the executable ([32-bit](https://julialang-s3.julialang.org/bin/winnt/x86/1.1/julia-1.1.0-win32.exe) / [64-bit](https://julialang-s3.julialang.org/bin/winnt/x64/1.1/julia-1.1.0-win64.exe))
2. Open with Admin Privileges
3. _GLOBAL_ | Change install directory to `%PROGRAMFILES%/Julia`
4. _GLOBAL_ | Add julia to PATH:

```powershell
$PATH = [Environment]::GetEnvironmentVariable("PATH")
$julia_path = "C:\Program Files\Julia"
[Environment]::SetEnvironmentVariable("PATH", "$PATH;$julia_path", "Machine")
```

5. _USER_ | Alias `julia` in your [$Profile](/posts/powershell-profile/)

## macOS

```bash
brew cask install julia   # GLOBAL
```

## Linux

```bash
curl https://julialang-s3.julialang.org/bin/linux/x64/1.1/julia-1.1.0-linux-x86_64.tar.gz \
  -o /tmp/julia.tar.gz
tar -C /tmp -xzf /tmp/julia.tar.gz
sudo cp -r /tmp/julia /opt/local   # GLOBAL
# Or: cp -r /tmp/julia ~/bin       # USER
```

## Verification

<picture>
    <source type="image/webp" srcset="/img/julia/install/bash_julia.webp">
    <img src="/img/julia/install/bash_julia.png" alt="Julia Install on Bash" style="height:80%;width:80%">
</picture>
