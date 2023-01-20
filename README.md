# go_authentication
This is  a authentication microservice implemented in golang

# flow
    main package -> This package will init the app and set everything up
    server package -> This package will be implementing the server in this app, containing routes
    handler package -> This packages implement handlers for routes 

    request -> request handler -> server -> handler -> service -> response handler -> response