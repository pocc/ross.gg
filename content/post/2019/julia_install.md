---
title: "Installing Julia"
date: 2019-03-10T23:32:38Z
author: Ross Jacobs
desc: "Install Julia locally to your $System"
keywords: install
tags:
  - julia
  - setup
  - guide
image: https://dl.dropboxusercontent.com/s/8k0asasb2jjzu6a/install_julia.png

draft: false
---

_Install Julia locally to your $System_

This article covers the local installation of Julia using packages from the
[Download Page](https://julialang.org/downloads/). Other options for using Julia
include an [Online REPL](https://repl.it/languages/julia), a
[Docker image](https://hub.docker.com/_/julia), and a
[Desktop Julia IDE](https://juliacomputing.com/products/juliapro.html).

Installs can either be without sudo/admin _USER_ or _GLOBAL_. For context, most
other programming languages default to being installed with _GLOBAL_ scope.
Choose one and follow the instructins for that target.

## Installation

### Windows

1. (Windows 7/Server 2012 only): Enable
   [Windows Management Framework](https://docs.microsoft.com/en-us/powershell/wmf/overview)
   3+
2. Download the executable
   ([32-bit](https://julialang-s3.julialang.org/bin/winnt/x86/1.1/julia-1.1.0-win32.exe)
   /
   [64-bit](https://julialang-s3.julialang.org/bin/winnt/x64/1.1/julia-1.1.0-win64.exe))
3. Open the executable with Admin Privileges
4. _GLOBAL_ | Change the install directory to `%PROGRAMFILES%/Julia`.
5. Hit Install > Hit Finish
6. _GLOBAL_ | Using Powershell with Admin Privileges, add julia permanently to
   your path

   ```powershell
   $PATH = [Environment]::GetEnvironmentVariable("PATH")
   $julia_path = "C:\Program Files\Julia"

   # -DO THIS- For all users on this machine
   [Environment]::SetEnvironmentVariable("PATH", "$PATH;$julia_path", "Machine")
   # -OR THIS- For just me
   [Environment]::SetEnvironmentVariable("PATH", "$PATH;$julia_path")
   ```

7. _USER_ | Alias `julia` in your \$Profile. If you do not have a profile,
   [Set one up](/post/setup-a-powershell-profile/).

   ```powershell
   $julia_path = $home\AppData\local\Julia-1.1.0\bin\julia
   Add-Content -Path $Profile -Value "function julia { Invoke-Expression $julia_path }"
   ```

### Macos

0. Get [brew](https://brew.sh/) if not installed
1. _GLOBAL_ | `brew cask install julia`
1. _USER_ | Download, un/mount dmg, install to ~/Applications, add julia it to
   PATH

   ```bash
   #!/usr/bin/env bash
   curl https://julialang-s3.julialang.org/bin/mac/x64/1.1/julia-1.1.0-mac64.dmg \
     -o /tmp/julia-1.1.0.dmg
   hdiutil attach /tmp/julia-1.1.0.dmg
   mkdir -p ~/Applications
   cp -r /Volumes/Julia-1.1.0/Julia-1.1.app ~/Applications/Julia-1.1
   ln -s ~/Applications/Julia-1.1/Contents/Resources/julia/bin/julia ~/Applications/julia
   hdiutil detach /Volumes/Julia-1.1.0

   if [[ ":$PATH:" != *":$HOME/Applications:"* ]]; then
   echo "export PATH=$PATH:$HOME/Applications" >> ~/.profile
     source ~/.profile
   fi
   ```

### Linux

1. Download the binary, depending on your architecture

   ```bash
   # Download 64-bit
   curl https://julialang-s3.julialang.org/bin/linux/x64/1.1/julia-1.1.0-linux-x86_64.tar.gz \
     -o /tmp/julia.tar.gz
   # Or 32-bit
   curl https://julialang-s3.julialang.org/bin/linux/x86/1.1/julia-1.1.0-linux-i686.tar.gz \
     -o /tmp/julia.tar.gz
   ```

2. _GLOBAL_ | Download, extract, and copy to /opt/local, add to PATH

   ```bash
   tar -C /tmp -xzf /tmp/julia.tar.gz
   cp -r /tmp/julia /opt/local

   if [[ ":$PATH:" != *":/opt/local:"* ]]; then
     echo "export PATH=$PATH:/opt/local" >> ~/.profile
     source ~/.profile
   fi
   ```

3. _USER_ | Download, extract, copy to ~/bin, and add to PATH

   ```bash
   tar -C /tmp -xzf /tmp/julia.tar.gz
   mkdir -p ~/bin
   cp -r /tmp/julia ~/bin

   if [[ ":$PATH:" != *":~/bin:"* ]]; then
     echo "export PATH=$PATH:~/bin" >> ~/.profile
     source ~/.profile
   fi
   ```

## Verification

Open a new shell to make sure julia==1.1.0 is installed correctly

### Bash

<picture>
    <source type="image/webp" srcset="https://dl.dropboxusercontent.com/s/k4hjgm4wpt3v3vd/bash_julia.webp">
    <source type="image/png" srcset="https://dl.dropboxusercontent.com/s/kwkmv93bzky7dze/bash_julia.png">
    <img
	src="https://dl.dropboxusercontent.com/s/kwkmv93bzky7dze/bash_julia.png"
	alt="Julia Install on Bash" style="height:80%;width:80%;text-align:left;margins:0px"
	Title="Julia Install on Bash">
</picture>

### Powershell

<picture>
    <source type="image/webp" srcset="https://dl.dropboxusercontent.com/s/aw99vi0qst959v9/pwsh_julia.webp">
    <source type="image/png" srcset="https://dl.dropboxusercontent.com/s/9eku6li80jqjftq/pwsh_julia.png">
    <img src="https://dl.dropboxusercontent.com/s/9eku6li80jqjftq/pwsh_julia.png"
	alt="Julia Install on Powershell" style="height:80%;width:80%;text-align:left;margins:0px"
	Title="Julia Install on Powershell">
</picture>

## Next Steps

Now that you have Julia, install some packages!

- [Popular Julia Packages](https://juliaobserver.com/packages)
- [Github-Trending Julia Projects](https://github.com/trending/julia)
