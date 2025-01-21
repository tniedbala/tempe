# Development Environment

## Docker
From the root of the repository execute the following to build a development image (see [Dockerfile](../Dockerfile)).
```shell
docker build --tag tempe-dev .
```

## Go
Commands for building and testing with Go:

1. ### Build
    Run ANTLR to generate Go source code in [../pkg/parser/base](./pkg/parser/base):
    ```shell
    docker run --rm -v $PWD:/workspace tempe-dev go build
    ```

2. ### Test
    Executes `go run test`:
    ```shell
    docker run --rm -v $PWD:/workspace tempe-dev go test 
    ```
    - Additional arguments following `go test` will be passed as cli arguments to `go run test`.

## ANTLR & Java
Commands for building and testing with Java:

1. ### ANTLR Test
    Generate java source code & jars to [../bin/grammar](./bin/grammar), and then test the generated files
    using the [ANTLR TestRig](https://github.com/mobileink/lab.clj.antlr/blob/master/doc/testrig.md):
    ```shell
    docker run --rm -v $PWD:/workspace tempe-dev antlr test
    ```

2. ### ANTLR Copy
    Copies the ANTLR `.jar` file from the container to the [../bin](./bin) directory. This can be helpful 
    if you'd like to run ANTLR from your local environment, which is necessary for executing the TestRig
    with the `-gui` option:
    ```shell
    docker run --rm -v $PWD:/workspace tempe-dev test go
    ```

3. ### Java Build
    Generate java source code & jars to [../bin/grammar](./bin/grammar):
    ```shell
    docker run --rm -v $PWD:/workspace tempe-dev java build
    ```