# Javascript

## Thoughworks 
The Novo Nordisk tech radar was built on existing code from open source project [Thoughtworks](https://github.com/thoughtworks/build-your-own-radar/), and revised to fit the specifications of a decentralized tech radar.

## Library dependencies
The current tech radar requires specific Javascript libraries to function. To utilize each of these libraries, a Javascript library called requireJS binds them all together.

An example of how requireJS is used, with library dependencies defined and used in the tech radar script to fully render:

```
define([
  "d3",
  "d3tip",
  "d3-collection",
  "d3-selection",
  "chance",
  "lodash",
  "jquery",
  "jquery-autocomplete",
], function facModel(d3, d3tip, d3col, d3sel, Chance, _, $, AutoComplete) { ...
```

The defined libraries in the example above are linked in <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/requireConfig.js </mark>. The library and usage of requireJS in placed in <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/makeHtml.go</mark> after the closing body tag in the html template.


## Quadrant and Ring names
The quadrant and ring names of the tech radar is available to be changed in the javascript. You can currently change these names while searching for: js/graphing/config.js in the <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/renderingTechRadar.js </mark> file.

Example:
```
  const quadrantNames =
    '["Techniques", "Platforms", "Tools", "Languages & Frameworks"]';
  const ringNames = '["Adopt", "Trial", "Assess", "Hold"]';
```

## Thoughtworks one file JS
The original Thoughtworks tech radar is build on having a Docker container or a Webpack to start a server. In this version of the Tech Radar, these technologies have been removed. Furthermore, all Thoughtworks specific javascript files have been moved into one large js file called renderingTechRadar.js.

## Updating Thoughtworks radar
In the original Thoughtworks tech radar, the js files are split into their own specific "need"-case. In this version of the Tech Radar, all js files are combined into one large script file called <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/renderingTechRadar.js </mark>. Inside renderingTechRadar.js comments are inserted before each block of code, showcasing which Thoughtworks js file the code originated from. To update the Tech Radar properly, locate the commented code block you seek and update them from the Thoughtworks js file.

## Updating Thoughtworks radar with a library or script

To add a library or another script to the Tech Radar you can add it by using requireJS. First, add you script link or library link inside <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/requireConfig.js </mark>.

Example:
```
require.config({
    paths: {
      'YourScriptName': './your/path/to/js',
      ...  
    },
    ...
  }
});
```

then add the name you chose for your link when you define the scripts used in <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/renderingTechRadar.js </mark>.

Example:
```
define([
  "YourScriptName",
  ...
], function facModel(YourReferencedScriptName, ... ) { ...
```

## Code structure
Given that the javascript is in one file, ere