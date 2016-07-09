# What is this

Monoslidegenic is a slide generator just only pass one markdown file.

## Features

- Monoslidegenic contains [remarkjs](http://remarkjs.com/).
- Give a single markdown file to it.
- It is just only output the single HTML file.
- Also you could customize CSS and Javascript.

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

Also name the output html.


```
$ monoslidegenic --output ./RENAME.html ./YOUR_FILENAME.md
```

### Customize CSS or Javascript

How a replace default setting to your customize file.
Would get default using assets by command options.
And pass it for replace to your codes.

```
$ monoslidegenic --default-css > default.css
$ monoslidegenic --default-js > default.js
$ $EDITOR default.css                                           # could edit it ...
$ monoslidegenic --replace-css default.css ./YOUR_FILENAME.md   # replace to your css
$ monoslidegenic --replace-js default.js ./YOUR_FILENAME.md     # replace to your js
```

## Installation

Please download from [github release page](https://github.com/ktrysmt/monoslidegenic/releases).

Also you could install by `go get`, too.

    $ go get github.com/ktrysmt/monoslidegenic

## Author

[@ktrysmt](https://twitter.com/ktrysmt)

## License

MIT License

