## Rest
```
curl -X POST  -v localhost:8081/rest/v1/rest/graphql -H "Content-Type: application/json" -d '{"value":"received"}'

...

{"value":"received received"}
```

## GraphQL
```
curl -X POST -L -v -g localhost:8081/graphql -d '{
  restAndGraphQL(value: "testdata") {
    value
  }
}'

...

{"data":{"restAndGraphQL":{"value":"testdata received"}}}
```