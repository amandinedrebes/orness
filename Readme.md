REST API server
============================

The server will store and return notes. It will work with the JSON representation of the notes

Running api server:
* go run main.go [-ip <ip>] [-port <port>]

IP is an optional argument: 127.0.0.1 means localhost only (default value: 127.0.0.1)
PORT is an optional argument (default value: 4000)


Running test compains:
* cd internal/core
* go run test

Bonus:
*  swagger documentation available (in webbrowser): http://127.0.0.1:4000/swaggerui/

Information:
This exercise can be easily solved by a simple main.go file as the scope is very limited.
It is a real position to start splitting code in order to separate API from the backend in order to prepare a real database integration.
The concurrency subject is solved by adding a Mutex.