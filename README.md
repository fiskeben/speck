# speck - a micro.blog CLI

`speck` is a command line tool for managing a micro.blog account.

## Getting started

### Building

The tool is written in Go and currently requires you to have Go installed
in order to build the binary.

`make build` will build the binary and put it in this folder.

Use `make install` to install `speck`. This will create an executable
called `speck` and add it to your Go bin folder.

### Configuring `speck`

You need a token from micro.blog to use the tool.
[Create a new token here](https://micro.blog/account/apps)
and put it in a file called `.speck.yml`, like this:

```yaml
---
token: your-token-here
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
* [ ] Parsing the HTML from micro.blog and show posts in a more terminal friendly manner.
* [x] Adding help text to the editor (like `git commit`)
* [x] Save the post locally as well as posting it to micro.blog.
* [ ] Multiple accounts.
* [ ] Implement more of the features from the API such as (un)following, reading users' timelines etc.
* [ ] Open user's micro.blog profile/timeline in a browser.
* [ ] Some sort of CI/CD and Homebrew cask for easy installation.

If you have any requests for new features or just want to give some feedback
create an issue or reach out at [hi@ricco.me](mailto:hi@ricco.me).
You can also [follow me on micro.blog](https://micro.blog/ricco).