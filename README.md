# feeder

[![Go Report Card](https://goreportcard.com/badge/github.com/trandoshan-io/feeder)](https://goreportcard.com/report/github.com/trandoshan-io/feeder)

Feeder is a oneshot Go written program designed to initialize a Trandoshan instance.

## features

- use scalable messaging protocol (nats)

## how it work

- The Crawler process connect to a nats server (specified by env variable *NATS_URI*)
- Read system property *INITIAL_URI* and send this url to nats with subject *todoSubject*