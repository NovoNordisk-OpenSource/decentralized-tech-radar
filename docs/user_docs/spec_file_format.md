# Spec file format
**Will be rewritten** 
This product uses CSV for file reading, this means that group definition is separated by a  ",". The spec for this product will have the following group definitions: name, ring quadrant, isNew, moved and description. For the CSV, this will look like:
```
name,ring,quadrant,isNew,moved,description
```
The reader will then translate these to structs which get translated to types using unmarshalling.

The types foreach group:  
```  
- name        : string
- ring        : string
- quadrant    : string
- isNew       : bool
- move       : int
- description : string
```

Some of the groups is defined to take specific inputs. Groups can have these inputs:
```
- name        : The name of a data management, datastore, infrastructure or language.
- ring        : The name of the ring it's place like adopt, trial, assess or hold.
- quadrant    : The name of the quadrant like data management, datastore, infrastructure or language.
- isNew       : True or false if the tech is new or not for the tech radar.
- move        : The integer representation of moving from one ring to an other. So moving from, for example, assess to trial would be a move of +1 and moving out from adopt to trial would be a move of -1.
- description : A small description on the tech.
```

A example of a CSV file:
```
name,ring,quadrant,isNew,moved,description
Python,hold,language,false,0,Lorem ipsum dolor sit amet, consectetur adipiscing elit.
golang,adopt,Language,true,3,golang would be useful for...
WebAssembly,trial,infrastructure,false,2,Lorem ipsum...
```