---
title: "Setup A Powershell Profile"
date: 2019-03-12T14:13:07Z
author: Ross Jacobs
desc: "Regular post"
keywords: draft
tags:
  - powershell
  - setup
image: https://www.dl.dropboxusercontent.com/s/nvrhskzzhn1c4tf/ps_profile_setup.png

draft: false
---

_Create a $Profile without Admin Privileges_

Like `~/.bashrc` for bash, a
[Powershell profile](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-6)
allows you to load aliases or functions when you open a new terminal. While
there are technically
[6 Powershell profiles](https://devblogs.microsoft.com/scripting/understanding-and-using-powershell-profiles/),
we are only concerned about `$Profile`, which aliases to
`$Profile.CurrentUserCurrentHost`.

1. Open powershell and add a profile file if it does not exist

   `New-Item -Path $Profile -ItemType "file" *> NUL`

2. (Required only on Windows) Set your
   [ExecutionPolicy](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-6)
   to allow local scripts.

   `Set-ExecutionPolicy RemoteSigned -Scope CurrentUser -Force`

3. (Optional) Edit your profile with `Start notepad $Profile` on Windows /
   `nano $Profile` on *nix. Some people like to add
   [Useful Functions](https://blog.dantup.com/2013/10/useful-powershell-profile-snippets/)
   while others like tricking out their consoles with
   [🎛module paths](http://draith.com/?p=253),
   [⛈️weather](https://dev.to/hf-solutions/how-to-uniquify-your-powershell-profile-2b35),
   and
   [👽aliens](https://blog.ukotic.net/2017/04/12/make-powershell-as-cool-as-you-modify-your-default-profile/).
   I like
   [Windows-Screenfetch](https://github.com/JulianChow94/Windows-screenFetch)
   (Add `Screenfetch 2> $Null` to your `$Profile`)

<picture>
    <source type="image/webp" srcset="https://dl.dropboxusercontent.com/s/yerrw9utxavdju4/ps_screenfetch.webp">
    <source type="image/png" srcset="https://dl.dropboxusercontent.com/s/jxndwak4o1n0on9/ps_screenfetch.png">
    <img src="https://dl.dropboxusercontent.com/s/jxndwak4o1n0on9/ps_screenfetch.png"
	alt="Setting up Powershell Profiles with Screenfetch"
	style="height:100%;width:100%;margin:0px"
	title="Screenfetch can take a while to load...">
</picture>
