This folder contains examples of workflows that can be setup to automatically generate radars. This example uses a combination of dagger to create the radar and a github action yaml file that uploads it on github as an artifact. This can be extended to work with any other workflow system. The most basic version of this would be to create a cronjob that looks like so:
```
* * 1 * * /bin/bash git tag $RADAR_VER && git push --tags && /bin/update_radar_ver.sh
```
 This would result in a cronjob that on the first of every month generates a new radar file based on the name given to the $RADAR_VER name this could easily be updated using the bash script at the end of the command. Hypothetically one could even setup this as a kubernetes cronjob if you wanna get really crazy

github is the github action specific part and the CI folder contains the go file example of the dagger workflow


<!-- TODO: Change to novo main later when if they run the workflow -->
[Example of a generated Artifact](https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/actions/runs/8990468587)