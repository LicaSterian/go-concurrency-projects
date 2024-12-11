# Go Concurrency Projects

### ubuntu-ssh
Spin-up a Ubuntu Docker container with these login credentials ```ubuntu:secret_password```
This container will be used to show the ```scp``` command.

### slice-search/linear and slice-search/concurrent
These projects mark the difference of searching linear and concurrently in a slice of 10 million entries.
The entry that it needs to find is the last one in the slice, highlighting the worst time complexity.

### cp/files 
This is a helper program that creates 1000 files that we will use in the ```cp/sequential``` and ```cp/concurrent``` projects.

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
We will use the ```requests/sequential``` and ```requests/concurrent``` projects to call the POST "users/new" endpoint.

### requests/sequential
Call the new user POST endpoint 10 times one request at a time.

### requests/concurrent
Call the new user POST endpoint 10 times concurrently.

Make Sequential and Concurrent requests to the ```http server``` 
