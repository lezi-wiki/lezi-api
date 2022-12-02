# Lezi Api

A fast API program for people who are silly and funny.

## Demo

[https://api.lezi.wiki/](https://api.lezi.wiki/)

## API Docs

[Wiki pages](https://github.com/lezi-wiki/lezi-api/wiki) - [Docs v1](https://github.com/lezi-wiki/lezi-api/wiki/LeziAPI-Docs-v1)

## Use

### Run

Go to [Releases](https://github.com/lezi-wiki/lezi-api/releases) to download the corresponding version of the program.Then unpack it and get the main program.

Start command(Linux):

```shell
./lezi-api
```

### Build

```shell
git clone https://github.com/lezi-wiki/lezi-api
cd lezi-api
go run github.com/google/wire/cmd/wire@latest ./...
go build -o ../lezi-api
cd ..
```

## License

under [GPL-3.0 license](https://github.com/lezi-wiki/lezi-api/blob/master/LICENSE)
