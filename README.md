
birthday-manager-cli
====================

birthday-manager-cli is small cli tool for birthday managing. It stores and shows birthdays.

Commands
--------

```
create   Create birthday
remove   Remove birthdays
show     Print name, age and next birthday date
help, h  Shows a list of commands or help for one command
```

.bashrc file
-----------

```
### Birthdays
date
birthday-manager-cli \
--file $HOME'/.config/birthday-manager-cli/conf.json' \
show --duration 240h
```