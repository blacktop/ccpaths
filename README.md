<p align="center">
  <picture>
  <source media="(prefers-color-scheme: dark)" srcset="docs/logo-dark.png" height="400">
  <source media="(prefers-color-scheme: light)" srcset="docs/logo-light.png" height="400">
  <img alt="Fallback logo" src="docs/logo-dark.png" height="400">
</picture>

  <h4><p align="center">Update <code>compile_commands.json</code> source code paths</p></h4>
  <p align="center">
    <a href="https://github.com/blacktop/ccpaths/actions" alt="Actions">
          <img src="https://github.com/blacktop/ccpaths/actions/workflows/go.yml/badge.svg" /></a>
    <a href="https://github.com/blacktop/ccpaths/releases/latest" alt="Downloads">
          <img src="https://img.shields.io/github/downloads/blacktop/ccpaths/total.svg" /></a>
    <a href="https://github.com/blacktop/ccpaths/releases" alt="GitHub Release">
          <img src="https://img.shields.io/github/release/blacktop/ccpaths.svg" /></a>
    <a href="http://doge.mit-license.org" alt="LICENSE">
          <img src="https://img.shields.io/:license-mit-blue.svg" /></a>
</p>
<br>

## Why? 🤔

Simple utility to fix a common annoyance with building a `compile_commands.json` in a CI server and then trying to use it locally.

## Getting Started

### Install

```bash
go install github.com/blacktop/ccpaths@latest
```

Or download the latest [release](https://github.com/blacktop/ccpaths/releases/latest)

### Running

To replace all the source paths in the JSON with the corrected *local* paths

```bash
❱ ccpaths compile_commands.json /path/to/target /path/in/db
```

### Example

```bash
git clone https://github.com/apple-oss-distributions/xnu.git ~/dev/apple/xnu
cd ~/dev/apple/xnu
wget https://github.com/blacktop/darwin-xnu-build/releases/download/v15.4/xnu-15.4-jsondb.zip
unzip xnu-15.4-jsondb.zip
rm xnu-15.4-jsondb.zip
ccpaths compile_commands.json /Users/runner/work/darwin-xnu-build/darwin-xnu-build ~/dev/apple
```

## License

MIT Copyright (c) 2025 **blacktop**