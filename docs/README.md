# Documentation

## Running


Just running golden will work, but you will get an error with each interaction as it 
trys to reach the triplestore to POST the RDF to.  Fine for just testing calls to golden

```bash
podman run -p 6789:6789 localhost/fcore/golden
```

Docker or Podman compose will bring up a companion triplestore 

```bash
podman-compose up  --privileged --group-add keep-groups -e GRANT_SUDO=yes
```

[Docker Hub Page for GoLDeN](https://hub.docker.com/repository/docker/fils/golden/general)

## Interactions with LDN

Some example interactions with LDN endpoints

```bash
curl -X POST  'http://localhost:6789/id/ldn/ID/inbox'  --data @./examples/testPackage.json
```



## Index

* [Example curl calls](interactions.md)
* [Thoughts on URIs](onURIs.md)