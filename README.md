# Hostrewind

## Install
```shell
go get -u github.com/kwiesmueller/hostrewind/bin/hostrewind
```

```shell
make prepare
make install
```

## Params
* `-sm`: handles the custom suffix smhss.de (remove or append when necessary)

## Usage Samples
Not that the configured custom domain is always being stripped when being at the start
```
ssh $(echo "lf.office.build" | hostrewind -sm)
-> build.office.lf.smhss.de

ssh $(echo "build.office.lf.smhss.de" | hostrewind)
-> lf.office.build
```

## Example Aliases
```
function sssh() {
ssh $(echo "$1" | hostrewind -sm)
}
```