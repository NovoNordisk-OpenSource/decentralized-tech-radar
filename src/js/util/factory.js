define([
  'd3',
  '../models/radar.js',
  '../models/quadrant.js',
  '../models/ring.js',
  '../models/blip.js',
  '../graphing/radar.js',
  '../config.js',
  '../graphing/config.js'
], function(d3, Radar, Quadrant, Ring, Blip, GraphingRadar, configFuntion, config) {
  
  const featureToggles = configFuntion().featureToggles;
  const { getGraphSize, graphConfig } = config;

  const plotRadar = function (blips, currentRadarName) {
    
    document.title = "plotRaderTitle"

    var rings = _.map(_.uniqBy(blips, 'ring'), 'ring')
    var ringMap = {}

    _.each(rings, function (ringName, i) {
      ringMap[ringName] = new Ring(ringName, i)
    })

    var quadrants = {}
    _.each(blips, function (blip) {
      if (!quadrants[blip.quadrant]) {
        quadrants[blip.quadrant] = new Quadrant(blip.quadrant[0].toUpperCase() + blip.quadrant.slice(1))
      }
      quadrants[blip.quadrant].add(
        new Blip(blip.name, ringMap[blip.ring], blip.isNew.toLowerCase() === 'true', blip.topic, blip.description),
      )
    })

    var radar = new Radar()
    _.each(quadrants, function (quadrant) {
      radar.addQuadrant(quadrant)
    })

    if (currentRadarName !== undefined || true) {
      radar.setCurrentSheet(currentRadarName)
    }

    const size = featureToggles.UIRefresh2022
      ? getGraphSize()
      : window.innerHeight - 133 < 620
      ? 620
      : window.innerHeight - 133
    new GraphingRadar(size, radar).init().plot()
  }

  function validateInputQuadrantOrRingName(allQuadrantsOrRings, quadrantOrRing) {
    const quadrantOrRingNames = Object.keys(allQuadrantsOrRings)
    const regexToFixLanguagesAndFrameworks = /(-|\s+)(and)(-|\s+)|\s*(&)\s*/g
    const formattedInputQuadrant = quadrantOrRing.toLowerCase().replace(regexToFixLanguagesAndFrameworks, ' & ')
    return quadrantOrRingNames.find((quadrantOrRing) => quadrantOrRing.toLowerCase() === formattedInputQuadrant)
  }

  const plotRadarGraph = function (blips, currentRadarName) {
    document.title = "Novo Nordisk Tech Radar"

    const ringMap = graphConfig.rings.reduce((allRings, ring, index) => {
      allRings[ring] = new Ring(ring, index)
      return allRings
    }, {})

    const quadrants = graphConfig.quadrants.reduce((allQuadrants, quadrant) => {
      allQuadrants[quadrant] = new Quadrant(quadrant)
      return allQuadrants
    }, {})

    blips.forEach((blip) => {
      const currentQuadrant = validateInputQuadrantOrRingName(quadrants, blip.quadrant)
      const ring = validateInputQuadrantOrRingName(ringMap, blip.ring)
      if (currentQuadrant && ring) {
        const blipObj = new Blip(
          blip.name,
          ringMap[ring],
          blip.isNew.toLowerCase() === 'true',
          blip.topic,
          blip.description,
        )
        quadrants[currentQuadrant].add(blipObj)
      }
    })

    const radar = new Radar()
    radar.addRings(Object.values(ringMap))

    _.each(quadrants, function (quadrant) {
      radar.addQuadrant(quadrant)
    })

    radar.setCurrentSheet(currentRadarName)

    const graphSize = window.innerHeight - 133 < 620 ? 620 : window.innerHeight - 133
    const size = featureToggles.UIRefresh2022 ? getGraphSize() : graphSize
    new GraphingRadar(size, radar).init().plot()
  }

  const CSVDocument = function (csvData) {
    var self = {}
    
    self.build = function () {
      csvfile = d3.csvParse(csvData)
      createBlips(csvfile)
    }

    var createBlips = function (data) {
      delete data.columns
      var blips = _.map(data)
      featureToggles.UIRefresh2022
        ? plotRadarGraph(blips, 'CSV File', [])
        : plotRadar(blips, 'CSV File', [])
    }

    return self
  }


  const Factory = function (sheetData) {
    var sheet
    sheet = CSVDocument(sheetData)
    return sheet
  }

  return Factory
});