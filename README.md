# Galaxy

Description: the goal of this api is to provide the ability to triangulate the position of a distress signal


## Assumptions:

### Position
If there are two valid positions, then the possible points of intersection are calculated.

![twoPoints](https://i.ibb.co/wQnrfPN/Screen-Shot-2022-05-23-at-06-09-15.png)

If there are three valid points, the last one will be used to break the tie.

![threePoints](https://i.ibb.co/m81hzSs/Screen-Shot-2022-05-23-at-06-12-21.png)

## Running the project

**Requirements:** golang 1.18

**Build:**

```shell
    go build main/main.go
```

**Execute (build included):**

```shell
    go run main.go
```

**Run tests:**

Using the native go command:

```shell
    go test ./...
```

Alternatively, using the included formatter and test runner (recommended):

```shell
    ./test.sh
```

**Test coverage:**

The test runner includes the package code coverage. To get a detailed code report, run the following runner to generate an html view of it 

```shell
    ./test_coverage.sh
```

**Deployment:**

To be able to deploy it is necessary to be authenticated against the cluster and the docker registry.

```shell
    ./deploy.sh
```

## Live API

The api can be tested on the following domain

```
    https://api.galaxy.fast.ar
```

## CURL example

```
curl --location --request POST 'https://api.galaxy.fast.ar/v1/topsecret' \
--header 'Content-Type: application/json' \
--data-raw '{
    "satellites": [
        {
            "name":"kenobi",
            "distance": 200.0,
            "message": ["este","","","mensaje",""]
        },
         {
            "name":"skywalker",
            "distance": 500,
            "message": ["","es","","","secreto"]
        },
         {
            "name":"sato",
            "distance": 904.3973,
            "message": ["este","","un","",""]
        }
    ]
}'
```