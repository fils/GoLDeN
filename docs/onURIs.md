# URI management

## Notes and thoughts (nothing real coherent)

A typical URI for an open data resource might look like:

http://opencoredata.org/pkg/id/8448c71edc22a06a26501a967223e5502dd4678be06c5761440167229ec9b715

So to make an inbox you might see something like:

http://opencoredata.org/pkg/id/RESID/inbox

However, I don't know if I want to maintain all those..  
I'd like not to have to instantiate them ahead of time, since the majority will never be used
and I don't need to maintain useless triples.  

So.. I'm faced with doing a check like

```go
func CheckForInbox(resid string) bool {
    // "if there is or isn't an inbox"
return true
}
```

then

``` go
func MakeNewInbox(resid string) (bool, error) {
// make a new inbox for a resource so we can store notifications
return true, nil
}
```

I'd likely maintain that as a separate triple database, but it could just be quads too.
Not sure which is best here.  

However, this raises some routing issues.  Either I have to mix my DO and LDN code and 
keep this one one code base, or I make the base URI different for an inbox which frees
me in the routing space to maintain that separately.

Something like
http://opencoredata.org/ldn/id/RESID/inbox

Wwhere I sub "pkg" with "ldn" and have a different URI.  Now, though, I want to make sure I maintain
association with this new data graph inside the object or package graph.   I could do that by adding a quad
to those graphs to point to this new resources (see also or something more explicate).

Or I can simply query across the two datasets.   If I needed only 1 named graph for all notifications
it could be a single named graph inside the object or package database.   However, then I have notifications in
two graphs.   If I maintain them in one graph, it's easier that way.   I just need to query across to see if
there is a notification.   I could also keep that resources triple for that relation in the notifiation graph (or
just push it to the object or package graph....   (as already stated above))

The inbox and inserted notification are new URIs..   they can maintain the old pattern if I can
resolve routing.    Then can stay in any graph (name graph or database) they wish.

Routing is typically done in traefik, then in the app runtime.

If LDN becomes a thing, people look for the inbox as an extension to the resource ID???  
Or as a property of the resource?  (In the later, then we could point to it)
Inspect pyldn and solid for how they do this....
