# GoLDeN  
## Linked Data Notifications for Data Sets

### About

GoLDeN (Go Linked Data Notifications ....   I just made up the e) is a Go based implementation of the W3C Linked
Data Notifications recommendation.  This is a work in progress and currently only the receiver is has been fleshed
out and run through the validation.  The sender and consumer code is just place holder stuff till I get a chance to 
work on it.   Special thanks to https://github.com/albertmeronyo/pyldn which allowed me to bootstrap this work far 
quicker.  Note this work is focused on LDN for data sets in particular schema.org/Dataset and DCAT Dataset. 

### Goal

This work is related to a larger body of work with BCO-DMO related to the application of digital object architecture
as defined by RDA work (citations needed here obviously).   The goal is to apply the LDN pattern to data sets to allow
the publication and subscription to inboxes/notifications (as defined by the LDN recommendation) associated with data sets.

### Details

The exact implementation details of how we are approaching needs to be detailed out a bit more here.  Connecting LDN to the 
unique IDs of datasets is important to provide the datasets an identity that can be associated with an LDN inbox.  Also we don't 
actually generate the inbox until the first notification is sent, but yet need to make sure to publish the capacity on all resources.
The goal is to not build huge amounts of triples that then don't ever get used.

So, I'll attempt to get a document published that details out how we are approaching this goal and adapting the LDN pattern
to data sets. 


