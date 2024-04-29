/**
 * Tech radar dependencies
 */
require.config({
    paths: {
      'd3': './HTML/js/libraries/d3',
      'd3tip': './HTML/js/libraries/d3-tip',
      'chance': './HTML/js/libraries/chance',
      'lodash': './HTML/js/libraries/lodash',
      'd3-collection': './HTML/js/libraries/d3-collection',
      'd3-selection': './HTML/js/libraries/d3-selection' ,
      'jquery': './HTML/js/libraries/jquery',
      'jquery-autocomplete': './HTML/js/libraries/jquery-ui'   
    },
    shim: {
      'd3tip': {
        deps: ['d3','d3-collection', 'd3-selection'],
        exports: 'd3.tip'
      },
      'jquery-autocomplete': {
        deps: ['jquery'],
        exports: '$.ui.autocomplete'
      }
  }
});