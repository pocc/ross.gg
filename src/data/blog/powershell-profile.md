---
title: "Setup A Powershell Profile"
pubDatetime: 2019-03-01T14:13:07Z
description: "Create and customize a PowerShell $Profile without admin privileges."
tags:
  - powershell
  - setup
---

![](@/assets/images/powershell-profile/ps_profile_setup.webp)

Like `~/.bashrc` for bash, a [Powershell profile](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-6) allows you to load aliases or functions when you open a new terminal.

## Table of contents

## Create a $Profile without Admin Privileges

1. Create a profile file if it does not exist:

```powershell
New-Item -Path $Profile -ItemType "file"
```

2. (Windows only) Set ExecutionPolicy to allow local scripts:

```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser -Force
```

3. (Optional) Edit your profile with `Start notepad $Profile` and add your customizations. For example, display system info or ASCII art on launch:

![PS Screenfetch](@/assets/images/powershell-profile/ps_screenfetch.webp)

![PS Profile Alien](@/assets/images/powershell-profile/ps_profile_alien.webp)
