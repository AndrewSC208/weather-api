# Weather Forecast Api
Accepts latitude and longitude coordinates.
Returns the short forecast for that area for Today (“Partly Cloudy” etc).
Returns a characterization of whether the temperature is “hot”, “cold”, or “moderate” (use your discretion on mapping temperatures to each type).
Use the National Weather Service API Web Service as a data source.

## Run
To run the service `cd ./services/weather/forecast && go run main.go`. If your having trouble running the service make sure that the `config.yaml` file is in the same directory when running. This is also the case when running as a compiled binary. Once the service is running get the forecast for the coordinates using the below as an example.

```shell
curl -X POST --header "Content-Type: application/json" \
    --data '{"latitude": 39.7456, "longitude": -97.0892}' \
    http://localhost:9090/weather.forecast.v1.ForecastService/ShortForecast
```

## Notes
A few notes. [Connect RPC](https://connectrpc.com/docs/go/getting-started/) and [Draft](https://github.com/steady-bytes/draft) were used to build the service. `Draft` is an open-source
micro-services framework that I personally maintain for building reliable, real-time distributed systems that scale, and thought I this would be a good time to show it off. Additional features like `pub/sub`, `service discovery`, and a `control plane` are built into the framework.

Additionally, `ConnectRPC` is an implementation of `gRPC` using [Protocol Buffers](https://protobuf.dev/) an  `IDL` created by google. It's used to define an api, and generate client/server code to a target programming language (in this case golang). The service is defined in `api/weather/forecast/v1/service.proto`. Protocol buffers abstract a lot of the boilerplate code that is needed when interacting with other web services. `gRPC` also has a few opinions on how the protocol should be implemented and thus an http `POST` is used with a message body instead of a `GET` with url parameters following the rest standard.

Finally, the nesting of the service code is an opinion of `Draft` (below layout as an example).
```shell
.
├── api
│   ├── buf.gen.gotag.yaml
│   ├── buf.gen.go.yaml
│   ├── buf.lock
│   ├── buf.yaml
│   ├── go.mod
│   ├── go.sum
│   └── weather
│       └── forecast
│           └── v1
│               ├── service.pb.go
│               ├── service.pb.validate.go
│               ├── service.proto
│               └── v1connect
│                   └── service.connect.go
├── draft.yaml
├── README.md
└── services
    └── weather
        └── forecast
            ├── config.yaml
            ├── controller
            │   └── weatherController.go
            ├── go.mod
            ├── go.sum
            └── main.go
```

This is because a mono repo is used to maintain multiple micro-services coordinated around a specific business domain following the principles of `Domain Driven Design`. In this case, the domain would be `weather` and the service `forecast`. As you can imagine if a business had many different business verticals, domains would be established around those verticals and services would be established within the domain to implement required business logic of that domain. This gives domains a loose coupling in the system as a whole with clear definition of what domain owns what resources (i.e data, events, and or api's).

## Build
To build the service binary run from the project root `cd ./services/weather/forecast && go build -o forecast main.go`.

## System Requirements
Golang `1.22.5` installed on your machine