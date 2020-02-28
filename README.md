# Logtest Service

This is the Logtest service

Generated with

```
micro new gitlab.com/krayohu/logtest --namespace=ps.clientele --fqdn=ps.clientele.svc.logtest --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: ps.clientele.svc.logtest
- Type: srv
- Alias: logtest

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./logtest-srv
```

Build a docker image
```
make docker
```