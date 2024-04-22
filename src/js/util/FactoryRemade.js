define([
    'd3',
    'lodash',
    'jquery',
    'jquery-autocomplete',
    '../graphing/radar.js',
    'd3',
    'd3tip',
    'd3-collection',
    'd3-selection'
  ], function(d3, _, $, autocomplete, GraphingRadar, d3, d3tip) {
    

    // ROOT JS:                     js/config.js 

    const mainConfig = () => {
        const env = {
            production: {
            featureToggles: {
                UIRefresh2022: true,
            },
            },
            development: {
            featureToggles: {
                UIRefresh2022: true,
            },
            },
        }

        environ = window.APP_CONFIG = {
            ENVIRONMENT: 'production',
            featureToggles: {
            UIRefresh2022: true
            }
        };

        return window.APP_CONFIG && window.APP_CONFIG.ENVIRONMENT
        ? env[window.APP_CONFIG.ENVIRONMENT]
        : env.development;
    }










    // MODELS:                      js/graphing/config.js

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


  // MODELS:                      js/models/quadrant.js'






  // MODELS:                      js/models/radar.js'

    const Radar = function () {
        const featureToggles = mainConfig().featureToggles
    
        let self, quadrants, blipNumber, addingQuadrant, alternatives, currentSheetName, rings
    
        blipNumber = 0
        addingQuadrant = 0
        quadrants = featureToggles.UIRefresh2022
          ? [
              { order: 'first', startAngle: 0 },
              { order: 'second', startAngle: -90 },
              { order: 'third', startAngle: 90 },
              { order: 'fourth', startAngle: -180 },
            ]
          : [
              { order: 'first', startAngle: 90 },
              { order: 'second', startAngle: 0 },
              { order: 'third', startAngle: -90 },
              { order: 'fourth', startAngle: -180 },
            ]
        alternatives = []
        currentSheetName = ''
        self = {}
        rings = {}
    
        function setNumbers(blips) {
          blips.forEach(function (blip) {
            ++blipNumber
            blip.setBlipText(blipNumber)
            blip.setId(blipNumber)
          })
        }
    
        self.addAlternative = function (sheetName) {
          alternatives.push(sheetName)
        }
    
        self.getAlternatives = function () {
          return alternatives
        }
    
        self.setCurrentSheet = function (sheetName) {
          currentSheetName = sheetName
        }
    
        self.getCurrentSheet = function () {
          return currentSheetName
        }
    
        self.addQuadrant = function (quadrant) {
          
          quadrants[addingQuadrant].quadrant = quadrant
          setNumbers(quadrant.blips())
          addingQuadrant++
        }
        self.addRings = function (allRings) {
          rings = allRings
        }
    
        function allQuadrants() {
          
    
          return _.map(quadrants, 'quadrant')
        }
    
        function allBlips() {
          return allQuadrants().reduce(function (blips, quadrant) {
            return blips.concat(quadrant.blips())
          }, [])
        }
    
        self.rings = function () {
          if (featureToggles.UIRefresh2022) {
            return rings
          }
    
          return _.sortBy(
            _.map(
              _.uniqBy(allBlips(), function (blip) {
                return blip.ring().name()
              }),
              function (blip) {
                return blip.ring()
              },
            ),
            function (ring) {
              return ring.order()
            },
          )
        }
    
        self.quadrants = function () {
          return quadrants
        }
    
        return self
    }










    
    // THIS IS CONFIG FROM js/models/quadrant.js'
    

    const Quadrant = function (name) {
        var self, blips

        self = {}
        blips = []

        self.name = function () {
        return name
        }

        self.add = function (newBlips) {
        if (Array.isArray(newBlips)) {
            blips = blips.concat(newBlips)
        } else {
            blips.push(newBlips)
        }
        }

        self.blips = function () {
        return blips.slice(0)
        }

        return self
    }
      












    // THIS IS ring.js FROM js/models/ring.js'

    const Ring = function (name, order) {
        var self = {}
    
        self.name = function () {
          return name
        }
    
        self.order = function () {
          return order
        }
    
        return self
    }















    // THIS IS blip.js FROM js/models/blip.js'

    const IDEAL_BLIP_WIDTH = 22
    const Blip = function (name, ring, isNew, topic, description) {
        let self, blipText, isGroup, id, groupIdInGraph

        self = {}
        isGroup = false

        self.width = IDEAL_BLIP_WIDTH

        self.name = function () {
        return name
        }

        self.id = function () {
        return id || -1
        }

        self.groupBlipWidth = function () {
        return isNew ? graphConfig.newGroupBlipWidth : graphConfig.existingGroupBlipWidth
        }

        self.topic = function () {
        return topic || ''
        }

        self.description = function () {
        return description || ''
        }

        self.isNew = function () {
        return isNew
        }

        self.isGroup = function () {
        return isGroup
        }

        self.groupIdInGraph = function () {
        return groupIdInGraph || ''
        }

        self.setGroupIdInGraph = function (groupId) {
        groupIdInGraph = groupId
        }

        self.ring = function () {
        return ring
        }

        self.blipText = function () {
        return blipText || ''
        }

        self.setBlipText = function (newBlipText) {
        blipText = newBlipText
        }

        self.setId = function (newId) {
        id = newId
        }

        self.setIsGroup = function (isAGroupBlip) {
        isGroup = isAGroupBlip
        }

        return self
    }



    // THIS IS mathUtils.js FROM js/util/mathUtils.js

    function toRadian(angleInDegrees) {
      return (Math.PI * angleInDegrees) / 180
    }





    //THIS IS htmlUtils.js FROM js/util/htmlUtils.js

    function getElementWidth(element) {
      return element.node().getBoundingClientRect().width
    }
  
    function decodeHTML(encodedText) {
      const parser = new DOMParser()
      return parser.parseFromString(encodedText, 'text/html').body.textContent
    }
  
    function getElementHeight(element) {
      return element.node().getBoundingClientRect().height
    }




    // THIS IS stringUtil.js FROM js/util/stringUtil.js

    function getRingIdString(ringName) {
      return ringName.replaceAll(/[^a-zA-Z0-9]/g, '-').toLowerCase()
    }
  
    function replaceSpaceWithHyphens(anyString) {
      return anyString.trim().replace(/\s+/g, '-').toLowerCase()
    }
  
    function removeAllSpaces(blipId) {
      return blipId.toString().replace(/\s+/g, '')
    }

















  //THIS IS UTIL FROM js/components/quadrants.js

  // const { getElementWidth, getElementHeight, decodeHTML } = htmlUtil;
  // const { toRadian } = mathUtils;
  // const { getRingIdString } = stringUtil;
  // const {
  //   graphConfig,
  //   getGraphSize,
  //   getScaledQuadrantWidth,
  //   getScaledQuadrantHeightWithGap,
  //   getScale,
  //   uiConfig,
  // } = config;

  const ANIMATION_DURATION = 1000
  
  //maybe not-needed!
  const { quadrantHeight, quadrantWidth, quadrantsGap, effectiveQuadrantWidth } = graphConfig

  let prevLeft, prevTop
  let quadrantScrollHandlerReference
  let scrollFlag = false

  const createElement = (tagName, text, attributes) => {
    const tag = document.createElement(tagName)
    Object.keys(attributes).forEach((keyName) => {
      tag.setAttribute(keyName, attributes[keyName])
    })
    tag.appendChild(document.createTextNode(text))
    return tag
  }

  const replaceChild = (element, child) => {
    element.textContent = ''
    element.appendChild(child)
  }

  function selectRadarQuadrant(order, startAngle, name) {
    const noOfBlips = d3.selectAll('.quadrant-group-' + order + ' .blip-link').size()
    d3.select('#radar').classed('no-blips', noOfBlips === 0)

    d3.select('.graph-header').node().scrollIntoView({
      behavior: 'smooth',
    })

    d3.selectAll(`.quadrant-group rect`).attr('tabindex', undefined)

    const svg = d3.select('svg#radar-plot')
    svg.attr('data-quadrant-selected', order)

    svg.classed('sticky', false)
    svg.classed('enable-transition', true)

    const size = getGraphSize()

    d3.selectAll('.quadrant-table').classed('selected', false)
    d3.selectAll('.quadrant-table.' + order).classed('selected', true)

    d3.selectAll('.blip-item-description').classed('expanded', false)

    const scale = getScale()

    const adjustX = Math.sin(toRadian(startAngle)) - Math.cos(toRadian(startAngle))
    const adjustY = Math.cos(toRadian(startAngle)) + Math.sin(toRadian(startAngle))

    const translateXAll = (((1 - adjustX) / 2) * size * scale) / 2 + ((1 - adjustX) / 2) * (1 - scale / 2) * size
    const translateYAll = (((1 + adjustY) / 2) * size * scale) / 2

    const radarContainer = d3.select('#radar')
    const parentWidth = getElementWidth(radarContainer)

    const translateLeftRightValues = {
      first: {
        left: parentWidth - quadrantWidth * scale,
        top: 0,
        right: 'unset',
      },
      second: {
        left: parentWidth - quadrantWidth * scale,
        top: 0,
        right: 'unset',
      },
      third: {
        left: 0,
        top: 0,
        right: 'unset',
      },
      fourth: {
        left: 0,
        top: 0,
        right: 'unset',
      },
    }

    svg
      .style(
        'left',
        window.innerWidth < uiConfig.tabletViewWidth
          ? `calc((100% - ${quadrantWidth * scale}px) / 2)`
          : translateLeftRightValues[order].left + 'px',
      )
      .style('top', translateLeftRightValues[order].top + 'px')
      .style('right', translateLeftRightValues[order].right)
      .style('box-sizing', 'border-box')

    if (window.innerWidth < uiConfig.tabletViewWidth) {
      svg.style('margin', 'unset')
    }

    svg
      .attr('transform', `scale(${scale})`)
      .style('transform', `scale(${scale})`)
      .style('transform-origin', `0 0`)
      .attr('width', quadrantWidth)
      .attr('height', quadrantHeight + quadrantsGap)
    svg.classed('quadrant-view', true)

    const quadrantGroupTranslate = {
      first: { x: 0, y: 0 },
      second: { x: 0, y: -quadrantHeight },
      third: { x: -(quadrantWidth + quadrantsGap), y: 0 },
      fourth: { x: -(quadrantWidth + quadrantsGap), y: -quadrantHeight },
    }

    d3.select('.quadrant-group-' + order)
      .transition()
      .duration(ANIMATION_DURATION)
      .style('left', 'unset')
      .style('right', 0)
      .style('transform', `translate(${quadrantGroupTranslate[order].x}px, ${quadrantGroupTranslate[order].y}px)`)
      .attr('transform', `translate(${quadrantGroupTranslate[order].x}px, ${quadrantGroupTranslate[order].y}px)`)

    d3.selectAll('.quadrant-group-' + order + ' .blip-link text').each(function () {
      d3.select(this.parentNode).transition().duration(ANIMATION_DURATION)
    })

    const otherQuadrants = d3.selectAll('.quadrant-group:not(.quadrant-group-' + order + ')')
    otherQuadrants
      .transition()
      .duration(ANIMATION_DURATION)
      .style('pointer-events', 'none')
      .attr('transform', 'translate(' + translateXAll + ',' + translateYAll + ')scale(0)')
      .style('transform', null)
      .on('end', function () {
        otherQuadrants.style('display', 'none')
      })

    d3.selectAll('.quadrant-group').style('opacity', 0)
    d3.selectAll('.quadrant-group-' + order)
      .style('display', 'block')
      .style('opacity', '1')

    d3.select('li.quadrant-subnav__list-item.active-item').classed('active-item', false)
    d3.select(`li#subnav-item-${getRingIdString(name)}`).classed('active-item', true)

    d3.selectAll(`li.quadrant-subnav__list-item button`).attr('aria-selected', null)
    d3.select(`li#subnav-item-${getRingIdString(name)} button`).attr('aria-selected', true)

    d3.select('.quadrant-subnav__dropdown-selector').text(name)

    d3.select('#radar').classed('mobile', true) // shows the table
    d3.select('.all-quadrants-mobile').classed('show-all-quadrants-mobile', false) // hides the quadrants

    if (order === 'first' || order === 'second') {
      d3.select('.radar-legends').classed('right-view', true)
    } else {
      d3.select('.radar-legends').classed('left-view', true)
    }

    if (window.innerWidth < uiConfig.tabletViewWidth) {
      d3.select('#radar').style('height', null)
    }

    const radarLegendsContainer = d3.select('.radar-legends')
    radarLegendsContainer.style('top', `${getScaledQuadrantHeightWithGap(scale)}px`)

    d3.selectAll('.quadrant-table.selected button').attr('aria-hidden', null).attr('tabindex', null)
    d3.selectAll('.quadrant-table:not(.selected) button').attr('aria-hidden', 'true').attr('tabindex', -1)

    d3.selectAll('svg#radar-plot a').attr('aria-hidden', 'true').attr('tabindex', -1)

    d3.selectAll('.blip-list__item-container__name').attr('aria-expanded', 'false')

    if (window.innerWidth >= uiConfig.tabletViewWidth) {
      if (order === 'first' || order === 'second') {
        radarLegendsContainer.style(
          'left',
          `${
            parentWidth -
            getScaledQuadrantWidth(scale) +
            (getScaledQuadrantWidth(scale) / 2 - getElementWidth(radarLegendsContainer) / 2)
          }px`,
        )
      } else {
        radarLegendsContainer.style(
          'left',
          `${getScaledQuadrantWidth(scale) / 2 - getElementWidth(radarLegendsContainer) / 2}px`,
        )
      }

      prevLeft = d3.select('#radar-plot').style('left')
      prevTop = d3.select('#radar-plot').style('top')
      stickQuadrantOnScroll()
    } else {
      radarLegendsContainer.style('left', `${window.innerWidth / 2 - getElementWidth(radarLegendsContainer) / 2}px`)
    }
  }

  function wrapQuadrantNameInMultiLine(elem, isTopQuadrants, quadrantNameGroup, tip) {
    const maxWidth = 150
    const element = elem.node()
    const text = decodeHTML(element.innerHTML)
    const dy = isTopQuadrants ? 0 : -20

    const words = text.split(' ')
    let line = ''

    replaceChild(element, createElement('tspan', text, { id: 'text-width-check' }))
    const testElem = document.getElementById('text-width-check')

    function maxCharactersToFit(testLine, suffix) {
      let j = 1
      let firstLineWidth = 0
      const testElem1 = document.getElementById('text-width-check')
      testElem1.textContent = testLine
      if (testElem1.getBoundingClientRect().width < maxWidth) {
        return testLine.length
      }
      while (firstLineWidth < maxWidth && testLine.length > j) {
        testElem1.textContent = testLine.substring(0, j) + suffix
        firstLineWidth = testElem1.getBoundingClientRect().width

        j++
      }
      return j - 1
    }

    function ellipsis(lineBreakIndex, secondLine) {
      if (lineBreakIndex >= secondLine.length) {
        return ''
      } else {
        quadrantNameGroup.on('mouseover', () => tip.show(text, quadrantNameGroup.node()))
        quadrantNameGroup.on('mouseout', () => tip.hide(text, quadrantNameGroup.node()))
        return '...'
      }
    }

    if (testElem.getBoundingClientRect().width > maxWidth) {
      for (let i = 0; i < words.length; i++) {
        let testLine = line + words[i] + ' '
        testElem.textContent = testLine
        const textWidth = testElem.getBoundingClientRect().width

        if (textWidth > maxWidth) {
          if (i === 0) {
            let lineBreakIndex = maxCharactersToFit(testLine, '-')
            const elementText = `${words[i].substring(0, lineBreakIndex)}-`
            element.appendChild(createElement('tspan', elementText, { x: '0', dy }))
            const secondLine = words[i].substring(lineBreakIndex, words[i].length) + ' ' + words.slice(i + 1).join(' ')
            lineBreakIndex = maxCharactersToFit(secondLine, '...')
            const text = `${secondLine.substring(0, lineBreakIndex)}${ellipsis(lineBreakIndex, secondLine)}`
            element.appendChild(createElement('tspan', text, { x: '0', dy: '20' }))
            break
          } else {
            element.appendChild(createElement('tspan', line, { x: '0', dy }))
            const secondLine = words.slice(i).join(' ')
            const lineBreakIndex = maxCharactersToFit(secondLine, '...')
            const text = `${secondLine.substring(0, lineBreakIndex)}${ellipsis(lineBreakIndex, secondLine)}`
            element.appendChild(createElement('tspan', text, { x: '0', dy: '20' }))
          }
          line = words[i] + ' '
        } else {
          line = testLine
        }
      }
    } else {
      element.appendChild(createElement('tspan', text, { x: '0' }))
    }

    document.getElementById('text-width-check').remove()
  }

  function renderRadarQuadrantName(quadrant, parentGroup, tip) {
    const adjustX = Math.sin(toRadian(quadrant.startAngle)) - Math.cos(toRadian(quadrant.startAngle))
    const adjustY = -Math.cos(toRadian(quadrant.startAngle)) - Math.sin(toRadian(quadrant.startAngle))
    const quadrantNameGroup = parentGroup.append('g').classed('quadrant-name-group', true)

    let quadrantNameToDisplay = quadrant.quadrant.name()
    let translateX,
      translateY,
      anchor,
      ctaArrowXOffset,
      ctaArrowYOffset = -12

    const quadrantName = quadrantNameGroup.append('text').attr('data-quadrant-name', quadrantNameToDisplay)
    const ctaArrow = quadrantNameGroup
      .append('polygon')
      .attr('class', 'quadrant-name-cta')
      .attr('points', '5.2105e-4 11.753 1.2874 13 8 6.505 1.2879 0 0 1.2461 5.4253 6.504')
      .attr('fill', '#e16a7c')
    quadrantName.text(quadrantNameToDisplay).attr('font-weight', 'bold')

    wrapQuadrantNameInMultiLine(quadrantName, adjustY < 0, quadrantNameGroup, tip)
    const quadrantTextElement = document.querySelector(`.quadrant-group-${quadrant.order} .quadrant-name-group text`)
    const renderedText = quadrantTextElement.getBoundingClientRect()
    ctaArrowXOffset = renderedText.width + 10
    anchor = 'start'

    if (adjustX < 0) {
      translateX = 60
    } else {
      translateX = quadrantWidth * 2 - quadrantsGap - renderedText.width
    }
    if (adjustY < 0) {
      ctaArrowYOffset = quadrantTextElement.childElementCount > 1 ? 8 : ctaArrowYOffset
      translateY = 60
    } else {
      translateY = effectiveQuadrantWidth * 2 - 60
    }
    quadrantName.attr('text-anchor', anchor)
    quadrantNameGroup.attr('transform', 'translate(' + translateX + ', ' + translateY + ')')
    ctaArrow.attr('transform', `translate(${ctaArrowXOffset}, ${ctaArrowYOffset})`)
  }

  function renderRadarQuadrants(size, svg, quadrant, rings, ringCalculator, tip) {
    const quadrantGroup = svg
      .append('g')
      .attr('class', 'quadrant-group quadrant-group-' + quadrant.order)
      .on('mouseover', mouseoverQuadrant.bind({}, quadrant.order))
      .on('mouseout', mouseoutQuadrant.bind({}, quadrant.order))
      .on('click', selectRadarQuadrant.bind({}, quadrant.order, quadrant.startAngle, quadrant.quadrant.name()))
      .on('keydown', function (e) {
        if (e.key === 'Enter') selectRadarQuadrant(quadrant.order, quadrant.startAngle, quadrant.quadrant.name())
      })

    const rectCoordMap = {
      first: { x: 0, y: 0, strokeDashArray: `0, ${quadrantWidth}, ${quadrantHeight + quadrantWidth}, ${quadrantHeight}` },
      second: {
        x: 0,
        y: quadrantHeight + quadrantsGap,
        strokeDashArray: `${quadrantWidth + quadrantHeight}, ${quadrantWidth + quadrantHeight}`,
      },
      third: {
        x: quadrantWidth + quadrantsGap,
        y: 0,
        strokeDashArray: `0, ${quadrantWidth + quadrantHeight}, ${quadrantWidth + quadrantHeight}`,
      },
      fourth: {
        x: quadrantWidth + quadrantsGap,
        y: quadrantHeight + quadrantsGap,
        strokeDashArray: `${quadrantWidth}, ${quadrantWidth + quadrantHeight}, ${quadrantHeight}`,
      },
    }

    quadrantGroup
      .append('rect')
      .attr('width', `${quadrantWidth}px`)
      .attr('height', `${quadrantHeight}px`)
      .attr('fill', '#edf1f3')
      .attr('x', rectCoordMap[quadrant.order].x)
      .attr('y', rectCoordMap[quadrant.order].y)
      .style('pointer-events', 'none')

    rings.forEach(function (ring, i) {
      const arc = d3
        .arc()
        .innerRadius(ringCalculator.getRingRadius(i))
        .outerRadius(ringCalculator.getRingRadius(i + 1))
        .startAngle(toRadian(quadrant.startAngle))
        .endAngle(toRadian(quadrant.startAngle - 90))

      quadrantGroup
        .append('path')
        .attr('d', arc)
        .attr('class', 'ring-arc-' + ring.order())
        .attr(
          'transform',
          'translate(' + graphConfig.effectiveQuadrantWidth + ', ' + graphConfig.effectiveQuadrantHeight + ')',
        )
    })

    quadrantGroup
      .append('rect')
      .classed('quadrant-rect', true)
      .attr('width', `${quadrantWidth}px`)
      .attr('height', `${quadrantHeight}px`)
      .attr('fill', 'transparent')
      .attr('stroke', 'black')
      .attr('x', rectCoordMap[quadrant.order].x)
      .attr('y', rectCoordMap[quadrant.order].y)
      .attr('stroke-dasharray', rectCoordMap[quadrant.order].strokeDashArray)
      .attr('stroke-width', 2)
      .attr('stroke', '#71777d')
      .attr('tabindex', 0)

    renderRadarQuadrantName(quadrant, quadrantGroup, tip)
    return quadrantGroup
  }

  function renderRadarLegends(radarElement) {
    const legendsContainer = radarElement.append('div').classed('radar-legends', true)

    const newImage = legendsContainer
      .append('img')
      .attr('src', '../src/js/images/new.svg')
      .attr('width', '37px')
      .attr('height', '37px')
      .attr('alt', 'new blip legend icon')
      .node().outerHTML

    const existingImage = legendsContainer
      .append('img')
      .attr('src', '../src/js/images/existing.svg')
      .attr('width', '37px')
      .attr('height', '37px')
      .attr('alt', 'existing blip legend icon')
      .node().outerHTML

    legendsContainer.html(`${newImage} New ${existingImage} Existing`)
  }

  function renderMobileView(quadrant) {
    const quadrantBtn = d3.select('.all-quadrants-mobile').append('button')
    quadrantBtn
      .attr('class', 'all-quadrants-mobile--btn')
      .style('background-image', `url('/images/${quadrant.order}-quadrant-btn-bg.svg')`)
      .attr('id', quadrant.order + '-quadrant-mobile')
      .append('div')
      .attr('class', 'btn-text-wrapper')
      .text(quadrant.quadrant.name().replace(/[^a-zA-Z0-9\s!&]/g, ' '))
    quadrantBtn.node().onclick = () => {
      selectRadarQuadrant(quadrant.order, quadrant.startAngle, quadrant.quadrant.name())
    }
  }

  function mouseoverQuadrant(order) {
    d3.select('.quadrant-group-' + order).style('opacity', 1)
    d3.selectAll('.quadrant-group:not(.quadrant-group-' + order + ')').style('opacity', 0.3)
  }

  function mouseoutQuadrant(order) {
    d3.selectAll('.quadrant-group:not(.quadrant-group-' + order + ')').style('opacity', 1)
  }

  function quadrantScrollHandler(
    scale,
    radarElement,
    offset,
    selectedOrder,
    leftQuadrantLeftValue,
    rightQuadrantLeftValue,
    radarHeight,
    selectedQuadrantTable,
    radarLegendsContainer,
    radarLegendsWidth,
  ) {
    const quadrantTableHeight = getElementHeight(selectedQuadrantTable)
    const quadrantTableOffset = offset + quadrantTableHeight

    if (window.scrollY >= offset) {
      radarElement.classed('enable-transition', false)
      radarElement.classed('sticky', true)
      radarLegendsContainer.classed('sticky', true)

      if (window.scrollY + uiConfig.subnavHeight + radarHeight >= quadrantTableOffset) {
        radarElement.classed('sticky', false)
        radarLegendsContainer.classed('sticky', false)

        radarElement.style('top', `${quadrantTableHeight - radarHeight - uiConfig.subnavHeight}px`)
        radarElement.style('left', prevLeft)

        radarLegendsContainer.style(
          'top',
          `${quadrantTableHeight - radarHeight - uiConfig.subnavHeight + getScaledQuadrantHeightWithGap(scale)}px`,
        )
        radarLegendsContainer.style(
          'left',
          `${parseFloat(prevLeft.slice(0, -2)) + (getScaledQuadrantWidth(scale) / 2 - radarLegendsWidth / 2)}px`,
        )
      } else {
        if (selectedOrder === 'first' || selectedOrder === 'second') {
          radarElement.style('left', `${leftQuadrantLeftValue}px`)
          radarLegendsContainer.style(
            'left',
            `${
              leftQuadrantLeftValue + (getScaledQuadrantWidth(scale) / 2 - getElementWidth(radarLegendsContainer) / 2)
            }px`,
          )
        } else {
          radarElement.style('left', `${rightQuadrantLeftValue}px`)
          radarLegendsContainer.style(
            'left',
            `${
              rightQuadrantLeftValue + (getScaledQuadrantWidth(scale) / 2 - getElementWidth(radarLegendsContainer) / 2)
            }px`,
          )
        }

        radarLegendsContainer.style('top', `${getScaledQuadrantHeightWithGap(scale) + uiConfig.subnavHeight}px`)
      }
    } else {
      radarElement.style('top', prevTop)
      radarElement.style('left', prevLeft)
      radarElement.classed('sticky', false)

      radarLegendsContainer.style('top', `${parseFloat(prevTop.slice(0, -2)) + getScaledQuadrantHeightWithGap(scale)}px`)
      radarLegendsContainer.style(
        'left',
        `${parseFloat(prevLeft.slice(0, -2)) + (getScaledQuadrantWidth(scale) / 2 - radarLegendsWidth / 2)}px`,
      )
      radarLegendsContainer.classed('sticky', false)
    }
  }

  function stickQuadrantOnScroll() {
    if (!scrollFlag) {
      const scale = getScale()

      const radarContainer = d3.select('#radar')
      const radarElement = d3.select('#radar-plot')
      const selectedQuadrantTable = d3.select('.quadrant-table.selected')
      const radarLegendsContainer = d3.select('.radar-legends')

      const radarHeight = quadrantHeight * scale + quadrantsGap * scale
      const offset = radarContainer.node().offsetTop - uiConfig.subnavHeight
      const radarWidth = radarContainer.node().getBoundingClientRect().width
      const selectedOrder = radarElement.attr('data-quadrant-selected')

      const leftQuadrantLeftValue =
        (window.innerWidth + radarWidth) / 2 - effectiveQuadrantWidth * scale + (quadrantsGap / 2) * scale
      const rightQuadrantLeftValue = (window.innerWidth - radarWidth) / 2

      const radarLegendsWidth = getElementWidth(radarLegendsContainer)

      quadrantScrollHandlerReference = quadrantScrollHandler.bind(
        this,
        scale,
        radarElement,
        offset,
        selectedOrder,
        leftQuadrantLeftValue,
        rightQuadrantLeftValue,
        radarHeight,
        selectedQuadrantTable,
        radarLegendsContainer,
        radarLegendsWidth,
      )

      if (
        uiConfig.subnavHeight + radarHeight + quadrantsGap * 2 + uiConfig.legendsHeight <
        getElementHeight(selectedQuadrantTable)
      ) {
        window.addEventListener('scroll', quadrantScrollHandlerReference)
        scrollFlag = true
      } else {
        removeScrollListener()
      }
    }
  }

  function removeScrollListener() {
    window.removeEventListener('scroll', quadrantScrollHandlerReference)
    scrollFlag = false
  }












    //THIS IS UTIL FROM js/components/quadrantTables.js
    
    //THIS IS UTIL FROM js/components/quadrantTables.js

    // THIS IS COMING FROM js/graphing/blips.js





    // THIS IS COMING FROM js/graphing/radar.js























    const featureToggles = mainConfig.featureToggles;
    //const { getGraphSize, graphConfig } = config;


    
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
  
  define(function modelRing() {
    const Rings = function (name, order) {
      var self = {}
  
      self.name = function () {
        return name
      }
  
      self.order = function () {
        return order
      }
  
      return self
    }
    
    return Rings
  })
  
  