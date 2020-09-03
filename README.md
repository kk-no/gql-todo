# GraphQL-TODO

### inithialize

$ go mod init github.com/[username]/[repogitory]  
$ go get github.com/99designs/gqlgen
$ go run github.com/99designs/gqlgen init

[GqlGen Official](https://gqlgen.com/getting-started/)

### run

$ make run
$ make test

## memo

```
mutation createTodo {
  createTodo(input:{text:"add user", userId: "test1"}) {
    user {
      id
      name
    }a
    text
    done
  }
}
```
