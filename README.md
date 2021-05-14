# 5d7fc6c5-a614-47ec-92d0-01b1b09f54f5

## Description

DataImpact backend developer test assignment.

## How to run

### Deploy the stack

```shell
make build
make up
```

### Minimal tests

Import the [Postman collection](postman/minimal_test_collection.json) to Postman and run it. 

**Important:** Make sure the file in `Create users` request point to [sample_data.json](postman/sample_data.json).

If you encounter a 401 error saying the JWT token is expired, rerun one of the login requests, since the token is valid 
for 5 minutes only.

## What have not been done

* Update user: I did not have time to. Either I would have done a full PUT or using PATCH with 
  [JSONPatch](https://datatracker.ietf.org/doc/html/rfc6902).
  
* Concurrent creation: It would require some atomic write at MongoDB and file system level, which I did not have time 
  dig into.
  
* Unit tests: It's a bad practice but again, time constraint. 
  
## What can be improved

### Code

* Unit tests and functional tests are to be added.
* Better error handling. Right now only green paths are correctly handled, many edge cases leads to error 500 where 
  they should not.
* Using `[]byte` instead of `string` for sensitive data.
* Bcrypt seems to have a max length constraint, so the password cannot be combined with the user ID (very long) to form 
  a better credentials.
* The app can be a three-tier architecture, but I merged the HTTP handler with the logic controller for the sake of 
  simplicity.
* Deletion should be soft-delete instead of hard-delete.
  
### Architecture

* Creation step should be asynchronous, since the input file can be large. A better way to do it might be to 
  temporarily save the file, push a message into a queue and let another process consume the file afterward.