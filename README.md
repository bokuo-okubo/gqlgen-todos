

# sample queries

```graphql
mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}
```

```gql
query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}
```
