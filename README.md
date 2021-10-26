# gocart
Golang Configuration and Revision Tracker

## Proposed Structure

### Starting a Go Cart

1. Create a git (or other SCM) repo
1. `gocart init`

Gocart init will

1. Create a `.gocart/` folder for storing configuration
1. Create a `.gocart/symlink_mapping` file for storing the mapping of `$CONFIG_NAME` to `$OLD_PATH`

### Adding Files

`gocart add $PATH_TO_CONFIG_FILE $CONFIG_NAME`

Add a new config file at the path to gocart's tracking.

1. copy `$PATH_TO_CONFIG` to `$GOCART_PATH/$CONFIG_NAME`
1. remove the file at `$PATH_TO_CONFIG`
1. place a symlink to `$GOCART_PATH/$CONFIG_NAME` in the old file's path, where applications expecting a config file will expect it

### Removing Files

`gocart remove $CONFIG_NAME`

Remove a file from gocart's tracking and place the tracked file at it's original path.

1. remove the symlink located at `$PATH_TO_CONFIG`
1. move the config file at `$GOCART_PATH/$CONFIG_NAME` back to `$PATH_TO_CONFIG` 
