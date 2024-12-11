# Go Concurrency Projects

### slice-search/sequential

This project searches for a specific entry in a slice of 10 million entries by checking for a match one entry at a time.

### slice-search/concurrent

This project searches for a specific entry in a slice of 10 million entries concurrently. It check's how many CPU's your machine has, splits the slice into the number of CPU's you have.

### cp/files

This is a helper program that creates 1000 files that we will use in the `cp/sequential` and `cp/concurrent` projects.

### cp/sequential

This program copies 1000 files from one folder to another one at a time.

### cp/concurrent

This program copies 1000 files from one folder to another all at once

### http

gin-gonic http server.
It exposes the following routes that we can call:

- POST "users/new"
- GEt "users"
- GET "users/:userID"
  We will use the `requests/sequential` and `requests/concurrent` projects to call the POST "users/new" endpoint.

### requests/sequential

Call the new user POST endpoint 10 times one request at a time.

### requests/concurrent

Call the new user POST endpoint 10 times concurrently.

Make Sequential and Concurrent requests to the `http server`
