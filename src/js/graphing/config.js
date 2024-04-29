define(function() {
  const quadrantSize = 512
  const quadrantGap = 32

  const quadrantNames = '["Language", "Infrastructure", "Datastore", "Data management"]';
  const ringNames = '["Adopt", "Trial", "Assess", "Hold"]';

  const getQuadrants = () => {
    return JSON.parse(quadrantNames)
  }

  const getRings = () => {
    return JSON.parse(ringNames)
  }

  const isBetween = (number, startNumber, endNumber) => {
    return startNumber <= number && number <= endNumber
  }
  const isValidConfig = () => {
    return getQuadrants().length === 4 && isBetween(getRings().length, 1, 4)
  }

  const graphConfig = {
    effectiveQuadrantHeight: quadrantSize + quadrantGap / 2,
    effectiveQuadrantWidth: quadrantSize + quadrantGap / 2,
    quadrantHeight: quadrantSize,
    quadrantWidth: quadrantSize,
    quadrantsGap: quadrantGap,
    minBlipWidth: 12,
    blipWidth: 22,
    groupBlipHeight: 24,
    newGroupBlipWidth: 88,
    existingGroupBlipWidth: 124,
    rings: getRings(),
    quadrants: getQuadrants(),
    groupBlipAngles: [30, 35, 60, 80],
    maxBlipsInRings: [8, 22, 17, 18],
  }

  const uiConfig = {
    subnavHeight: 60,
    bannerHeight: 200,
    tabletBannerHeight: 300,
    headerHeight: 80,
    legendsHeight: 42,
    tabletViewWidth: 1280,
    mobileViewWidth: 768,
  }

  function getScale() {
    return window.innerWidth < 1800 ? 1.25 : 1.5
  }

  function getGraphSize() {
    return graphConfig.effectiveQuadrantHeight + graphConfig.effectiveQuadrantWidth
  }

  function getScaledQuadrantWidth(scale) {
    return graphConfig.quadrantWidth * scale
  }

  function getScaledQuadrantHeightWithGap(scale) {
    return (graphConfig.quadrantHeight + graphConfig.quadrantsGap) * scale
  }

  return {
    graphConfig: graphConfig,
    uiConfig: uiConfig,
    getScale: getScale,
    getGraphSize: getGraphSize,
    getScaledQuadrantWidth: getScaledQuadrantWidth,
    getScaledQuadrantHeightWithGap: getScaledQuadrantHeightWithGap,
    isValidConfig: isValidConfig
  }
})