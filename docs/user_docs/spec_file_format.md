# Specification file format
The specification file (*specfile* for short) defines the data that will go into the tech radar.This is done by adding individual blips to a CSV file with the format defined in this file.

**A specfile has to have the following format:**

### Header
The header of the CSV specfile has to have *name*, *ring*, *quadrant*, *isNew*, *moved*, and *description*. The header of the file should look like this:
```
name,ring,quadrant,isNew,moved,description
```

### Content
The type for the content on each line is different depending on the column it is in. Here are the types for each column:
```  
- name        : string
- ring        : string
- quadrant    : string
- isNew       : bool
- move        : int
- description : string
```

Some of the columns are defined to take specific inputs. Columns can have these inputs:
- **name**: The name of a technique, platform, tool, or language or framework.
- **ring**: The name of the ring it's place being either *adopt*, *trial*, *assess* or *hold*.
- **quadrant**: The name of the quadrant being either *techniques*, *platforms*, *tools*, or *languages & frameworks*.
- **isNew**: *true* if the blip is new on the radar, or *false* if the blip isn't new on the radar.
- **move**: The integer representation of moving from one ring to an other. So moving from, for example, assess to trial would be a move of +1 and moving out from adopt to trial would be a move of -1. Since there are 4 rings the possible values for this column are [3,2,1,0,-1,-2,-3]
- **description**: A small description on the tech. This is an HTML string and as such HTML tags can be used to define the look of the description.

### An example of a CSV file:
```
name,ring,quadrant,isNew,moved,description
Python,hold,languages & frameworks,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.
golang,adopt,languages & frameworks,true,3,golang would be useful for...
WebAssembly,trial,tools,false,2,Lorem ipsum...
```