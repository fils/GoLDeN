package graph

const queries = `
# Comments are ignored, except those tagging a query.

# tag: test
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
select *
where {
  { ?s ?p ?o } 
  UNION { GRAPH ?g { ?s ?p ?o } } 
}
LIMIT 10


# tag: orgsearch
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
select DISTINCT  ?name ?url ?desc
where {
  ?s rdf:type schema:Organization .
  ?s schema:description ?desc .
  ?s schema:name ?name .
  ?s schema:url ?url .
  ?desc bds:search "{{.}}" .
}


# tag: describeCall
DESCRIBE <{{.}}>

# tag: detailsCall
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
select distinct ?s ?aname ?name ?url ?description ?citation ?datepublished ?curl  ?keywords ?license
where {
VALUES ?url { <{{.}}> }.
?s rdf:type schema:Dataset .
?s <http://www.w3.org/2000/01/rdf-schema#seeAlso> ?url
OPTIONAL { ?s schema:alternateName ?aname } .
OPTIONAL { ?s schema:citation      ?citation }
OPTIONAL { ?s schema:datePublished ?datepublished }
OPTIONAL { ?s schema:description   ?description }
OPTIONAL { ?s schema:distribution ?distribution }
OPTIONAL { ?s schema:distribution ?distribution .
          ?distribution schema:contentUrl ?curl .
        }
OPTIONAL { ?s schema:identifier ?identifier }
OPTIONAL { ?s schema:keywords ?keywords }
OPTIONAL { ?s schema:license       ?license}
OPTIONAL { ?s schema:name         ?name}
OPTIONAL { ?s schema:url ?url}
OPTIONAL { ?s schema:measurementTechnique ?measurementtechnique }
}



# tag: LogoCall
PREFIX schema: <http://schema.org/>
SELECT DISTINCT ?graph ?type ?resource ?logo
WHERE {
  VALUES ?resource {
    <{{.RESID}}>
  }
  VALUES ?type {
    schema:Organization
    schema:Dataset
  }
  GRAPH ?graph {
    ?resource rdf:type ?type .
    OPTIONAL { ?resource schema:logo [ schema:url ?logo ] }

  }
}
ORDER BY ?graph ?type ?resource ?logo


# tag: ResourceResults
prefix schema: <http://schema.org/>
SELECT ?val ?desc ?pubname ?puburl
WHERE
{
  BIND(<{{.RESID}}> AS ?ID)
  ?ID schema:publisher ?pub .
  ?pub schema:name ?pubname .
  ?pub schema:url ?puburl .
  ?ID schema:variableMeasured ?res  .
  ?res a schema:PropertyValue .
  ?res schema:value ?val   .
  ?res schema:description ?desc
}

# tag: ResourceSetIGSN
prefix schema: <http://schema.org/>
SELECT DISTINCT ?igsn
WHERE
{
VALUES ?ID
{  <http://get.iedadata.org/doi/100415> <http://get.iedadata.org/doi/100416>
}
 ?dataset a schema:Dataset .
 ?dataset rdfs:seeAlso ?ID .
 ?dataset schema:hasPart ?sample .
 ?sample schema:additionalType "http://schema.geolink.org/1.0/base/main#PhysicalSample" .
 ?sample schema:identifier ?id .
 ?id schema:propertyID "IGSN" .
 ?id schema:value ?igsn .
}



# tag: ResourceSetResults
prefix schema: <http://schema.org/>
SELECT DISTINCT ?val ?desc ?pubname ?puburl
WHERE
{
VALUES ?ID
{  {{.}}
}
?ID schema:variableMeasured ?res .
OPTIONAL {
?res a schema:PropertyValue .
?res schema:value ?val .
?res schema:description ?desc
}
OPTIONAL {
?ID schema:publisher ?pub .
OPTIONAL { ?pub schema:name ?pubname }
OPTIONAL { ?pub schema:url ?puburl }
}
}

# tag: ResourceSetPeople
PREFIX schema: <http://schema.org/>
SELECT DISTINCT ?g ?person ?person_role (IF(?role = schema:contributor, "Contributor", IF(?role = schema:creator, "Creator/Author", "Author")) as ?rolename) ?name ?url ?orcid
WHERE
{
 GRAPH ?g {
  VALUES ?dataset
  {   {{.}}
  }
  VALUES ?role
  {
   schema:author
   schema:creator
   schema:contributor
  }
  VALUES ?orcid_type {
    "datacite:orcid"
    "http://purl.org/spar/datacite/orcid"
  }
  { ?dataset ?role ?person_role }
  OPTIONAL {
    {
     ?person_role a schema:Role .
     ?person_role ?role ?person .
     OPTIONAL { ?person schema:name ?name }
     OPTIONAL { ?person schema:url ?url }
     OPTIONAL {
       ?person schema:identifier ?id .
       ?id schema:propertyID ?orcid_type .
       ?id schema:value ?orcid
     }
    } UNION {
     ?person_role a schema:Person .
     OPTIONAL { ?person_role schema:name ?name }
     OPTIONAL { ?person_role schema:url ?url }
     OPTIONAL {
       ?person_role schema:identifier ?id .
       ?id schema:propertyID ?orcid_type .
       ?id schema:value ?orcid
     }
    }
  }
 }
}


`
