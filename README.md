# Monoslidegenic

This is a slide generator just only pass one markdown file.

## Features

- Monoslidegenic contains [remarkjs](http://remarkjs.com/) in the standard.
- You could give a single markdown file to it.
- It is output the single HTML file only.
- Also you can use your customized CSS or Javascript for the file.

For more information, see `monoslidegenic --help`.

## Usage

### Basic

Prepare your markdown file. and give it to monoslidegenic.  
And open only it by your browser.  

```
$ monoslidegenic ./YOUR_FILE.md
$ ls
YOUR_FILE.md YOUR_FILE.html
$ open ./YOUR_FILE.html
```
Name the slide by yourself, too.

```
$ monoslidegenic --output awesome-slide.html ./YOUR_FILE.md
```

### Customize CSS or Javascript

You can get default using assets by command options.

```
$ monoslidegenic --css awesome.css --js awesome.js ./YOUR_FILE.md
```

## Installation

Please download from [github release page](https://github.com/ktrysmt/monoslidegenic/releases).

Also you can install via `go get`, too.

    $ go get github.com/ktrysmt/monoslidegenic

## Author

[@ktrysmt](https://twitter.com/ktrysmt)

## License

MIT License

