define(function() {
  function renderButtons(radarFooter) {
    const buttonsRow = radarFooter.append('div').classed('buttons', true)

    buttonsRow
      .append('button')
      .classed('buttons__wave-btn', true)
      .text('Print this Radar')
      .on('click', window.print.bind(window))
  }

  return renderButtons;
});
