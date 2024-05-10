# CSS

The CSS file changes the styling of the Tech Radar. The path to the Tech Radar CSS can be found in <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/css/style.css </mark>

## Image paths in CSS
To change the image paths in the css navigate to <mark style="background-color: #69a8f5; margin:0 4px; padding: 0 4px"> ./src/HTML/css/style.css </mark>.

Locate the image paths by searching for, for instance, "HTML/images/", followed by changing the link path to yours.

Example:
```
#alternative-buttons .search-radar {
    background-color: inherit;
    background-image: url(../src/HTML/images/search-logo-2x.svg);
    background-position: 10px;
    background-repeat: no-repeat;
    border: 1px solid #aaa
}
```