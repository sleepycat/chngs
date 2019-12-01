# chngs

This is a command line tool for monitoring a directory (and it's contents) for
file system events. There are [other
ways](https://superuser.com/questions/181517/how-to-execute-a-command-whenever-a-file-changes)
to do this, but I had this problem and I'm learning Go... this is the result.

## Usage

```sh
chngs [path to directory you want to monitor]
```

If no path is supplied, it's assumed that you want to monitor `.`, aka the current directory.

## Installation

```sh
go build
go install
```
