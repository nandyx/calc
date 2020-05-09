

# Install

1. Download stable version
[https://golang.org/dl/](https://golang.org/dl/)
2. vscode extension
[https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

## Run app

```
go run main.go
```

## Results

success result:
```
----------------------------
----- Addition console -----
----------------------------
2+3
resultado: 5
```
fail results:
```
----------------------------
----- Addition console -----
----------------------------
2&5
La operación no es una adición
```
```
----------------------------
----- Addition console -----
----------------------------
a+b
Error: Error de sintaxis
```
## Tests
```
go test -cover
```


