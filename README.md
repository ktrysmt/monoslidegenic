# What is this

You just simply create the manuscript by markdown syntax.
Monoslidegenic is a slide generator just only pass one markdown file.

## Features

- Monoslidegenic contains [remarkjs](http://remarkjs.com/).
- Give a single markdown file to it.
- It is output the single HTML file only.
- You can customize CSS and Javascript.

For more information, see `monoslidegenic --help`.

## Usage

### Basic

Prepare your markdown file. and give it to monoslidegenic.  
And open only it by your browser.  

```
$ monoslidegenic ./YOUR_FILENAME.md
$ ls
YOUR_FILENAME.html
$ open ./YOUR_FILENAME.html
```
### Customize CSS or Javascript

You can get default using assets by command options.

```
$ monoslidegenic --output-css > default.css
$ monoslidegenic --output-js > default.js
```

## Installation

Please download from [github release page](https://github.com/ktrysmt/monoslidegenic/releases).

Also you can install by `go get`, too.

    $ go get github.com/ktrysmt/monoslidegenic

## Author

[@ktrysmt](https://twitter.com/ktrysmt)

## License

MIT License

