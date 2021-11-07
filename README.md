# gocart

Golang Configuration and Revision Tracker

## Why?

To make it easier to maintain your dotfiles across hosts ands platforms. It was built for use in tandem with git, but any SCM tool should suffice. 

## How?

### Install

1. clone it

2. `go build`

3. copy to path

4. ???

5. Profit

### Setup

1. `mkdir -p ~/src/dotfile_repo` or wherever you want your system config to live

2. `cd ~/src/dotfile_repo`

3. `git init`

4. `gocart repoInit`

5. `gocart platformSet $YOUR_OS_NAME` could be osx, debian, bsd...

6. `git add .gocart.json`

7. `git commit -m "started versioning my dotfiles with gocart and I got a raise at work"`

8. `git push` tell all your friends!

### Track a config

1. `gocart configAdd -n vimrc -p ~/.vimrc`

2. `git add vimrc .gocart.json`

3. `git commit -m "tracking my vimrc with gocart, now I can run a 6 minute mile!"`

### Remove a config

1. `gocart configDel vimrc`

2. `git add .gocart.json vimrc`

3. `commit -m "remove vimrc from gocart tracking, I stubbed my toe after I got up from my workspace"`

## ...but How does it?

Gocart works by copying your config files into your working directory, deleting the originals and replacing them with symlinks to the new copy. To remove a config from tracking, it does the reverse. When used inside of a git repository, you can version your personal configuration. You can also use Github or Gitlab to share your dotfiles across machines.

## TODO:

* Version Json Schema

* Setup detailed logging

* Add command to install dotfiles system-wide after `git clone`

* Add nested `kubectl` style subcommands

### Notes

I built this mainly as a project to learn `Go`, and solve a marginally painful problem in my day-to-day work as a software engineer. 

If you find this tool useful, please tell your friends, and consider submitting a feature request to the issue tracker, or opening a PR on an existing issue. 
