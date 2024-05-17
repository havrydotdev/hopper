# Hopper
> Robust vanilla minecraft server

Minecraft server core with focus on customization, performance and minimal hardware requirements.

## Installation

### To install hopper, check [releases](https://github.com/gavrylenkoIvan/hopper/releases) page

## Usage example

To start Hopper server, install latest release binary from releases page and execute it

_For config settings and more usage examples, please refer to the [Wiki][wiki]._

## Development setup

Hopper requires [Go](https://go.dev/dl/) 1.16+ installed since it uses `go mod` dependencies.

```sh
# Run hopper
make run
```

To enable hot reload, [air](https://github.com/cosmtrek/air) is required

```sh
# Run hopper in hot reload mode
make hot
```

## Release History

* 0.0.1
    * Implement [server list ping](https://wiki.vg/Server_List_Ping)

## Contributing

1. Fork it (<https://github.com/gavrylenkoIvan/hopper>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[wiki]: https://github.com/gavrylenkoIvan/hopper/wiki
