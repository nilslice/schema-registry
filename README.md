# Schema Registry

The registry is focused on Protocol Buffer definitions for data and services to 
continuously validate de/serialization compatibility amongst its users. The 
goal is to provide a runtime compatibility check, much like a "health check" is
made by interdependent services.

## Process

**1. Registration / Initialization**

Teams must register their `git` repo containing their schema. A `proto.lock` 
file must also be present within the root of the repo, or otherwise denoted in 
the configuration object in the `Init` request.

Registering a schema will result in a response from the registry containing a 
`token`, for the registrant to update their schema in the future, as well as a 
`Schema` object, containing (roughly) the URL locating the schema in storage, 
the lock file in byte form, the origin git repo location, and any metadata (?)
that must be kept along with the schema in key/value format.

**2. Retrieval**

In order to validate against an existing schema, the registered schema must be
retrieved by a peer system. With the schema URL, a request can be made by using
the `Get` rpc, which will return a `Schema` object for status checks and further
inspection. 

**3. Status / Compatibility Checking**

Just as a service should do health-checks prior to making cross-boundry requests,
a compatibility check should be made any time data is transmitted across service
boundaries. This includes service-to-service calls, data drops (into S3 or elsewhere),
data downloads (from S3 or elsewhere), etc. 

With a retrieved `Schema` object, users can call the `Status` rpc, which takes a
`Schema` and a `proto.lock` file which would have been stored in-memory or along
side the service executable. The `proto.lock` represents the Protocol Buffer 
definition (contract) known by the user system at the time the system was built
& deployed. The `Schema` object is the live, registered schema by another peer
system being interacted with from the user's system.

A `Status` call will return an array of warnings, which when non-zero in length,
indicate that a breaking, incompatible change was made to a schema, and that the
user's service making the status check should not proceed with its request since
the de/serialization of their data payload will fail on one end of the contract.

**4. Commit / Update Schema Changes**

With the originating `token` issued at `Init`-time, and the updated repo URL
containing the Protocol Buffer contract, a user can call the `Commit` rpc and
update their schema. The registry will first check the schema located in the 
`git` repo, and compare it to the previous version's `proto.lock` file to check
for any breaking or incompatible changes. A `force` boolean option can be passed
to indicate that despite breaking changes, the schema should be updated.

The call to `Commit` will return an array of warnings, similar that of `Status` 
and the user can adjust their schema to fix these warnings. If the user passed 
a truthy `force` option, the warnings will still be returned, but the schema in
the return response will have been updated in the registry.

## Components

1. Database
    - activity / access logs

2. Cloud Storage
    - S3/GCS for versioned file store
    - Bucket URLs are defacto schema locators

3. Compatibility Engine
    - [`protolock`](https://github.com/nilslice/protolock) libraries will drive Schema checks
    


