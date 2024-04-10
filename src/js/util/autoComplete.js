/*const $ = require('jquery')
require('jquery-ui/ui/widgets/autocomplete')

const config = require('../config')
const featureToggles = config().featureToggles*/
/*
define([
  '../config',
  'jquery',
  'jquery-autocomplete'
], function(config, $, autocomplete) {
  const featureToggles = config().featureToggles;

  const createAutoComplete = (el, blips, onSelect) => {
    const input = document.querySelector(el);
    const resultList = document.createElement('ul');
    resultList.className = 'autocomplete-results';
    input.parentNode.appendChild(resultList);

    input.addEventListener('input', () => {
      const searchTerm = input.value.toLowerCase();
      const matches = blips.filter(({ blip }) => {
        const searchable = `${blip.name()} ${blip.description()}`.toLowerCase();
        return searchTerm.split(' ').every(term => searchable.includes(term));
      });

      // Clear previous results
      resultList.innerHTML = '';

      // Add new results to the list
      matches.forEach(match => {
        const li = document.createElement('li');
        li.className = 'autocomplete-result';
        li.textContent = match.blip.name();
        li.addEventListener('click', () => {
          onSelect(null, { item: match });
        });
        resultList.appendChild(li);
      });
    });

    document.addEventListener('click', (event) => {
      if (!input.contains(event.target)) {
        resultList.innerHTML = ''; // Clear results when clicking outside
      }
    });
  };

  const AutoComplete = (el, quadrants, onSelect) => {
    const blips = quadrants.reduce((acc, quadrant) => {
      return [...acc, ...quadrant.quadrant.blips().map(blip => ({ blip, quadrant }))];
    }, []);

    if (featureToggles.UIRefresh2022) {
      createAutoComplete(el, blips, onSelect);
    } else {
      // Define alternative behavior if the toggle is off
      createAutoComplete(el, blips, onSelect);
    }
  };

  return AutoComplete;
});
*/

// const $ = require('jquery')
// require('jquery-ui/ui/widgets/autocomplete')

// const config = require('../config')
// const featureToggles = config().featureToggles

// $.widget('custom.radarcomplete', $.ui.autocomplete, {


define([
  'jquery',
  'jquery-autocomplete',
  '../config.js',
], function( $, autocomplete, configFuntion) {
  
  const featureToggles = configFuntion().featureToggles;
  
  $.widget('custom.radarcomplete', $.ui.autocomplete, {
    _create: function () {
      this._super()
      this.widget().menu('option', 'items', '> :not(.ui-autocomplete-quadrant)')
    },
    _renderMenu: function (ul, items) {
      let currentQuadrant = ''

      items.forEach((item) => {
        const quadrantName = item.quadrant.quadrant.name()
        if (quadrantName !== currentQuadrant) {
          ul.append(`<li class='ui-autocomplete-quadrant'>${quadrantName}</li>`)
          currentQuadrant = quadrantName
        }
        const li = this._renderItemData(ul, item)
        if (quadrantName) {
          li.attr('aria-label', `${quadrantName}:${item.value}`)
        }
      })
    },
  })

  const AutoComplete = (el, quadrants, cb) => {
    const blips = quadrants.reduce((acc, quadrant) => {
      return [...acc, ...quadrant.quadrant.blips().map((blip) => ({ blip, quadrant }))]
    }, [])

    if (featureToggles.UIRefresh2022) {
      $(el).autocomplete({
        appendTo: '.search-container',
        source: (request, response) => {
          const matches = blips.filter(({ blip }) => {
            const searchable = `${blip.name()} ${blip.description()}`.toLowerCase()
            return request.term.split(' ').every((term) => searchable.includes(term.toLowerCase()))
          })
          response(matches.map((item) => ({ ...item, value: item.blip.name() })))
        },
        select: cb.bind({}),
      })
    } else {
      $(el).radarcomplete({
        source: (request, response) => {
          const matches = blips.filter(({ blip }) => {
            const searchable = `${blip.name()} ${blip.description()}`.toLowerCase()
            return request.term.split(' ').every((term) => searchable.includes(term.toLowerCase()))
          })
          response(matches.map((item) => ({ ...item, value: item.blip.name() })))
        },
        select: cb.bind({}),
      })
    }
  }

  return AutoComplete;
});
