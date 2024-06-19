# ODIS Match Making

## About

This is the start of documentation on the use of this
package to support the OIDS "Matching Making"

## Simple Running

If you are running from this repo with Go directly use;

```bash
go run  cmd/reciever/main.go
```

This will run the receiver on port 6789.

You can test via curl with 

```bash
curl -v -X POST  'http://localhost:6789/id/ldn/ID/inbox'  --data @./examples/testPackage.json
```

> Note:  GoLDeN does not currently do any SHACL validation.  So it expects well-formed JSON-LD
> only.  It does not validate the data graph against something like a SHACL shape.  

If you have a running triplestore on the default 7878 port and GoLDeN can load the RDF 
into the triplestore, you will get a 201 CREATED return.

If you test without a triplestore, but the JSON-LD is valid,  GoLDeN
will return a status code 500 due to not being able to contact the server.

Sending a non-compliant package will result in an error 422 for an unprocessed entity

The return value will currently look like:

```bash
 http://mm.oceaninfohub.org/id/ldn/NAMEDGRAPH/inbox/1095553104
```

The trailing number is a unique hash that the user can use to check on their submission with.  
The structure of this URL will change as might the hashing approach.  This is provided for testing
only.  

## Errors

The standard error codes GoLDeN will return.

StatusCreated (201)

StatusBadRequest (400)

StatusUnprocessableEntity (422)

StatusInternalServerError (500)


## Docker

