# GraphQL examples

## create Todo
```
mutation createTodo {
  createTodo(input:{text:"create todo", userId: "test1"}) {
    id
    user {
      id
      name
    }
    text
    done
  }
}
```

## read Todo
```
query Todos {
    todos {
        id
        text
        done
    }
}
```

## update Todo
```
mutation updateTodo {
  updateTodo(input:{text:"update todo", userId: "test1", todoId: "hoge"}) {
    id
    user {
      id
      name
    }
    text
    done
  }
}
```