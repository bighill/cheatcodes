# curl cheatsheet #

Server authentication
```bash
curl -u username:password http://localhost
```

Send POST data
```bash
curl -d foo1=bar1 -d foo2=bar2 http://localhost
```

Display header in output
```bash
curl -i http://localhost
```

Set HTTP method
```bash
curl -X POST http://localhost
curl -X PUT http://localhost
curl -X DELETE http://localhost
```

___

Note to self: this is only helpful to an extent. At a certain point, better to just use the [Insomnia REST Client](https://insomnia.rest/)
