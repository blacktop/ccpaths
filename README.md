<p align="center">
  <a href="https://github.com/blacktop/ccpaths"><img alt="ccpaths Logo" src="https://raw.githubusercontent.com/blacktop/ccpaths/main/docs/logo.png" height="200" /></a>
  <h1 align="center">ccpaths</h1>
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

Simple utility to fix a common annoyance with building a compile_commands.json in a CI server and then trying to use it locally.

## Getting Started

### Install

```bash
go install github.com/blacktop/ccpaths@latest
```

### Running

To replace all the source paths in the JSON with the corrected *local* paths

```bash
❱ ccpaths compile_commands.json /path/to/target /path/in/db
```

## License

MIT Copyright (c) 2025 **blacktop**