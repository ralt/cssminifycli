# cssminifycli

This is a CSS minifier in Go. It relies on [cssminify][0] for the parsing/minifying job.

It recursively looks for all CSS files in your current working directory and outputs the minified CSS.

## Usage

    cssminify

Yep. It's that simple! You usually want the minified CSS to go in a file though. So you can use it this way:

    cssminify > minified.css

## Install

    git clone https://github.com/Ralt/cssminifycli
    cd cssminifycli
    make
    make install

## Development

    git clone https://github.com/Ralt/cssminifycli
    cd cssminifycli
    make
    ./dist/cssminify

## Contributors

- [Florian Margaine](http://margaine.com)

## License

MIT License.
