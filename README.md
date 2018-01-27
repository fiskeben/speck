# speck - a micro.blog CLI

`speck` is a command line tool for managing a micro.blog account.

## Getting started

### Install with `homebrew`

You can install `speck` with
[Homebrew](https://brew.sh)
in two easy steps
(given that you have Homebrew installed):

1. Add my tap: `brew tap fiskeben/homebrew-tap`
1. Install `speck`: `brew install speck`

Don't forget to read the paragraph about configuration below.

### Building and installing manually

The tool is written in Go and currently requires you to have Go installed
in order to build the binary.

`make build` will build the binary and put it in this folder.

Use `make install` to install `speck`. This will create an executable
called `speck` and add it to your Go bin folder.
Add the folder to your `$PATH` if you want.

### Configuring `speck`

You need a token from micro.blog to use the tool.
[Create a new token here](https://micro.blog/account/apps)
and put it in a file called `.speck.yml`,
along with your username, like this:

```yaml
---
username: your-micro-blog-username
token: your-token
```

Save the file to either your home directory or current working directory.

## Usage and features

Currently `speck` is very limited.
It can retrieve your timeline and create new posts.

### Timeline

Run `speck` without any parameters to get your timeline.
Pipe the output to `less` to better read it.

### Posting

Run `speck post` to create a new post.
This will open your `$EDITOR` (or `vi` if it's not set)
and you can write your post there. When you save and exit the editor
the post will be created on your micro blog.

If what you wrote is longer than 280 characters,
your editor will repoen and you need to make it shorter.

It's also possible to post something you already have on file,
like so: `speck post <path/to/file>`

## TODO and upcoming features

These are some of the features I want to add soon:

* [x] Pass a file to `speck post` so that you can write posts independently from posting.
* [x] A `--dry-run` flag.
* [ ] Post display:
  * [ ] Parsing/stripping HTML
  * [ ] Limiting output width to avoid super long lines of text
  * [ ] URL to photos.
* [x] Adding help text to the editor (like `git commit`)
* [x] Save the post locally as well as posting it to micro.blog.
* [ ] Multiple accounts.
* [ ] Setup command that writes the `.speck.yml` config.
* [ ] Implement more of the features from the API such as ~(un)following~, reading users' timelines etc.
* [ ] Open user's micro.blog profile/timeline in a browser.
* [ ] Don't be an "app" with lots of friendly, formatted output. Output should be possible to pipe to another program.
* [ ] Use `stdin` as source for post text.
* [ ] Some sort of CI/CD and ~Homebrew cask for easy installation~.

If you have any requests for new features or just want to give some feedback
create an issue or reach out at [hi@ricco.me](mailto:hi@ricco.me).
You can also [follow me on micro.blog](https://micro.blog/ricco).