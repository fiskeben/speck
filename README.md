# micro.blog CLI

A command line tool for managing a micro.blog account.

## Getting started

### Building

The tool is written in Go and currently requires you to have Go installed.

Use `make install` to install the CLI tool. This will create an executable
called `mcro` and add it to your Go bin folder.

### Configuration

You need a token from micro.blog to use the tool.
[Create a new token here](https://micro.blog/account/apps)
and put it in a file called `.microdotblog.yml`, like this:

```yaml
---
token: your-token-here
```

Save the file to either your home directory or current working directory.

## Usage and features

Currently the tool is very limited.
It can retrieve your timeline and create new posts.

### Timeline

Run `mcro` without any parameters to get your timeline.
Pipe the output to `less` to better read it.

### Posting

Run `micro post` to create a new post.
This will open your `$EDITOR` (or `vi` if it's not set)
and you can write your post there. When you save and exit the editor
the post will be created on your micro blog.

If what you wrote is longer than 280 characters,
your editor will repoen and you need to make it shorter.

## TODO and upcoming features

These are some of the features I want to add soon:

* [ ] Pass a file to `mcro post` so that you can write posts independently from posting.
* [ ] A `--dry-run` flag.
* [ ] Parsing the HTML from micro.blog and show posts in a more terminal friendly manner.
* [ ] Multiple accounts.
* [ ] Implement more of the features from the API such as (un)following, reading users' timelines etc.

If you have any requests for new features or just want to give some feedback
create an issue or reach out at [hi@ricco.me](mailto:hi@ricco.me).
