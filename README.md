# Berlioz GoLang SDK

## Using

## Development

### Running locally

```
$ berlioz local push-run
$ docker ps -a | grep berliozgo-example | awk '{print $1}' | xargs docker logs
```

Further launches can be made much quicker using:
```
berlioz local push-run --quick --cluster berliozgo --service example
```

berlioz local push-run --quick --cluster berliozgo --service example; docker ps -a | grep berliozgo-example | awk '{print $1}' | xargs docker logs

**Helper script to run and fetch output**
```
$ ./test-run.sh
```

cleanup:
```
berlioz local stop --cluster berliozgo
```