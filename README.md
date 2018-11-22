# Config - A synchronised config map

A synchronised string map for holding and passing around a common config object.

## What?
A synchronised string map for holding and passing around a common config object. It was built and intended for use with [auth] but can be used as a standalone.

## Why?
This was part of a learning exercise to create [auth] which is a very rough Go equivalent of dotnet core Identity services.

## How?
See the tests for usage examples.

## Examples
See [examples] for a http/appengine implementations which uses config and [auth]. This is written for appengine standard 2nd gen, but also works as a standalone.

## Dependencies and services
None.

## Installation
Install using go get.

```sh
$ go get -u github.com/lidstromberg/config
```

   [auth]: <https://github.com/lidstromberg/auth>
