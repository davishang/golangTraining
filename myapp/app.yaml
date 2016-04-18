application: trusty-dialect-127419    //unique application ID
version: 1
runtime: go         //this code runs in the go runtime envrinment, with API version go1
api_version: go1

handlers:
- url: /.*          //every request to a URL whose path matches the regular expression /.* (all URLs) should be handles by the Go program
  script: _go_app   //this value is a magic string recognized by the development web server; it is ignore by the production App Engine servers.



