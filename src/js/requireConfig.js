/**
 * Tech radar dependencies
 */
require.config({
    paths: {
      'd3': './js/libraries/d3',
      'd3tip': './js/libraries/d3-tip',
      'chance': './js/libraries/chance',
      'lodash': './js/libraries/lodash',
      'd3-collection': './js/libraries/d3-collection',
      'd3-selection': './js/libraries/d3-selection' ,
      'jquery': './js/libraries/jquery',
      'jquery-autocomplete': './js/libraries/jquery-ui'   
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