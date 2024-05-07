This folder contains examples of workflows that can be setup to automatically generate radars. This example uses a combination of dagger to create the radar and a github action yaml file that uploads it on github as an artifact. This can be extended to work with any other workflow system (Could maybe even be setup as a kubernetes cronjob if you wanna get really crazy)

github is the github action specific part and the CI folder contains the go file example of the dagger workflow

<!-- TODO: Change to novo main later when if they run the workflow -->
[Example of a generated Artifact](https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/actions/runs/8990468587)