
birthday-manager-cli
====================

Description
-----------

birthday-manager-cli is small cli tool for birthday managing. It stores and shows birthdays.

Commands
--------

```
create   Create birthday
remove   Remove birthdays
show     Print name, age and next birthday date
help     Shows a list of commands or help for one command
```

Install
-------

```bash
$ go get github.com/ivan1993spb/birthday-manager-cli
$ cd $GOPATH/src/github.com/ivan1993spb/birthday-manager-cli
$ go build
$ go install
```

Usage
-----

### First example script

```bash
### Birthdays
#!/bin/bash
echo "Birthdays:"
birthday-manager-cli \
--file $HOME'/.config/birthday-manager-cli/conf.json' \
show --duration 240h
```

### Second example script

```bash
#!/bin/bash

birthdays="$(birthday-manager-cli --file $HOME'/.config/birthday-manager-cli/conf.json' show --duration 240h)"
title="Birthday"

if [ `wc -l <<< "$birthdays"` -gt 1 ] ; then
	title="$title"s
fi

notify-send --icon="$GOPATH/src/github.com/ivan1993spb/birthday-manager-cli/icon.png" "$title" "$birthdays"

```

example conf.json
-----------------

```json
[
	{
		"name": "Ivan",
		"time": "02 Jan 93 00:00 MSK"
	},
	{
		"name": "Masha",
		"time": "26 Feb 94 00:00 MSK"
	},
	{
		"name": "Kolya",
		"time": "26 Apr 97 00:00 MSK"
	}
]
```
