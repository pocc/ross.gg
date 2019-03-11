---
title: "Installing Julia"
date: 2019-03-10T23:32:38Z
author: Ross Jacobs
desc: "Install Julia locally to your $System"
keywords: install
tags:
  - draft
image: http://www.quickmeme.com/img/92/927d52fd29f08027c5356e5f8bfd78021dcd2351d18d717eb86d393132f7322a.jpg

draft: true
---

- [ ] Images: Cover image

_Install Julia locally to your $System_

This article covers the local installation of Julia using packages from the
[Download Page](https://julialang.org/downloads/). Other options for using Julia
include an [Online REPL](https://repl.it/languages/julia), a
[Docker image](https://hub.docker.com/_/julia), and a
[Desktop Julia IDE](https://juliacomputing.com/products/juliapro.html).

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
4. Change the install directory to `%PROGRAMFILES%/Julia`. The default of
   `%APPDATA%/local/julia-1.1.0` assumes you do not have Admin access.
5. Hit Install > Hit Finish
6. Using Powershell with Admin Privileges, add julia to your path for all users

   ```powershell
   $PATH = [Environment]::GetEnvironmentVariable("PATH")
   $julia_path = "C:\Program Files\Julia"
   [Environment]::SetEnvironmentVariable("PATH", "$PATH;$julia_path", "Machine")
   ```

### Macos

0. Get [brew](https://brew.sh/) if not installed
1. `brew cask install julia`

### Linux

1. Download the binary
   ([32-bit](https://julialang-s3.julialang.org/bin/linux/x86/1.1/julia-1.1.0-linux-i686.tar.gz)
   /
   [64-bit](https://julialang-s3.julialang.org/bin/linux/x64/1.1/julia-1.1.0-linux-x86_64.tar.gz))
2. Navigate to the download in your terminal and create a user copy of julia.

   ```bash
   tar -xvzf $filename
   mkdir -p ~/bin
   cp -r julia-1.1.0 ~/bin/julia
   ```

3. Add ~/bin to path if it's not already

   ```bash
   if [[ ":$PATH:" != *":~/bin:"* ]]; then
     echo "export PATH=$PATH:~/bin" >> ~/.profile
     source ~/.profile
   fi
   ```

## Verification

Open a shell and ensure julia==1.1.0 installed correctly

### Bash

<picture>
    <source type="image/webp" srcset="https://dl.dropboxusercontent.com/s/k4hjgm4wpt3v3vd/bash_julia.webp">
    <source type="image/png" srcset="https://dl.dropboxusercontent.com/s/kwkmv93bzky7dze/bash_julia.png">
    <img
	src="https://dl.dropboxusercontent.com/s/kwkmv93bzky7dze/bash_julia.png"
	alt="Julia Install on Bash" style="height:80%;width:80%;text-align:left;margins:0px">
</picture>

### Powershell

<picture>
    <source type="image/webp" srcset="https://dl.dropboxusercontent.com/s/aw99vi0qst959v9/pwsh_julia.webp">
    <source type="image/png" srcset="https://dl.dropboxusercontent.com/s/9eku6li80jqjftq/pwsh_julia.png">
    <img src="https://dl.dropboxusercontent.com/s/9eku6li80jqjftq/pwsh_julia.png"
	alt="Julia Install on Powershell" style="height:80%;width:80%;text-align:left;margins:0px">
</picture>

## Next Steps

Now that you have Julia, install some packages!

- [Popular Julia Packages](https://juliaobserver.com/packages)
- [Github-Trending Julia Projects](https://github.com/trending/julia)
