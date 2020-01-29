# Rocket Bot Go

This is an API for a Rocket Chat bot written in go. The library for this rocket bot is in the `rocket` folder. `main.go` is a demo implementation of a chat bot.

#### Building Demo

Before building the demo, ensure that the _go_ compiler is installed and that both the _yaml_ and _websocket_ libraries are installed. Then use the _go_ compiler to build the executable. Assuming the _go_ compiler is installed, these 3 commands will compile the demo.

```
go get gopkg.in/yaml.v2
go get github.com/gorilla/websocket
go build main.go
```
#### License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
