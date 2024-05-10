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

The defined libraries in the example above are linked in <mark style="background-color: #69a8f5"> ./src/HTML/js/requireConfig.js </mark>. The library and usage of requireJS in placed in <mark style="background-color: #69a8f5"> ./src/HTML/makeHtml.go</mark> after the closing body tag in the html template.

## Updating Thoughtworks radar
Since Thoughtworks radar is build on having container or starting a server, which then make it hard to update  is depended on 