<p align="center">
<img src="https://raw.githubusercontent.com/wiki/akiyosi/gonvim/images/gopher-with-neovim.png" width="180">
</p>

# Gonvim

[![Join the chat at https://gitter.im/gonvim/gonvim](https://badges.gitter.im/gonvim/gonvim.svg)](https://gitter.im/gonvim/gonvim?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Gonvim is a Neovim GUI written in Golang, using a [Golang qt backend](https://github.com/therecipe/qt).
 
This repository forked from the original [Gonvim](https://github.com/dzhou121/gonvim) for the purpose of maintenance and enhancement.

![](https://raw.githubusercontent.com/wiki/akiyosi/gonvim/images/0.3.0.png)

## Features

* [Modern UI](https://github.com/akiyosi/gonvim/wiki/Features#tabline-statusline-lint-message-command-line-and-message)
* [Fuzzy Finder](https://github.com/akiyosi/gonvim/wiki/Features#fuzzy-finder-in-gui)
* [Markdown Preview](https://github.com/akiyosi/gonvim/wiki/Features#markdown-preview)
* [MiniMap](https://github.com/akiyosi/gonvim/wiki/Features#minimap)
* [Dein.vim GUI](https://github.com/akiyosi/gonvim/wiki/Features#deinvim-gui)

<br>

## Getting Started
Pre-built packages for Windows, macOS, and Linux are found at the [Releases](https://github.com/akiyosi/gonvim/releases) page.

### MacOS
gonvim.app looks for the nvim process from the following.

```
/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/opt/local/bin:/opt/local/sbin
```

Deployment example:
```
cd /path/to
curl -LO https://github.com/neovim/neovim/releases/download/nightly/nvim-macos.tar.gz
tar xf nvim-macos.tar.gz
ln -s /path/to/bin/nvim /usr/local/bin/nvim
```

### Linux
Deploy Neovim under the `$PATH`.

### Windows
Add Gonvim `bin` path to `%PATH%` environment variable.


<br>

## Documenation

* [Configurations](https://github.com/akiyosi/gonvim/wiki/Configurations)
* [Development](https://github.com/akiyosi/gonvim/wiki/Development)

<br>

## Credits
* Gonvim was created by dzhou121 ([https://github.com/dzhou121/gonvim](https://github.com/dzhou121/gonvim))
* Gopher in Gonvim image was created by Takuya Ueda ([https://twitter.com/tenntenn](https://twitter.com/tenntenn))


