# tinyrp



## Command

```
go mod tidy
```

### Run reverse Proxy
```shell
make run
```
### Run spawn sample service

```shell
make run_sample_a
make run_sample_b
```

### Check discovery
```shell
curl localhost:8000/discovery | jq
```