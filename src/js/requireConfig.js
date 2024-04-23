/**
 * Tech radar dependencies
 */
require.config({
    paths: {
      'd3': 'https://d3js.org/d3.v7.min',
      'd3tip': 'https://cdnjs.cloudflare.com/ajax/libs/d3-tip/0.9.1/d3-tip.min',
      'chance': 'https://cdnjs.cloudflare.com/ajax/libs/chance/1.1.11/chance.min',
      'lodash': 'https://cdnjs.cloudflare.com/ajax/libs/lodash.js/4.17.21/lodash.min',
      'd3-collection': 'https://cdn.jsdelivr.net/npm/d3-collection@1.0.7/dist/d3-collection.min',
      'd3-selection': 'https://cdn.jsdelivr.net/npm/d3-selection@3.0.0/dist/d3-selection.min' ,
      'jquery': 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min',
      'jquery-autocomplete': 'https://code.jquery.com/ui/1.13.1/jquery-ui.min'   
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