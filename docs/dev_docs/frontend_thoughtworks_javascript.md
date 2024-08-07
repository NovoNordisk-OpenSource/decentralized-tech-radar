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
const quadrantNames = '["Techniques", "Platforms", "Tools", "Languages & Frameworks"]';
const ringNames = '["Adopt", "Trial", "Assess", "Hold"]';
```

## Thoughtworks one file JS
The original Thoughtworks tech radar is build on having a Docker container or a Webpack to start a server. Furthermore, the original Thoughtworks js files are split into their own specific "need"-case. In this version of the Tech Radar, all js files are combined into one large script file called <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/renderingTechRadar.js </mark>.

## Updating Thoughtworks radar
Inside renderingTechRadar.js comments outline each block of code, showcasing which Thoughtworks js file the code originated from. To update the Tech Radar properly, locate the commented code block you seek and update them from the Thoughtworks js file.

## Updating Thoughtworks radar with a library or script
To add a library or another script to the Tech Radar you can add it by using requireJS. First, add your script link or library link inside <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/requireConfig.js </mark>.

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

## Updating Tech Radar Images
To update the images navigate to <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/js/renderingTechRadar.js </mark>. Search for "/HTML/images/" and replace the url with your image path.

Example image url path: 
```
.attr("src", "./src/HTML/images/existing.svg")
```

## Code structure
Given the one file of javascript, here is list of how the code is placed. 

References to files from Thoughtworks and their correspondent code lines:

```
1.  js/config.js                              (line 24-48) 
2.  js/graphing/config.js                     (line: 53-118)
3.  js/models/quadrant.js                     (line: 123-146)
4.  js/models/radar.js                        (line: 151-250)
5.  js/models/ring.js                         (line: 256-268)
6.  js/models/blip.js                         (line: 273-341)
7.  js/util/mathUtils.js                      (line: 346-348)
8.  js/util/htmlUtils.js                      (line: 353-364)
9.  js/util/stringUtil.js                     (line: 369-379)
10. js/graphing/components/quadrants.js       (line: 385-1128)
11. js/graphing/components/quadrantTables.js  (line: 1133-1405)
12. js/graphing/blips.js                      (line: 1410-1835)
13. js/util/ringCalculator.js                 (line: 1840-1865)
14. js/graphing/components/quadrantSubnav.js  (line: 1870-1986)
15. js/util/autoComplete.js                   (line: 1991-2063)
16. js/graphing/components/search.js          (line: 2068-2103)
17. js/graphing/components/buttons.js         (line: 2108-2116)
18. js/graphing/pdfPage.js                    (line: 2121-2134)
19. js/util/urlUtils.js                       (line: 2139-2166)
20. js/graphing/components/banner.js          (line: 2171-2202)
21. js/util/queryParamProcessor.js            (line: 2207-2221)
22. js/graphing/radar.js                      (line: 2226-3237)
23. js/util/factory.js                        (line: 3242-3374)
```