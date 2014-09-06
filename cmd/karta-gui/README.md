# Karta GUI

Experimental GUI for the [karta](https://github.com/peterhellberg/karta) library.

![Karta GUI](http://assets.c7.se/skitch/Karta_QML-20140906-165833.png)

## Dependencies

On Mac OS X you'll need [Qt 5](https://qt-project.org/wiki/Qt_5.0).
It's easiest to install with [Homebrew](http://brew.sh/).

Install the qt5 and pkg-config packages:

    $ brew install qt5 pkg-config

Then, force brew to "link" qt5 (this makes it available under /usr/local):

    $ brew link --force qt5

And finally, fetch and install go-qml:

    $ go get -u gopkg.in/qml.v1

## Karta.app

I’ve provided a skeleton app bundle that wraps the `karta-gui` binary.

Build the `karta-gui` executable inside the bundle like this:

    $ make build
