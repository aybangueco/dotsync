# DotSync

A simple cli application to sync your dotfiles from your local repository to specified path

## Motivation

I've been using Hyprland with my own custom dotfiles configuration. I've been enjoying it, but the problem is whenever i have
some changes in my dotfiles, I find it a hassle to synchronize those changes, because I will have to navigate on each directory
and files to sync them on my repository vice versa. With this CLI i built, I aim to eliminate those repetitive steps because i am lazy.

## Installation

```bash
go install github.com/aybangueco/dotsync@latest
```

## Commands

### Init

Generates a configuration file on current directory.

```bash
dotsync init
```

### Sync

Sync files and directory on specified target path.

```bash
dotsync sync
```

## Contributing

1. Fork this repository
2. Clone your forked repository:
3. Create a new branch

```bash
git checkout -b feature/your-feature-name
```

4. Make your changes and commit

```bash
git add .
git commit -m "feat: add your feature description"
```

5. Push your fork
6. Create a pull request (pr) to the main repo
7. Wait for a review and make any requested change
