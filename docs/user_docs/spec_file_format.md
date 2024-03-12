# Spec file format
**Will be rewritten** 
This product uses CSV for file reading, this means that group definition is separated by a  ",". The spec for this product will have the following group definitions: name, ring quadrant, isNew, moved and description. For the CSV, this will look like:
```
name,ring,quadrant,isNew,moved,description
```
The reader will then translate these to structs, and therefor these groups will also have types.

The types foreach group:  
```  
- name        : string
- ring        : string
- quadrant    : string
- isNew       : bool
- moved       : int
- description : string
```

Some of the groups is defined to take specific inputs. Groups can have these input
```
- name        : The name of a data management, datastore, infrastructure or langauge.
- ring        : The name of the ring it's place like adopt, trial, assess or hold.
- quadrant    : The name of the quadrant like data management, datastore, infrastructure or langauge.
- isNew       : True or false if the tech is new or not for the tech radar.
- moved       : 0-3 depending on what ring the tech moved to. so going from hold to adopt is going from 0 to 3 is moved=3.
- description : A small description on the tech.
```

A example of a CSV file:
```
name,ring,quadrant,isNew,moved,description
Python,hold,language,false,0,Lorem ipsum dolor sit amet, consectetur adipiscing elit.
golang,adopt,Language,ture,3,golang would be usefull for...
WebAssembly,trial,infrastucture,false,2,Lorem ipsum...
```