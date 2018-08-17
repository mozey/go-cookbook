#!/usr/bin/env bash

export APP_FOO=bar

go build \
-ldflags "-X main.foo=${APP_FOO}
-X github.com/mozey/go-cookbook/ldflags/config.Foo=${APP_FOO}
-X github.com/mozey/go-cookbook/ldflags/config.fiz=${APP_FOO}" \
-o ./main.out ./
