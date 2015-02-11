# turbo-octo-avenger
Web services in Go

Project status - design phase

# Architecture

- Load balancer node
- Web service nodes
- Database node


The Load balancer distributes the the incoming requests across the cluster of 
Web service nodes.  Web service nodes retieve shared state from the Database node.
Service discovery and configuration is shared across all nodes.

## Building blocks

The system consists of these major components:

- HAProxy - Load balancing
- Consul - Service discovery & configuration
- Go app - Web service application
- Postgres - Datastore

Also these components are used:

- Syslong - logging

## Non functionals

- Logging is recorded at the system level & then centralised for access, & ease
  of search
- Security is handled at the Web service layer
  
# Goals

- Provide a standardised REST API to hang new services off. The system standardises
  all the way that each service is built.
- Allow inter-process communication. Allow external systems and services to be called.
- Emphasise intra-process communication. Calling another service is as fast as a method call, and eliminates
  the ipc failure conditions of network down etc, but retains the same semantics as if called remotely, ie:
  has its own db connection and transaction etc.
- Explore the speed of such a system. How many requests can it respond to from a single server?
- Explore the scaleability of such a system. How many service components can be supported in a single process? 

# Resources

- http://www.alexedwards.net/blog/making-and-using-middleware
- http://www.gorillatoolkit.org/pkg/context
- https://justinas.org/writing-http-middleware-in-go/
- https://www.consul.io/
- http://syslog-ng.org/
- http://golang.org
- https://github.com/julienschmidt/go-http-routing-benchmark
- http://jmoiron.github.io/sqlx/
- https://github.com/xeipuuv/gojsonschema
# https://justinas.org/alice-painless-middleware-chaining-for-go/
# 