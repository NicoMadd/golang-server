# Golang HTTP Server

## Motive

This is a simple HTTP server implementation in Golang. It handles the TCP connection, parses the request and delivers it to a pool of workers. A worker should handle the incoming request as the developer desires as is intended by using the _net/http_ module that Go already provides. This implementation doesn't use that module as it would make everything easier.

The project, though basic and not trying to be the most performant available, is an attempt to learn and improve golang usage, testing and its heuristic on design usage.

## How to run it

### Requirements

Go : ^1.22.x

### Commands

Execute _go run ._

