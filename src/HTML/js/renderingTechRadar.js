/**
 * The code below is stitched together from Thoughtworks build your radar open sourced project
 * 
 * Its Thoughtwork file path is indicated as a comment above each code block
 * 
 * @url [Thoughtworks](https://github.com/thoughtworks/build-your-own-radar/)
 */

define([
  "d3",
  "d3tip",
  "d3-collection",
  "d3-selection",
  "chance",
  "lodash",
  "jquery",
  "jquery-autocomplete",
], function facModel(d3, d3tip, d3col, d3sel, Chance, _, $, AutoComplete) {
  

  /**
   * ROOT JS: js/config.js
   */
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
    };

    environ = window.APP_CONFIG = {
      ENVIRONMENT: "production",
      featureToggles: {
        UIRefresh2022: true,
      },
    };

    return window.APP_CONFIG && window.APP_CONFIG.ENVIRONMENT
      ? env[window.APP_CONFIG.ENVIRONMENT]
      : env.development;
  };
  
  /** 
   * GRAPHING: js/graphing/config.js
   * Change your quadrant names and ring names here
   */
  const quadrantSize = 512;
  const quadrantGap = 32;

  const quadrantNames =
    '["Techniques", "Platforms", "Tools", "Languages & Frameworks"]';
  const ringNames = '["Adopt", "Trial", "Assess", "Hold"]';

  const getQuadrants = () => {
    return JSON.parse(quadrantNames);
  };

  const getRings = () => {
    return JSON.parse(ringNames);
  };

  const isBetween = (number, startNumber, endNumber) => {
    return startNumber <= number && number <= endNumber;
  };
  const isValidConfig = () => {
    return getQuadrants().length === 4 && isBetween(getRings().length, 1, 4);
  };

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
  };

  const uiConfig = {
    subnavHeight: 60,
    bannerHeight: 200,
    tabletBannerHeight: 300,
    headerHeight: 80,
    legendsHeight: 42,
    tabletViewWidth: 1280,
    mobileViewWidth: 768,
  };

  function getScale() {
    return window.innerWidth < 1800 ? 1.25 : 1.5;
  }

  function getGraphSize() {
    return (
      graphConfig.effectiveQuadrantHeight + graphConfig.effectiveQuadrantWidth
    );
  }

  function getScaledQuadrantWidth(scale) {
    return graphConfig.quadrantWidth * scale;
  }

  function getScaledQuadrantHeightWithGap(scale) {
    return (graphConfig.quadrantHeight + graphConfig.quadrantsGap) * scale;
  }

  /** 
   * MODELS: js/models/quadrant.js
  */
  const Quadrant = function (name) {
    var self, blips;

    self = {};
    blips = [];

    self.name = function () {
      return name;
    };

    self.add = function (newBlips) {
      if (Array.isArray(newBlips)) {
        blips = blips.concat(newBlips);
      } else {
        blips.push(newBlips);
      }
    };

    self.blips = function () {
      return blips.slice(0);
    };

    return self;
  };

  /** 
   * MODELS: js/models/radar.js
   */   
  const Radar = function () {
    const featureToggles1 = mainConfig().featureToggles;

    let self,
      quadrants,
      blipNumber,
      addingQuadrant,
      alternatives,
      currentSheetName,
      rings;

    blipNumber = 0;
    addingQuadrant = 0;
    quadrants = featureToggles1.UIRefresh2022
      ? [
          { order: "first", startAngle: 0 },
          { order: "second", startAngle: -90 },
          { order: "third", startAngle: 90 },
          { order: "fourth", startAngle: -180 },
        ]
      : [
          { order: "first", startAngle: 90 },
          { order: "second", startAngle: 0 },
          { order: "third", startAngle: -90 },
          { order: "fourth", startAngle: -180 },
        ];
    alternatives = [];
    currentSheetName = "";
    self = {};
    rings = {};

    function setNumbers(blips) {
      blips.forEach(function (blip) {
        ++blipNumber;
        blip.setBlipText(blipNumber);
        blip.setId(blipNumber);
      });
    }

    self.addAlternative = function (sheetName) {
      alternatives.push(sheetName);
    };

    self.getAlternatives = function () {
      return alternatives;
    };

    self.setCurrentSheet = function (sheetName) {
      currentSheetName = sheetName;
    };

    self.getCurrentSheet = function () {
      return currentSheetName;
    };

    self.addQuadrant = function (quadrant) {
      quadrants[addingQuadrant].quadrant = quadrant;
      setNumbers(quadrant.blips());
      addingQuadrant++;
    };
    self.addRings = function (allRings) {
      rings = allRings;
    };

    function allQuadrants() {
      return _.map(quadrants, "quadrant");
    }

    function allBlips() {
      return allQuadrants().reduce(function (blips, quadrant) {
        return blips.concat(quadrant.blips());
      }, []);
    }

    self.rings = function () {
      if (featureToggles1.UIRefresh2022) {
        return rings;
      }

      return _.sortBy(
        _.map(
          _.uniqBy(allBlips(), function (blip) {
            return blip.ring().name();
          }),
          function (blip) {
            return blip.ring();
          }
        ),
        function (ring) {
          return ring.order();
        }
      );
    };

    self.quadrants = function () {
      return quadrants;
    };

    return self;
  };

  /**
   * MODELS: js/models/ring.js 
   */
  const Ring = function (name, order) {
    var self = {};

    self.name = function () {
      return name;
    };

    self.order = function () {
      return order;
    };

    return self;
  };

  /**
   * MODELS: js/models/blip.js' 
   */
  const IDEAL_BLIP_WIDTH = 22;
  const Blip = function (name, ring, isNew, topic, description) {
    let self, blipText, isGroup, id, groupIdInGraph;

    self = {};
    isGroup = false;

    self.width = IDEAL_BLIP_WIDTH;

    self.name = function () {
      return name;
    };

    self.id = function () {
      return id || -1;
    };

    self.groupBlipWidth = function () {
      return isNew
        ? graphConfig.newGroupBlipWidth
        : graphConfig.existingGroupBlipWidth;
    };

    self.topic = function () {
      return topic || "";
    };

    self.description = function () {
      return description || "";
    };

    self.isNew = function () {
      return isNew;
    };

    self.isGroup = function () {
      return isGroup;
    };

    self.groupIdInGraph = function () {
      return groupIdInGraph || "";
    };

    self.setGroupIdInGraph = function (groupId) {
      groupIdInGraph = groupId;
    };

    self.ring = function () {
      return ring;
    };

    self.blipText = function () {
      return blipText || "";
    };

    self.setBlipText = function (newBlipText) {
      blipText = newBlipText;
    };

    self.setId = function (newId) {
      id = newId;
    };

    self.setIsGroup = function (isAGroupBlip) {
      isGroup = isAGroupBlip;
    };

    return self;
  };
  
  /**
   * UTIL: js/util/mathUtils.js  
   */
  function toRadian(angleInDegrees) {
    return (Math.PI * angleInDegrees) / 180;
  }

  /**
   * UTIL: js/util/htmlUtils.js
   */
  function getElementWidth(element) {
    return element.node().getBoundingClientRect().width;
  }

  function decodeHTML(encodedText) {
    const parser = new DOMParser();
    return parser.parseFromString(encodedText, "text/html").body.textContent;
  }

  function getElementHeight(element) {
    return element.node().getBoundingClientRect().height;
  }
  
  /**
   * UTIL: js/util/stringUtil.js  
   */
  function getRingIdString(ringName) {
    return ringName.replaceAll(/[^a-zA-Z0-9]/g, "-").toLowerCase();
  }

  function replaceSpaceWithHyphens(anyString) {
    return anyString.trim().replace(/\s+/g, "-").toLowerCase();
  }

  function removeAllSpaces(blipId) {
    return blipId.toString().replace(/\s+/g, "");
  }


  /**
   * COMPONENTS: js/graphing/components/quadrants.js
   */
  const ANIMATION_DURATION = 1000;
  
  const {
    quadrantHeight,
    quadrantWidth,
    quadrantsGap,
    effectiveQuadrantWidth,
  } = graphConfig;

  let prevLeft, prevTop;
  let quadrantScrollHandlerReference;
  let scrollFlag = false;

  function selectRadarQuadrant(order, startAngle, name) {
    const noOfBlips = d3
      .selectAll(".quadrant-group-" + order + " .blip-link")
      .size();
    d3.select("#radar").classed("no-blips", noOfBlips === 0);

    d3.select(".graph-header").node().scrollIntoView({
      behavior: "smooth",
    });

    d3.selectAll(`.quadrant-group rect`).attr("tabindex", undefined);

    const svg = d3.select("svg#radar-plot");
    svg.attr("data-quadrant-selected", order);

    svg.classed("sticky", false);
    svg.classed("enable-transition", true);

    const size = getGraphSize();

    d3.selectAll(".quadrant-table").classed("selected", false);
    d3.selectAll(".quadrant-table." + order).classed("selected", true);

    d3.selectAll(".blip-item-description").classed("expanded", false);

    const scale = getScale();

    const adjustX =
      Math.sin(toRadian(startAngle)) - Math.cos(toRadian(startAngle));
    const adjustY =
      Math.cos(toRadian(startAngle)) + Math.sin(toRadian(startAngle));

    const translateXAll =
      (((1 - adjustX) / 2) * size * scale) / 2 +
      ((1 - adjustX) / 2) * (1 - scale / 2) * size;
    const translateYAll = (((1 + adjustY) / 2) * size * scale) / 2;

    const radarContainer = d3.select("#radar");
    const parentWidth = getElementWidth(radarContainer);

    const translateLeftRightValues = {
      first: {
        left: parentWidth - quadrantWidth * scale,
        top: 0,
        right: "unset",
      },
      second: {
        left: parentWidth - quadrantWidth * scale,
        top: 0,
        right: "unset",
      },
      third: {
        left: 0,
        top: 0,
        right: "unset",
      },
      fourth: {
        left: 0,
        top: 0,
        right: "unset",
      },
    };

    svg
      .style(
        "left",
        window.innerWidth < uiConfig.tabletViewWidth
          ? `calc((100% - ${quadrantWidth * scale}px) / 2)`
          : translateLeftRightValues[order].left + "px"
      )
      .style("top", translateLeftRightValues[order].top + "px")
      .style("right", translateLeftRightValues[order].right)
      .style("box-sizing", "border-box");

    if (window.innerWidth < uiConfig.tabletViewWidth) {
      svg.style("margin", "unset");
    }

    svg
      .attr("transform", `scale(${scale})`)
      .style("transform", `scale(${scale})`)
      .style("transform-origin", `0 0`)
      .attr("width", quadrantWidth)
      .attr("height", quadrantHeight + quadrantsGap);
    svg.classed("quadrant-view", true);

    const quadrantGroupTranslate = {
      first: { x: 0, y: 0 },
      second: { x: 0, y: -quadrantHeight },
      third: { x: -(quadrantWidth + quadrantsGap), y: 0 },
      fourth: { x: -(quadrantWidth + quadrantsGap), y: -quadrantHeight },
    };

    d3.select(".quadrant-group-" + order)
      .transition()
      .duration(ANIMATION_DURATION)
      .style("left", "unset")
      .style("right", 0)
      .style(
        "transform",
        `translate(${quadrantGroupTranslate[order].x}px, ${quadrantGroupTranslate[order].y}px)`
      )
      .attr(
        "transform",
        `translate(${quadrantGroupTranslate[order].x}px, ${quadrantGroupTranslate[order].y}px)`
      );

    d3.selectAll(".quadrant-group-" + order + " .blip-link text").each(
      function () {
        d3.select(this.parentNode).transition().duration(ANIMATION_DURATION);
      }
    );

    const otherQuadrants = d3.selectAll(
      ".quadrant-group:not(.quadrant-group-" + order + ")"
    );
    otherQuadrants
      .transition()
      .duration(ANIMATION_DURATION)
      .style("pointer-events", "none")
      .attr(
        "transform",
        "translate(" + translateXAll + "," + translateYAll + ")scale(0)"
      )
      .style("transform", null)
      .on("end", function () {
        otherQuadrants.style("display", "none");
      });

    d3.selectAll(".quadrant-group").style("opacity", 0);
    d3.selectAll(".quadrant-group-" + order)
      .style("display", "block")
      .style("opacity", "1");

    d3.select("li.quadrant-subnav__list-item.active-item").classed(
      "active-item",
      false
    );
    d3.select(`li#subnav-item-${getRingIdString(name)}`).classed(
      "active-item",
      true
    );

    d3.selectAll(`li.quadrant-subnav__list-item button`).attr(
      "aria-selected",
      null
    );
    d3.select(`li#subnav-item-${getRingIdString(name)} button`).attr(
      "aria-selected",
      true
    );

    d3.select(".quadrant-subnav__dropdown-selector").text(name);

    d3.select("#radar").classed("mobile", true); // shows the table
    d3.select(".all-quadrants-mobile").classed(
      "show-all-quadrants-mobile",
      false
    ); // hides the quadrants

    if (order === "first" || order === "second") {
      d3.select(".radar-legends").classed("right-view", true);
    } else {
      d3.select(".radar-legends").classed("left-view", true);
    }

    if (window.innerWidth < uiConfig.tabletViewWidth) {
      d3.select("#radar").style("height", null);
    }

    const radarLegendsContainer = d3.select(".radar-legends");
    radarLegendsContainer.style(
      "top",
      `${getScaledQuadrantHeightWithGap(scale)}px`
    );

    d3.selectAll(".quadrant-table.selected button")
      .attr("aria-hidden", null)
      .attr("tabindex", null);
    d3.selectAll(".quadrant-table:not(.selected) button")
      .attr("aria-hidden", "true")
      .attr("tabindex", -1);

    d3.selectAll("svg#radar-plot a")
      .attr("aria-hidden", "true")
      .attr("tabindex", -1);

    d3.selectAll(".blip-list__item-container__name").attr(
      "aria-expanded",
      "false"
    );

    if (window.innerWidth >= uiConfig.tabletViewWidth) {
      if (order === "first" || order === "second") {
        radarLegendsContainer.style(
          "left",
          `${
            parentWidth -
            getScaledQuadrantWidth(scale) +
            (getScaledQuadrantWidth(scale) / 2 -
              getElementWidth(radarLegendsContainer) / 2)
          }px`
        );
      } else {
        radarLegendsContainer.style(
          "left",
          `${
            getScaledQuadrantWidth(scale) / 2 -
            getElementWidth(radarLegendsContainer) / 2
          }px`
        );
      }

      prevLeft = d3.select("#radar-plot").style("left");
      prevTop = d3.select("#radar-plot").style("top");
      stickQuadrantOnScroll();
    } else {
      radarLegendsContainer.style(
        "left",
        `${
          window.innerWidth / 2 - getElementWidth(radarLegendsContainer) / 2
        }px`
      );
    }
  }

  function wrapQuadrantNameInMultiLine(
    elem,
    isTopQuadrants,
    quadrantNameGroup,
    tip
  ) {
    const maxWidth = 150;
    const element = elem.node();
    const text = decodeHTML(element.innerHTML);
    const dy = isTopQuadrants ? 0 : -20;

    const words = text.split(" ");
    let line = "";

    element.innerHTML = `<tspan id="text-width-check">${text}</tspan >`
    
    const testElem = document.getElementById("text-width-check");

    function maxCharactersToFit(testLine, suffix) {
      let j = 1;
      let firstLineWidth = 0;
      const testElem1 = document.getElementById("text-width-check");
      testElem1.innerHTML = testLine
      if (testElem1.getBoundingClientRect().width < maxWidth) {
        return testLine.length;
      }
      while (firstLineWidth < maxWidth && testLine.length > j) {
        testElem1.innerHTML = testLine.substring(0, j) + suffix
        firstLineWidth = testElem1.getBoundingClientRect().width;

        j++;
      }
      return j - 1;
    }

    function ellipsis(lineBreakIndex, secondLine) {
      if (lineBreakIndex >= secondLine.length) {
        return "";
      } else {
        quadrantNameGroup.on("mouseover", () =>
          tip.show(text, quadrantNameGroup.node())
        );
        quadrantNameGroup.on("mouseout", () =>
          tip.hide(text, quadrantNameGroup.node())
        );
        return "...";
      }
    }

    if (testElem.getBoundingClientRect().width > maxWidth) {
      for (let i = 0; i < words.length; i++) {
        let testLine = line + words[i] + " ";
        testElem.innerHTML = testLine
        const textWidth = testElem.getBoundingClientRect().width;

        if (textWidth > maxWidth) {
          if (i === 0) {
            let lineBreakIndex = maxCharactersToFit(testLine, "-");
            const elementText = `${words[i].substring(0, lineBreakIndex)}-`;
            element.innerHTML += '<tspan x="0" dy="' + dy + '">' + words[i].substring(0, lineBreakIndex) + '-</tspan>'
          
            const secondLine =
              words[i].substring(lineBreakIndex, words[i].length) +
              " " +
              words.slice(i + 1).join(" ");
            lineBreakIndex = maxCharactersToFit(secondLine, "...");
            const text = `${secondLine.substring(0, lineBreakIndex)}${ellipsis(
              lineBreakIndex,
              secondLine
            )}`;

            element.innerHTML +=
            '<tspan x="0" dy="' +
            20 +
            '">' +
            secondLine.substring(0, lineBreakIndex) +
            ellipsis(lineBreakIndex, secondLine) +
            '</tspan>'
            
            break;
          } else {
            element.innerHTML += '<tspan x="0" dy="' + dy + '">' + line + '</tspan>'
            
            const secondLine = words.slice(i).join(" ");
            const lineBreakIndex = maxCharactersToFit(secondLine, "...");
            const text = `${secondLine.substring(0, lineBreakIndex)}${ellipsis(
              lineBreakIndex,
              secondLine
            )}`;

            element.innerHTML +=
            '<tspan x="0" dy="' +
            20 +
            '">' +
            secondLine.substring(0, lineBreakIndex) +
            ellipsis(lineBreakIndex, secondLine) +
            '</tspan>'
        
            
          }
          line = words[i] + " ";
        } else {
          line = testLine;
        }
      }
    } else {
      element.innerHTML += `<tspan x="0" class="">` + text + `</tspan>`;
    }

    document.getElementById("text-width-check").remove();
  }

  function renderRadarQuadrantName(quadrant, parentGroup, tip) {
    const adjustX =
      Math.sin(toRadian(quadrant.startAngle)) -
      Math.cos(toRadian(quadrant.startAngle));
    const adjustY =
      -Math.cos(toRadian(quadrant.startAngle)) -
      Math.sin(toRadian(quadrant.startAngle));
    const quadrantNameGroup = parentGroup
      .append("g")
      .classed("quadrant-name-group", true);

    let quadrantNameToDisplay = quadrant.quadrant.name();
    let translateX,
      translateY,
      anchor,
      ctaArrowXOffset,
      ctaArrowYOffset = -12;

    const quadrantName = quadrantNameGroup
      .append("text")
      .attr("data-quadrant-name", quadrantNameToDisplay);
    const ctaArrow = quadrantNameGroup
      .append("polygon")
      .attr("class", "quadrant-name-cta")
      .attr(
        "points",
        "5.2105e-4 11.753 1.2874 13 8 6.505 1.2879 0 0 1.2461 5.4253 6.504"
      )
      .attr("fill", "#e16a7c");
    quadrantName.text(quadrantNameToDisplay).attr("font-weight", "bold");

    wrapQuadrantNameInMultiLine(
      quadrantName,
      adjustY < 0,
      quadrantNameGroup,
      tip
    );
    const quadrantTextElement = document.querySelector(
      `.quadrant-group-${quadrant.order} .quadrant-name-group text`
    );
    const renderedText = quadrantTextElement.getBoundingClientRect();
    ctaArrowXOffset = renderedText.width + 10;
    anchor = "start";

    if (adjustX < 0) {
      translateX = 60;
    } else {
      translateX = quadrantWidth * 2 - quadrantsGap - renderedText.width;
    }
    if (adjustY < 0) {
      ctaArrowYOffset =
        quadrantTextElement.childElementCount > 1 ? 8 : ctaArrowYOffset;
      translateY = 60;
    } else {
      translateY = effectiveQuadrantWidth * 2 - 60;
    }
    quadrantName.attr("text-anchor", anchor);
    quadrantNameGroup.attr(
      "transform",
      "translate(" + translateX + ", " + translateY + ")"
    );
    ctaArrow.attr(
      "transform",
      `translate(${ctaArrowXOffset}, ${ctaArrowYOffset})`
    );
  }

  function renderRadarQuadrants(
    size,
    svg,
    quadrant,
    rings,
    ringCalculator,
    tip
  ) {
    const quadrantGroup = svg
      .append("g")
      .attr("class", "quadrant-group quadrant-group-" + quadrant.order)
      .on("mouseover", mouseoverQuadrant.bind({}, quadrant.order))
      .on("mouseout", mouseoutQuadrant.bind({}, quadrant.order))
      .on(
        "click",
        selectRadarQuadrant.bind(
          {},
          quadrant.order,
          quadrant.startAngle,
          quadrant.quadrant.name()
        )
      )
      .on("keydown", function (e) {
        if (e.key === "Enter")
          selectRadarQuadrant(
            quadrant.order,
            quadrant.startAngle,
            quadrant.quadrant.name()
          );
      });

    const rectCoordMap = {
      first: {
        x: 0,
        y: 0,
        strokeDashArray: `0, ${quadrantWidth}, ${
          quadrantHeight + quadrantWidth
        }, ${quadrantHeight}`,
      },
      second: {
        x: 0,
        y: quadrantHeight + quadrantsGap,
        strokeDashArray: `${quadrantWidth + quadrantHeight}, ${
          quadrantWidth + quadrantHeight
        }`,
      },
      third: {
        x: quadrantWidth + quadrantsGap,
        y: 0,
        strokeDashArray: `0, ${quadrantWidth + quadrantHeight}, ${
          quadrantWidth + quadrantHeight
        }`,
      },
      fourth: {
        x: quadrantWidth + quadrantsGap,
        y: quadrantHeight + quadrantsGap,
        strokeDashArray: `${quadrantWidth}, ${
          quadrantWidth + quadrantHeight
        }, ${quadrantHeight}`,
      },
    };

    quadrantGroup
      .append("rect")
      .attr("width", `${quadrantWidth}px`)
      .attr("height", `${quadrantHeight}px`)
      .attr("fill", "#edf1f3")
      .attr("x", rectCoordMap[quadrant.order].x)
      .attr("y", rectCoordMap[quadrant.order].y)
      .style("pointer-events", "none");

    rings.forEach(function (ring, i) {
      const arc = d3
        .arc()
        .innerRadius(ringCalculator.getRingRadius(i))
        .outerRadius(ringCalculator.getRingRadius(i + 1))
        .startAngle(toRadian(quadrant.startAngle))
        .endAngle(toRadian(quadrant.startAngle - 90));

      quadrantGroup
        .append("path")
        .attr("d", arc)
        .attr("class", "ring-arc-" + ring.order())
        .attr(
          "transform",
          "translate(" +
            graphConfig.effectiveQuadrantWidth +
            ", " +
            graphConfig.effectiveQuadrantHeight +
            ")"
        );
    });

    quadrantGroup
      .append("rect")
      .classed("quadrant-rect", true)
      .attr("width", `${quadrantWidth}px`)
      .attr("height", `${quadrantHeight}px`)
      .attr("fill", "transparent")
      .attr("stroke", "black")
      .attr("x", rectCoordMap[quadrant.order].x)
      .attr("y", rectCoordMap[quadrant.order].y)
      .attr("stroke-dasharray", rectCoordMap[quadrant.order].strokeDashArray)
      .attr("stroke-width", 2)
      .attr("stroke", "#71777d")
      .attr("tabindex", 0);

    renderRadarQuadrantName(quadrant, quadrantGroup, tip);
    return quadrantGroup;
  }

  function renderRadarLegends(radarElement) {
    const legendsContainer = radarElement
      .append("div")
      .classed("radar-legends", true);

    const newImage = legendsContainer
      .append("img")
      .attr("src", "./HTML/images/new.svg")
      .attr("width", "37px")
      .attr("height", "37px")
      .attr("alt", "new blip legend icon")
      .node().outerHTML;

    const existingImage = legendsContainer
      .append("img")
      .attr("src", "./HTML/images/existing.svg")
      .attr("width", "37px")
      .attr("height", "37px")
      .attr("alt", "existing blip legend icon")
      .node().outerHTML;

    legendsContainer.html(`${newImage} New ${existingImage} Existing`);
  }

  function renderMobileView(quadrant) {
    const quadrantBtn = d3.select(".all-quadrants-mobile").append("button");
    quadrantBtn
      .attr("class", "all-quadrants-mobile--btn")
      .style(
        "background-image",
        `url('./HTML/images/${quadrant.order}-quadrant-btn-bg.svg')`
      )
      .attr("id", quadrant.order + "-quadrant-mobile")
      .append("div")
      .attr("class", "btn-text-wrapper")
      .text(quadrant.quadrant.name().replace(/[^a-zA-Z0-9\s!&]/g, " "));
    quadrantBtn.node().onclick = () => {
      selectRadarQuadrant(
        quadrant.order,
        quadrant.startAngle,
        quadrant.quadrant.name()
      );
    };
  }

  function mouseoverQuadrant(order) {
    d3.select(".quadrant-group-" + order).style("opacity", 1);
    d3.selectAll(".quadrant-group:not(.quadrant-group-" + order + ")").style(
      "opacity",
      0.3
    );
  }

  function mouseoutQuadrant(order) {
    d3.selectAll(".quadrant-group:not(.quadrant-group-" + order + ")").style(
      "opacity",
      1
    );
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
    radarLegendsWidth
  ) {
    const quadrantTableHeight = getElementHeight(selectedQuadrantTable);
    const quadrantTableOffset = offset + quadrantTableHeight;

    if (window.scrollY >= offset) {
      radarElement.classed("enable-transition", false);
      radarElement.classed("sticky", true);
      radarLegendsContainer.classed("sticky", true);

      if (
        window.scrollY + uiConfig.subnavHeight + radarHeight >=
        quadrantTableOffset
      ) {
        radarElement.classed("sticky", false);
        radarLegendsContainer.classed("sticky", false);

        radarElement.style(
          "top",
          `${quadrantTableHeight - radarHeight - uiConfig.subnavHeight}px`
        );
        radarElement.style("left", prevLeft);

        radarLegendsContainer.style(
          "top",
          `${
            quadrantTableHeight -
            radarHeight -
            uiConfig.subnavHeight +
            getScaledQuadrantHeightWithGap(scale)
          }px`
        );
        radarLegendsContainer.style(
          "left",
          `${
            parseFloat(prevLeft.slice(0, -2)) +
            (getScaledQuadrantWidth(scale) / 2 - radarLegendsWidth / 2)
          }px`
        );
      } else {
        if (selectedOrder === "first" || selectedOrder === "second") {
          radarElement.style("left", `${leftQuadrantLeftValue}px`);
          radarLegendsContainer.style(
            "left",
            `${
              leftQuadrantLeftValue +
              (getScaledQuadrantWidth(scale) / 2 -
                getElementWidth(radarLegendsContainer) / 2)
            }px`
          );
        } else {
          radarElement.style("left", `${rightQuadrantLeftValue}px`);
          radarLegendsContainer.style(
            "left",
            `${
              rightQuadrantLeftValue +
              (getScaledQuadrantWidth(scale) / 2 -
                getElementWidth(radarLegendsContainer) / 2)
            }px`
          );
        }

        radarLegendsContainer.style(
          "top",
          `${getScaledQuadrantHeightWithGap(scale) + uiConfig.subnavHeight}px`
        );
      }
    } else {
      radarElement.style("top", prevTop);
      radarElement.style("left", prevLeft);
      radarElement.classed("sticky", false);

      radarLegendsContainer.style(
        "top",
        `${
          parseFloat(prevTop.slice(0, -2)) +
          getScaledQuadrantHeightWithGap(scale)
        }px`
      );
      radarLegendsContainer.style(
        "left",
        `${
          parseFloat(prevLeft.slice(0, -2)) +
          (getScaledQuadrantWidth(scale) / 2 - radarLegendsWidth / 2)
        }px`
      );
      radarLegendsContainer.classed("sticky", false);
    }
  }

  function stickQuadrantOnScroll() {
    if (!scrollFlag) {
      const scale = getScale();

      const radarContainer = d3.select("#radar");
      const radarElement = d3.select("#radar-plot");
      const selectedQuadrantTable = d3.select(".quadrant-table.selected");
      const radarLegendsContainer = d3.select(".radar-legends");

      const radarHeight = quadrantHeight * scale + quadrantsGap * scale;
      const offset = radarContainer.node().offsetTop - uiConfig.subnavHeight;
      const radarWidth = radarContainer.node().getBoundingClientRect().width;
      const selectedOrder = radarElement.attr("data-quadrant-selected");

      const leftQuadrantLeftValue =
        (window.innerWidth + radarWidth) / 2 -
        effectiveQuadrantWidth * scale +
        (quadrantsGap / 2) * scale;
      const rightQuadrantLeftValue = (window.innerWidth - radarWidth) / 2;

      const radarLegendsWidth = getElementWidth(radarLegendsContainer);

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
        radarLegendsWidth
      );

      if (
        uiConfig.subnavHeight +
          radarHeight +
          quadrantsGap * 2 +
          uiConfig.legendsHeight <
        getElementHeight(selectedQuadrantTable)
      ) {
        window.addEventListener("scroll", quadrantScrollHandlerReference);
        scrollFlag = true;
      } else {
        removeScrollListener();
      }
    }
  }

  function removeScrollListener() {
    window.removeEventListener("scroll", quadrantScrollHandlerReference);
    scrollFlag = false;
  }
  
  /**
   * COMPONENTS: js/graphing/components/quadrantTables.js
   */
  function fadeOutAllBlips() {
    d3.selectAll("g > a.blip-link").attr("opacity", 0.3);
  }

  function fadeInSelectedBlip(selectedBlipOnGraph) {
    selectedBlipOnGraph.attr("opacity", 1.0);
  }

  function highlightBlipInTable(selectedBlip) {
    selectedBlip.classed("highlight", true);
  }

  function highlightBlipInGraph(blipIdToFocus) {
    fadeOutAllBlips();
    const selectedBlipOnGraph = d3.select(
      `g > a.blip-link[data-blip-id='${blipIdToFocus}'`
    );
    fadeInSelectedBlip(selectedBlipOnGraph);
  }

  function renderBlipDescription(
    blip,
    ring,
    quadrant,
    tip,
    groupBlipTooltipText
  ) {
    let blipTableItem = d3.select(
      `.quadrant-table.${quadrant.order} ul[data-ring-order='${ring.order()}']`
    );
    if (!groupBlipTooltipText) {
      blipTableItem = blipTableItem
        .append("li")
        .classed("blip-list__item", true);
      const blipItemDiv = blipTableItem
        .append("div")
        .classed("blip-list__item-container", true)
        .attr("data-blip-id", blip.id());

      if (blip.groupIdInGraph()) {
        blipItemDiv.attr("data-group-id", blip.groupIdInGraph());
      }

      const blipItemContainer = blipItemDiv
        .append("button")
        .classed("blip-list__item-container__name", true)
        .attr("aria-expanded", "false")
        .attr("aria-controls", `blip-description-${blip.id()}`)
        .attr("aria-hidden", "true")
        .attr("tabindex", -1)
        .on("click search-result-click", function (e) {
          e.stopPropagation();

          const expandFlag = d3
            .select(e.target.parentElement)
            .classed("expand");

          d3.selectAll(".blip-list__item-container.expand").classed(
            "expand",
            false
          );
          d3.select(e.target.parentElement).classed("expand", !expandFlag);

          d3.selectAll(".blip-list__item-container__name").attr(
            "aria-expanded",
            "false"
          );
          d3.select(
            ".blip-list__item-container.expand .blip-list__item-container__name"
          ).attr("aria-expanded", "true");

          if (window.innerWidth >= uiConfig.tabletViewWidth) {
            stickQuadrantOnScroll();
          }
        });

      blipItemContainer
        .append("span")
        .classed("blip-list__item-container__name-value", true)
        .text(`${blip.blipText()}. ${blip.name()}`);
      blipItemContainer
        .append("span")
        .classed("blip-list__item-container__name-arrow", true);

      blipItemDiv
        .append("div")
        .classed("blip-list__item-container__description", true)
        .attr("id", `blip-description-${blip.id()}`)
        .html(blip.description());
    }
    const blipGraphItem = d3.select(
      `g a#blip-link-${removeAllSpaces(blip.id())}`
    );
    const mouseOver = function (e) {
      const targetElement = e.target.classList.contains("blip-link")
        ? e.target
        : e.target.parentElement;
      const isGroupIdInGraph = !targetElement.classList.contains("blip-link")
        ? true
        : false;
      const blipWrapper = d3.select(targetElement);
      const blipIdToFocus = blip.groupIdInGraph()
        ? blipWrapper.attr("data-group-id")
        : blipWrapper.attr("data-blip-id");
      const selectedBlipOnGraph = d3.select(
        `g > a.blip-link[data-blip-id='${blipIdToFocus}'`
      );
      highlightBlipInGraph(blipIdToFocus);
      highlightBlipInTable(blipTableItem);

      const isQuadrantView = d3
        .select("svg#radar-plot")
        .classed("quadrant-view");
      const displayToolTip = blip.isGroup()
        ? !isQuadrantView
        : !blip.groupIdInGraph();
      const toolTipText = blip.isGroup() ? groupBlipTooltipText : blip.name();

      if (displayToolTip && !isGroupIdInGraph) {
        tip.show(toolTipText, selectedBlipOnGraph.node());

        const selectedBlipCoords = selectedBlipOnGraph
          .node()
          .getBoundingClientRect();

        const tipElement = d3.select("div.d3-tip");
        const tipElementCoords = tipElement.node().getBoundingClientRect();

        tipElement
          .style(
            "left",
            `${parseInt(
              selectedBlipCoords.left +
                window.scrollX -
                tipElementCoords.width / 2 +
                selectedBlipCoords.width / 2
            )}px`
          )
          .style(
            "top",
            `${parseInt(
              selectedBlipCoords.top + window.scrollY - tipElementCoords.height
            )}px`
          );
      }
    };

    const mouseOut = function () {
      d3.selectAll("g > a.blip-link").attr("opacity", 1.0);
      blipTableItem.classed("highlight", false);
      tip.hide().style("left", 0).style("top", 0);
    };

    const blipClick = function (e) {
      const isQuadrantView = d3
        .select("svg#radar-plot")
        .classed("quadrant-view");
      const targetElement = e.target.classList.contains("blip-link")
        ? e.target
        : e.target.parentElement;
      if (isQuadrantView) {
        e.stopPropagation();
      }

      const blipId = d3.select(targetElement).attr("data-blip-id");
      highlightBlipInGraph(blipId);

      d3.selectAll(".blip-list__item-container.expand").classed(
        "expand",
        false
      );

      let selectedBlipContainer = d3.select(
        `.blip-list__item-container[data-blip-id="${blipId}"`
      );
      selectedBlipContainer.classed("expand", true);

      setTimeout(
        () => {
          if (window.innerWidth >= uiConfig.tabletViewWidth) {
            stickQuadrantOnScroll();
          }

          const isGroupBlip = isNaN(parseInt(blipId));
          if (isGroupBlip) {
            selectedBlipContainer = d3.select(
              `.blip-list__item-container[data-group-id="${blipId}"`
            );
          }
          const elementToFocus = selectedBlipContainer.select(
            "button.blip-list__item-container__name"
          );
          elementToFocus.node()?.scrollIntoView({
            behavior: "smooth",
          });
        },
        isQuadrantView ? 0 : 1500
      );
    };

    !groupBlipTooltipText &&
      blipTableItem
        .on("mouseover", mouseOver)
        .on("mouseout", mouseOut)
        .on("focusin", mouseOver)
        .on("focusout", mouseOut);
    blipGraphItem
      .on("mouseover", mouseOver)
      .on("mouseout", mouseOut)
      .on("focusin", mouseOver)
      .on("focusout", mouseOut)
      .on("click", blipClick);
  }

  function renderQuadrantTables(quadrants, rings) {
    const radarContainer = d3.select("#radar");

    const quadrantTablesContainer = radarContainer
      .append("div")
      .classed("quadrant-table__container", true);
    quadrants.forEach(function (quadrant) {
      const scale = getScale();
      let quadrantContainer;
      if (
        window.innerWidth < uiConfig.tabletViewWidth &&
        window.innerWidth >= uiConfig.mobileViewWidth
      ) {
        quadrantContainer = quadrantTablesContainer
          .append("div")
          .classed("quadrant-table", true)
          .classed(quadrant.order, true)
          .style(
            "margin",
            `${
              graphConfig.quadrantHeight * scale +
              graphConfig.quadrantsGap * scale +
              graphConfig.quadrantsGap * 2 +
              uiConfig.legendsHeight
            }px auto 0px`
          )
          .style("left", "0")
          .style("right", 0);
      } else {
        quadrantContainer = quadrantTablesContainer
          .append("div")
          .classed("quadrant-table", true)
          .classed(quadrant.order, true);
      }

      const ringNames = Array.from(
        new Set(
          quadrant.quadrant
            .blips()
            .map((blip) => blip.ring())
            .map((ring) => ring.name())
        )
      );
      ringNames.forEach(function (ringName) {
        quadrantContainer
          .append("h2")
          .classed("quadrant-table__ring-name", true)
          .attr("data-ring-name", ringName)
          .text(ringName);
        quadrantContainer
          .append("ul")
          .classed("blip-list", true)
          .attr(
            "data-ring-order",
            rings.filter((ring) => ring.name() === ringName)[0].order()
          );
      });
    });
  }

  /**
   * GRAPHING: js/graphing/blips.js
   */
  const featureToggles = mainConfig().featureToggles;

  const getRingRadius = function (ringIndex) {
    const ratios = [0, 0.316, 0.652, 0.832, 0.992];
    const radius = ratios[ringIndex] * graphConfig.quadrantWidth;
    return radius || 0;
  };

  function getBorderWidthOffset(quadrantOrder, adjustY, adjustX) {
    let borderWidthYOffset = 0,
      borderWidthXOffset = 0;

    if (quadrantOrder !== "first") {
      borderWidthYOffset = adjustY < 0 ? 0 : graphConfig.quadrantsGap;
      borderWidthXOffset = adjustX > 0 ? graphConfig.quadrantsGap : 0;
    }
    return { borderWidthYOffset, borderWidthXOffset };
  }

  function calculateRadarBlipCoordinates(
    minRadius,
    maxRadius,
    startAngle,
    quadrantOrder,
    chance,
    blip
  ) {
    const adjustX =
      Math.sin(toRadian(startAngle)) - Math.cos(toRadian(startAngle));
    const adjustY =
      -Math.cos(toRadian(startAngle)) - Math.sin(toRadian(startAngle));
    const { borderWidthYOffset, borderWidthXOffset } = getBorderWidthOffset(
      quadrantOrder,
      adjustY,
      adjustX
    );
    const radius = chance.floating({
      min: minRadius + blip.width / 2,
      max: maxRadius - blip.width,
    });

    let angleDelta =
      (Math.asin(blip.width / 2 / radius) * 180) / (Math.PI - 1.25);
    angleDelta = angleDelta > 45 ? 45 : angleDelta;
    const angle = toRadian(
      chance.integer({ min: angleDelta, max: 90 - angleDelta })
    );

    let x =
      graphConfig.quadrantWidth +
      radius * Math.cos(angle) * adjustX +
      borderWidthXOffset;
    let y =
      graphConfig.quadrantHeight +
      radius * Math.sin(angle) * adjustY +
      borderWidthYOffset;

    return avoidBoundaryCollision(x, y, adjustX, adjustY);
  }

  function thereIsCollision(coordinates, allCoordinates, blipWidth) {
    return allCoordinates.some(function (currentCoordinates) {
      return (
        Math.abs(currentCoordinates.coordinates[0] - coordinates[0]) <
          currentCoordinates.width / 2 + blipWidth / 2 + 10 &&
        Math.abs(currentCoordinates.coordinates[1] - coordinates[1]) <
          currentCoordinates.width / 2 + blipWidth / 2 + 10
      );
    });
  }

  function avoidBoundaryCollision(x, y, adjustX, adjustY) {
    const size = graphConfig.quadrantWidth * 2 + graphConfig.quadrantsGap;
    if (
      (adjustY > 0 && y + graphConfig.blipWidth > size) ||
      (adjustY < 0 && y + graphConfig.blipWidth > graphConfig.quadrantHeight)
    ) {
      y = y - graphConfig.blipWidth;
    }
    if (adjustX < 0 && x - graphConfig.blipWidth > graphConfig.quadrantWidth) {
      x += graphConfig.blipWidth;
    }
    if (
      adjustX > 0 &&
      x + graphConfig.blipWidth <
        graphConfig.quadrantWidth + graphConfig.quadrantsGap
    ) {
      x -= graphConfig.blipWidth;
    }
    return [x, y];
  }

  function findBlipCoordinates(
    blip,
    minRadius,
    maxRadius,
    startAngle,
    allBlipCoordinatesInRing,
    quadrantOrder
  ) {
    const maxIterations = 200;
    const chance = new Chance(
      Math.PI *
        graphConfig.quadrantWidth *
        graphConfig.quadrantHeight *
        graphConfig.quadrantsGap *
        graphConfig.blipWidth *
        maxIterations
    );
    let coordinates = calculateRadarBlipCoordinates(
      minRadius,
      maxRadius,
      startAngle,
      quadrantOrder,
      chance,
      blip
    );
    let iterationCounter = 0;
    let foundAPlace = false;

    while (iterationCounter < maxIterations) {
      if (thereIsCollision(coordinates, allBlipCoordinatesInRing, blip.width)) {
        coordinates = calculateRadarBlipCoordinates(
          minRadius,
          maxRadius,
          startAngle,
          quadrantOrder,
          chance,
          blip
        );
      } else {
        foundAPlace = true;
        break;
      }
      iterationCounter++;
    }
    if (
      !featureToggles.UIRefresh2022 &&
      !foundAPlace &&
      blip.width > graphConfig.minBlipWidth
    ) {
      blip.width = blip.width - 1;
      blip.scale = Math.max((blip.scale || 1) - 0.1, 0.7);
      return findBlipCoordinates(
        blip,
        minRadius,
        maxRadius,
        startAngle,
        allBlipCoordinatesInRing,
        quadrantOrder
      );
    } else {
      return coordinates;
    }
  }

  function blipAssistiveText(blip) {
    return blip.isGroup()
      ? `\`${blip.ring().name()} ring, group of ${blip.blipText()}`
      : `${blip.ring().name()} ring, ${blip.name()}, ${
          blip.isNew() ? "new" : "existing"
        } blip.`;
  }
  function addOuterCircle(parentSvg, order, scale = 1) {
    parentSvg
      .append("path")
      .attr("opacity", "1")
      .attr("class", order)
      .attr(
        "d",
        "M18 36C8.07 36 0 27.93 0 18S8.07 0 18 0c9.92 0 18 8.07 18 18S27.93 36 18 36zM18 3.14C9.81 3.14 3.14 9.81 3.14 18S9.81 32.86 18 32.86S32.86 26.19 32.86 18S26.19 3.14 18 3.14z"
      )
      .style("transform", `scale(${scale})`);
  }

  function drawBlipCircle(group, blip, xValue, yValue, order) {
    group
      .attr("transform", `scale(1) translate(${xValue - 16}, ${yValue - 16})`)
      .attr("aria-label", blipAssistiveText(blip));
    group
      .append("circle")
      .attr("r", "12")
      .attr("cx", "18")
      .attr("cy", "18")
      .attr("class", order)
      .style("transform", `scale(${blip.scale || 1})`);
  }

  function newBlip(blip, xValue, yValue, order, group) {
    drawBlipCircle(group, blip, xValue, yValue, order);
    addOuterCircle(group, order, blip.scale);
  }

  function existingBlip(blip, xValue, yValue, order, group) {
    drawBlipCircle(group, blip, xValue, yValue, order);
  }

  function groupBlip(blip, xValue, yValue, order, group) {
    group
      .attr("transform", `scale(1) translate(${xValue}, ${yValue})`)
      .attr("aria-label", blipAssistiveText(blip));
    group
      .append("rect")
      .attr("x", "1")
      .attr("y", "1")
      .attr("rx", "12")
      .attr("ry", "12")
      .attr("width", blip.groupBlipWidth())
      .attr("height", graphConfig.groupBlipHeight)
      .attr("class", order)
      .style("transform", `scale(${blip.scale || 1})`);
  }

  function drawBlipInCoordinates(blip, coordinates, order, quadrantGroup) {
    let x = coordinates[0];
    let y = coordinates[1];

    const blipId = removeAllSpaces(blip.id());

    const group = quadrantGroup
      .append("g")
      .append("a")
      .attr("href", "javascript:void(0)")
      .attr("class", "blip-link")
      .attr("id", "blip-link-" + blipId)
      .attr("data-blip-id", blipId)
      .attr("data-ring-name", blip.ring().name());

    if (blip.isGroup()) {
      groupBlip(blip, x, y, order, group);
    } else if (blip.isNew()) {
      newBlip(blip, x, y, order, group);
    } else {
      existingBlip(blip, x, y, order, group);
    }

    group
      .append("text")
      .attr("x", blip.isGroup() ? (blip.isNew() ? 45 : 64) : 18)
      .attr("y", blip.isGroup() ? 17 : 23)
      .style("font-size", "12px")
      .attr("font-style", "normal")
      .attr("font-weight", "bold")
      .attr("fill", "white")
      .text(blip.blipText())
      .style("text-anchor", "middle")
      .style("transform", `scale(${blip.scale || 1})`);
  }

  function getGroupBlipTooltipText(ringBlips) {
    let tooltipText = "Click to view all";
    if (ringBlips.length <= 15) {
      tooltipText = ringBlips.reduce((toolTip, blip) => {
        toolTip += blip.id() + ". " + blip.name() + "</br>";
        return toolTip;
      }, "");
    }
    return tooltipText;
  }

  const findExistingBlipCoords = function (ringIndex, deg) {
    const blipWidth = graphConfig.existingGroupBlipWidth;
    const ringWidth = getRingRadius(ringIndex) - getRingRadius(ringIndex - 1);
    const halfRingRadius = getRingRadius(ringIndex) - ringWidth / 2;
    const x =
      graphConfig.quadrantWidth -
      halfRingRadius * Math.cos(toRadian(deg)) -
      blipWidth / 2;
    const y =
      graphConfig.quadrantHeight - halfRingRadius * Math.sin(toRadian(deg));
    return [x, y];
  };

  function findNewBlipCoords(existingCoords) {
    const groupBlipGap = 5;
    const offsetX =
      graphConfig.existingGroupBlipWidth - graphConfig.newGroupBlipWidth;
    const offsetY = graphConfig.groupBlipHeight + groupBlipGap;
    return [existingCoords[0] + offsetX, existingCoords[1] - offsetY];
  }

  const groupBlipsBaseCoords = function (ringIndex) {
    const existingCoords = findExistingBlipCoords(
      ringIndex + 1,
      graphConfig.groupBlipAngles[ringIndex]
    );

    return {
      existing: existingCoords,
      new: findNewBlipCoords(existingCoords),
    };
  };

  const transposeQuadrantCoords = function (coords, blipWidth) {
    const transposeX =
      graphConfig.effectiveQuadrantWidth * 2 - coords[0] - blipWidth;
    const transposeY =
      graphConfig.effectiveQuadrantHeight * 2 -
      coords[1] -
      graphConfig.groupBlipHeight;
    return {
      first: coords,
      second: [coords[0], transposeY],
      third: [transposeX, coords[1]],
      fourth: [transposeX, transposeY],
    };
  };

  function createGroupBlip(blipsInRing, blipType, ring, quadrantOrder) {
    const blipText = `${blipsInRing.length} ${blipType} blips`;
    const blipId = `${quadrantOrder}-${replaceSpaceWithHyphens(
      ring.name()
    )}-group-${replaceSpaceWithHyphens(blipType)}-blips`;
    const groupBlip = new Blip(blipText, ring, blipsInRing[0].isNew(), "", "");
    groupBlip.setBlipText(blipText);
    groupBlip.setId(blipId);
    groupBlip.setIsGroup(true);
    return groupBlip;
  }

  function plotGroupBlips(
    ringBlips,
    ring,
    quadrantOrder,
    parentElement,
    quadrantWrapper,
    tooltip
  ) {
    let newBlipsInRing = [],
      existingBlipsInRing = [];
    ringBlips.forEach((blip) => {
      blip.isNew() ? newBlipsInRing.push(blip) : existingBlipsInRing.push(blip);
    });

    const blipGroups = [newBlipsInRing, existingBlipsInRing].filter(
      (group) => !_.isEmpty(group)
    );
    blipGroups.forEach((blipsInRing) => {
      const blipType = blipsInRing[0].isNew() ? "new" : "existing";
      const groupBlip = createGroupBlip(
        blipsInRing,
        blipType,
        ring,
        quadrantOrder
      );
      const groupBlipTooltipText = getGroupBlipTooltipText(blipsInRing);
      const ringIndex = graphConfig.rings.indexOf(ring.name());
      const baseCoords = groupBlipsBaseCoords(ringIndex)[blipType];
      const blipCoordsForCurrentQuadrant = transposeQuadrantCoords(
        baseCoords,
        groupBlip.groupBlipWidth()
      )[quadrantOrder];
      drawBlipInCoordinates(
        groupBlip,
        blipCoordsForCurrentQuadrant,
        quadrantOrder,
        parentElement
      );
      renderBlipDescription(
        groupBlip,
        ring,
        quadrantWrapper,
        tooltip,
        groupBlipTooltipText
      );
      blipsInRing.forEach(function (blip) {
        blip.setGroupIdInGraph(groupBlip.id());
        renderBlipDescription(blip, ring, quadrantWrapper, tooltip);
      });
    });
  }

  const plotRadarBlips = function (
    parentElement,
    rings,
    quadrantWrapper,
    tooltip
  ) {
    let blips, quadrant, startAngle, quadrantOrder;

    quadrant = quadrantWrapper.quadrant;
    startAngle = quadrantWrapper.startAngle;
    quadrantOrder = quadrantWrapper.order;

    blips = quadrant.blips();
    rings.forEach(function (ring, i) {
      const ringBlips = blips.filter(function (blip) {
        return blip.ring() === ring;
      });

      if (ringBlips.length === 0) {
        return;
      }

      const offset = 10;
      const minRadius = getRingRadius(i) + offset;
      const maxRadius = getRingRadius(i + 1) - offset;
      const allBlipCoordsInRing = [];

      if (ringBlips.length > graphConfig.maxBlipsInRings[i]) {
        plotGroupBlips(
          ringBlips,
          ring,
          quadrantOrder,
          parentElement,
          quadrantWrapper,
          tooltip
        );
        return;
      }

      ringBlips.forEach(function (blip) {
        const coordinates = findBlipCoordinates(
          blip,
          minRadius,
          maxRadius,
          startAngle,
          allBlipCoordsInRing,
          quadrantOrder
        );
        allBlipCoordsInRing.push({ coordinates, width: blip.width });
        drawBlipInCoordinates(blip, coordinates, quadrantOrder, parentElement);
        renderBlipDescription(blip, ring, quadrantWrapper, tooltip);
      });
    });
  };

  /**
   * UTIL: js/util/ringCalculator.js
   */
  const RingCalculator = function (numberOfRings, maxRadius) {
    var sequence = [0, 6, 5, 3, 2, 1, 1, 1];

    var self = {};

    self.sum = function (length) {
      return sequence.slice(0, length + 1).reduce(function (previous, current) {
        return previous + current;
      }, 0);
    };

    self.getRadius = function (ring) {
      var total = self.sum(numberOfRings);
      var sum = self.sum(ring);

      return (maxRadius * sum) / total;
    };

    self.getRingRadius = function (ringIndex) {
      const ratios = [0, 0.316, 0.652, 0.832, 1];
      const radius = ratios[ringIndex] * maxRadius;
      return radius || 0;
    };

    return self;
  };

  /**
   * COMPONENTS: js/graphing/components/quadrantSubnav.js
   */
  function addListItem(quadrantList, name, callback) {
    quadrantList
      .append("li")
      .attr("id", `subnav-item-${getRingIdString(name)}`)
      .classed("quadrant-subnav__list-item", true)
      .attr("title", name)
      .append("button")
      .classed("quadrant-subnav__list-item__button", true)
      .attr("role", "tab")
      .text(name)
      .on("click", function (e) {
        d3.select("#radar").classed("no-blips", false);
        d3.select("#auto-complete").property("value", "");
        removeScrollListener();

        d3.select(".graph-header").node().scrollIntoView({
          behavior: "smooth",
        });

        d3.select("span.quadrant-subnav__dropdown-selector").text(
          e.target.innerText
        );

        const subnavArrow = d3.select(".quadrant-subnav__dropdown-arrow");
        subnavArrow.classed(
          "rotate",
          !d3.select("span.quadrant-subnav__dropdown-arrow").classed("rotate")
        );
        quadrantList.classed(
          "show",
          !d3.select("ul.quadrant-subnav__list").classed("show")
        );

        const subnavDropdown = d3.select(".quadrant-subnav__dropdown");
        subnavDropdown.attr(
          "aria-expanded",
          subnavDropdown.attr("aria-expanded") === "false" ? "true" : "false"
        );

        d3.selectAll(".blip-list__item-container.expand").classed(
          "expand",
          false
        );

        if (callback) {
          callback();
        }
      });
  }

  function renderQuadrantSubnav(radarHeader, quadrants, renderFullRadar) {
    const subnavContainer = radarHeader
      .append("nav")
      .classed("quadrant-subnav", true);

    const subnavDropdown = subnavContainer
      .append("div")
      .classed("quadrant-subnav__dropdown", true)
      .attr("aria-expanded", "false");
    subnavDropdown
      .append("span")
      .classed("quadrant-subnav__dropdown-selector", true)
      .text("All quadrants");
    const subnavArrow = subnavDropdown
      .append("span")
      .classed("quadrant-subnav__dropdown-arrow", true);

    const quadrantList = subnavContainer
      .append("ul")
      .classed("quadrant-subnav__list", true);
    addListItem(quadrantList, "All quadrants", renderFullRadar);
    d3.select("li.quadrant-subnav__list-item")
      .classed("active-item", true)
      .select("button")
      .attr("aria-selected", "true");

    subnavDropdown.on("click", function () {
      subnavArrow.classed(
        "rotate",
        !d3.select("span.quadrant-subnav__dropdown-arrow").classed("rotate")
      );
      quadrantList.classed(
        "show",
        !d3.select("ul.quadrant-subnav__list").classed("show")
      );

      subnavDropdown.attr(
        "aria-expanded",
        subnavDropdown.attr("aria-expanded") === "false" ? "true" : "false"
      );
    });

    quadrants.forEach(function (quadrant) {
      addListItem(quadrantList, quadrant.quadrant.name(), () =>
        selectRadarQuadrant(
          quadrant.order,
          quadrant.startAngle,
          quadrant.quadrant.name()
        )
      );
    });

    const subnavOffset =
      (window.innerWidth < 1024
        ? uiConfig.tabletBannerHeight
        : uiConfig.bannerHeight) + uiConfig.headerHeight;

    window.addEventListener("scroll", function () {
      if (subnavOffset <= window.scrollY) {
        d3.select(".quadrant-subnav").classed("sticky", true);
        d3.select(".search-container").classed("sticky-offset", true);
      } else {
        d3.select(".quadrant-subnav").classed("sticky", false);
        d3.select(".search-container").classed("sticky-offset", false);
      }
    });
  }

  /**
   * UTIL: js/util/autoComplete.js
   */
  const featureToggles2 = mainConfig().featureToggles;
  $.widget("custom.radarcomplete", $.ui.autocomplete, {
    _create: function () {
      this._super();
      console.log("Custom radarcomplete widget created!");
      this.widget().menu(
        "option",
        "items",
        "> :not(.ui-autocomplete-quadrant)"
      );
    },
    _renderMenu: function (ul, items) {
      let currentQuadrant = "";

      items.forEach((item) => {
        const quadrantName = item.quadrant.quadrant.name();
        if (quadrantName !== currentQuadrant) {
          ul.append(
            `<li class='ui-autocomplete-quadrant'>${quadrantName}</li>`
          );
          currentQuadrant = quadrantName;
        }
        const li = this._renderItemData(ul, item);
        if (quadrantName) {
          li.attr("aria-label", `${quadrantName}:${item.value}`);
        }
      });
    },
  });

  const AutoComplete1 = (el, quadrants, cb) => {
    const blips = quadrants.reduce((acc, quadrant) => {
      return [
        ...acc,
        ...quadrant.quadrant.blips().map((blip) => ({ blip, quadrant })),
      ];
    }, []);

    if (featureToggles2.UIRefresh2022) {
      $(el).autocomplete({
        appendTo: ".search-container",
        source: (request, response) => {
          const matches = blips.filter(({ blip }) => {
            const searchable =
              `${blip.name()} ${blip.description()}`.toLowerCase();
            return request.term
              .split(" ")
              .every((term) => searchable.includes(term.toLowerCase()));
          });
          response(
            matches.map((item) => ({ ...item, value: item.blip.name() }))
          );
        },
        select: cb.bind({}),
      });
    } else {
      $(el).radarcomplete({
        source: (request, response) => {
          const matches = blips.filter(({ blip }) => {
            const searchable =
              `${blip.name()} ${blip.description()}`.toLowerCase();
            return request.term
              .split(" ")
              .every((term) => searchable.includes(term.toLowerCase()));
          });
          response(
            matches.map((item) => ({ ...item, value: item.blip.name() }))
          );
        },
        select: cb.bind({}),
      });
    }
  };

  /**
   * COMPONENTS: js/graphing/components/search.js
   */
  const AutoCompleteSearch = AutoComplete1;

  function renderSearch(radarHeader, quadrants) {
    const searchContainer = radarHeader
      .append("div")
      .classed("search-container", true);

    searchContainer
      .append("input")
      .classed("search-container__input", true)
      .attr("placeholder", "Search this radar")
      .attr("id", "auto-complete");

    AutoCompleteSearch("#auto-complete", quadrants, function (e, ui) {
      const blipId = ui.item.blip.id();
      const quadrant = ui.item.quadrant;

      selectRadarQuadrant(
        quadrant.order,
        quadrant.startAngle,
        quadrant.quadrant.name()
      );
      const blipElement = d3.select(
        `.blip-list__item-container[data-blip-id="${blipId}"] .blip-list__item-container__name`
      );

      removeScrollListener();
      blipElement.dispatch("search-result-click");

      setTimeout(() => {
        blipElement.node().scrollIntoView({
          behavior: "smooth",
        });
      }, 1500);
    });
  }

  /**
   * COMPONENTS: js/graphing/components/buttons.js
   */
  function renderButtons(radarFooter) {
    const buttonsRow = radarFooter.append("div").classed("buttons", true);

    buttonsRow
      .append("button")
      .classed("buttons__wave-btn", true)
      .text("Print this Radar")
      .on("click", window.print.bind(window));
  }

  /**
   * GRAPHING: js/graphing/pdfPage.js
   */
  const addPdfCoverTitle = (title) => {
    d3.select("main #pdf-cover-page .pdf-title").text(title);
  };

  const addRadarLinkInPdfView = () => {
    d3.select("#generated-radar-link").attr("href", window.location.href);
  };

  const addQuadrantNameInPdfView = (order, quadrantName) => {
    d3.select(`.quadrant-table.${order}`)
      .insert("div", ":first-child")
      .attr("class", "quadrant-table__name")
      .text(quadrantName);
  };

  /**
   *  UTIL: js/util/urlUtils.js
   */
  function constructSheetUrl(sheetName) {
    const noParamsUrl = window.location.href.substring(
      0,
      window.location.href.indexOf(window.location.search)
    );
    const queryParams = QueryParams(window.location.search.substring(1));
    const sheetUrl =
      noParamsUrl +
      "?" +
      ((queryParams.documentId &&
        `documentId=${encodeURIComponent(queryParams.documentId)}`) ||
        (queryParams.sheetId &&
          `sheetId=${encodeURIComponent(queryParams.sheetId)}`) ||
        "") +
      "&sheetName=" +
      encodeURIComponent(sheetName);
    return sheetUrl;
  }

  function getDocumentOrSheetId() {
    const queryParams = QueryParams(window.location.search.substring(1));
    return queryParams.documentId ?? queryParams.sheetId;
  }

  function getSheetName() {
    const queryParams = QueryParams(window.location.search.substring(1));
    return queryParams.sheetName;
  }

  /**
   * COMPONENTS: js/graphing/components/banner.js
   */
  function renderBanner(renderFullRadar) {
    if (featureToggles.UIRefresh2022) {
      const documentTitle =
        document.title[0].toUpperCase() + document.title.slice(1);

      document.title = documentTitle;
      d3.select(".hero-banner__wrapper")
        .append("p")
        .classed("hero-banner__subtitle-text", true)
        .text(document.title);
      d3.select(".hero-banner__title-text").on("click", renderFullRadar);

      addPdfCoverTitle(documentTitle);
    } else {
      const header = d3.select("body").insert("header", "#radar");
      header
        .append("div")
        .attr("class", "radar-title")
        .append("div")
        .attr("class", "radar-title__text")
        .append("h1")
        .text(document.title)
        .style("cursor", "pointer")
        .on("click", renderFullRadar);

      header
        .select(".radar-title")
        .append("div")
        .attr("class", "radar-title__logo")
        .html('<img src="../images/logo.png" />');
    }
  }

  /** 
   * UTIL: js/util/queryParamProcessor.js
   */
  const QueryParams = function (queryString) {
    var decode = function (s) {
      return decodeURIComponent(s.replace(/\+/g, " "));
    };

    var search = /([^&=]+)=?([^&]*)/g;

    var queryParams = {};
    var match;
    while ((match = search.exec(queryString))) {
      queryParams[decode(match[1])] = decode(match[2]);
    }

    return queryParams;
  };

  /**
   * GRAPHING: js/graphing/radar.js
   */
  const MIN_BLIP_WIDTH = 12;

  const GraphingRadar = function (size, gRadar) {
    const CENTER = size / 2;
    var svg,
      radarElement,
      quadrantButtons,
      buttonsGroup,
      header,
      alternativeDiv;

    var tip = d3tip()
      .attr("class", "d3-tip")
      .html(function (text) {
        return text;
      });

    tip.direction(function () {
      return "n";
    });

    var ringCalculator = new RingCalculator(gRadar.rings().length, CENTER);

    var self = {};
    var chance;

    function plotLines(quadrantGroup, quadrant) {
      const startX =
        size * (1 - (-Math.sin(toRadian(quadrant.startAngle)) + 1) / 2);
      const endX =
        size * (1 - (-Math.sin(toRadian(quadrant.startAngle - 90)) + 1) / 2);

      let startY =
        size * (1 - (Math.cos(toRadian(quadrant.startAngle)) + 1) / 2);
      let endY =
        size * (1 - (Math.cos(toRadian(quadrant.startAngle - 90)) + 1) / 2);

      if (startY > endY) {
        const aux = endY;
        endY = startY;
        startY = aux;
      }
      const strokeWidth = featureToggles.UIRefresh2022
        ? graphConfig.quadrantsGap
        : 10;

      quadrantGroup
        .append("line")
        .attr("x1", CENTER)
        .attr("y1", startY)
        .attr("x2", CENTER)
        .attr("y2", endY)
        .attr("stroke-width", strokeWidth);

      quadrantGroup
        .append("line")
        .attr("x1", endX)
        .attr("y1", CENTER)
        .attr("x2", startX)
        .attr("y2", CENTER)
        .attr("stroke-width", strokeWidth);
    }

    function plotQuadrant(rings, quadrant) {
      var quadrantGroup = svg
        .append("g")
        .attr("class", "quadrant-group quadrant-group-" + quadrant.order)
        .on("mouseover", mouseoverQuadrant.bind({}, quadrant.order))
        .on("mouseout", mouseoutQuadrant.bind({}, quadrant.order))
        .on(
          "click",
          selectQuadrant.bind({}, quadrant.order, quadrant.startAngle)
        );

      rings.forEach(function (ring, i) {
        var arc = d3
          .arc()
          .innerRadius(ringCalculator.getRadius(i))
          .outerRadius(ringCalculator.getRadius(i + 1))
          .startAngle(toRadian(quadrant.startAngle))
          .endAngle(toRadian(quadrant.startAngle - 90));

        quadrantGroup
          .append("path")
          .attr("d", arc)
          .attr("class", "ring-arc-" + ring.order())
          .attr("transform", "translate(" + CENTER + ", " + CENTER + ")");
      });

      return quadrantGroup;
    }

    function plotTexts(quadrantGroup, rings, quadrant) {
      rings.forEach(function (ring, i) {
        if (quadrant.order === "first" || quadrant.order === "fourth") {
          quadrantGroup
            .append("text")
            .attr("class", "line-text")
            .attr("y", CENTER + 4)
            .attr(
              "x",
              CENTER +
                (ringCalculator.getRadius(i) +
                  ringCalculator.getRadius(i + 1)) /
                  2
            )
            .attr("text-anchor", "middle")
            .text(ring.name());
        } else {
          quadrantGroup
            .append("text")
            .attr("class", "line-text")
            .attr("y", CENTER + 4)
            .attr(
              "x",
              CENTER -
                (ringCalculator.getRadius(i) +
                  ringCalculator.getRadius(i + 1)) /
                  2
            )
            .attr("text-anchor", "middle")
            .text(ring.name());
        }
      });
    }

    function plotRingNames(quadrantGroup, rings, quadrant) {
      rings.forEach(function (ring, i) {
        const ringNameWithEllipsis =
          ring.name().length > 6
            ? ring.name().slice(0, 6) + "..."
            : ring.name();
        if (quadrant.order === "third" || quadrant.order === "fourth") {
          quadrantGroup
            .append("text")
            .attr("class", "line-text")
            .attr("y", CENTER + 5)
            .attr(
              "x",
              CENTER +
                (ringCalculator.getRingRadius(i) +
                  ringCalculator.getRingRadius(i + 1)) /
                  2
            )
            .attr("text-anchor", "middle")
            .text(ringNameWithEllipsis);
        } else {
          quadrantGroup
            .append("text")
            .attr("class", "line-text")
            .attr("y", CENTER + 5)
            .attr(
              "x",
              CENTER -
                (ringCalculator.getRingRadius(i) +
                  ringCalculator.getRingRadius(i + 1)) /
                  2
            )
            .attr("text-anchor", "middle")
            .text(ringNameWithEllipsis);
        }
      });
    }

    function triangle(blip, x, y, order, group) {
      return group
        .append("path")
        .attr(
          "d",
          "M412.201,311.406c0.021,0,0.042,0,0.063,0c0.067,0,0.135,0,0.201,0c4.052,0,6.106-0.051,8.168-0.102c2.053-0.051,4.115-0.102,8.176-0.102h0.103c6.976-0.183,10.227-5.306,6.306-11.53c-3.988-6.121-4.97-5.407-8.598-11.224c-1.631-3.008-3.872-4.577-6.179-4.577c-2.276,0-4.613,1.528-6.48,4.699c-3.578,6.077-3.26,6.014-7.306,11.723C402.598,306.067,405.426,311.406,412.201,311.406"
        )
        .attr(
          "transform",
          "scale(" +
            blip.width / 34 +
            ") translate(" +
            (-404 + x * (34 / blip.width) - 17) +
            ", " +
            (-282 + y * (34 / blip.width) - 17) +
            ")"
        )
        .attr("class", order);
    }

    function triangleLegend(x, y, group) {
      return group
        .append("path")
        .attr(
          "d",
          "M412.201,311.406c0.021,0,0.042,0,0.063,0c0.067,0,0.135,0,0.201,0c4.052,0,6.106-0.051,8.168-0.102c2.053-0.051,4.115-0.102,8.176-0.102h0.103c6.976-0.183,10.227-5.306,6.306-11.53c-3.988-6.121-4.97-5.407-8.598-11.224c-1.631-3.008-3.872-4.577-6.179-4.577c-2.276,0-4.613,1.528-6.48,4.699c-3.578,6.077-3.26,6.014-7.306,11.723C402.598,306.067,405.426,311.406,412.201,311.406"
        )
        .attr(
          "transform",
          "scale(" +
            22 / 64 +
            ") translate(" +
            (-404 + x * (64 / 22) - 17) +
            ", " +
            (-282 + y * (64 / 22) - 17) +
            ")"
        );
    }

    function circle(blip, x, y, order, group) {
      return (group || svg)
        .append("path")
        .attr(
          "d",
          "M420.084,282.092c-1.073,0-2.16,0.103-3.243,0.313c-6.912,1.345-13.188,8.587-11.423,16.874c1.732,8.141,8.632,13.711,17.806,13.711c0.025,0,0.052,0,0.074-0.003c0.551-0.025,1.395-0.011,2.225-0.109c4.404-0.534,8.148-2.218,10.069-6.487c1.747-3.886,2.114-7.993,0.913-12.118C434.379,286.944,427.494,282.092,420.084,282.092"
        )
        .attr(
          "transform",
          "scale(" +
            blip.width / 34 +
            ") translate(" +
            (-404 + x * (34 / blip.width) - 17) +
            ", " +
            (-282 + y * (34 / blip.width) - 17) +
            ")"
        )
        .attr("class", order);
    }

    function circleLegend(x, y, group) {
      return (group || svg)
        .append("path")
        .attr(
          "d",
          "M420.084,282.092c-1.073,0-2.16,0.103-3.243,0.313c-6.912,1.345-13.188,8.587-11.423,16.874c1.732,8.141,8.632,13.711,17.806,13.711c0.025,0,0.052,0,0.074-0.003c0.551-0.025,1.395-0.011,2.225-0.109c4.404-0.534,8.148-2.218,10.069-6.487c1.747-3.886,2.114-7.993,0.913-12.118C434.379,286.944,427.494,282.092,420.084,282.092"
        )
        .attr(
          "transform",
          "scale(" +
            22 / 64 +
            ") translate(" +
            (-404 + x * (64 / 22) - 17) +
            ", " +
            (-282 + y * (64 / 22) - 17) +
            ")"
        );
    }

    function addRing(ring, order) {
      var table = d3.select(".quadrant-table." + order);
      table.append("h3").text(ring);
      return table.append("ul");
    }

    function calculateBlipCoordinates(
      blip,
      chance,
      minRadius,
      maxRadius,
      startAngle
    ) {
      var adjustX =
        Math.sin(toRadian(startAngle)) - Math.cos(toRadian(startAngle));
      var adjustY =
        -Math.cos(toRadian(startAngle)) - Math.sin(toRadian(startAngle));

      var radius = chance.floating({
        min: minRadius + blip.width / 2,
        max: maxRadius - blip.width / 2,
      });
      var angleDelta =
        (Math.asin(blip.width / 2 / radius) * 180) / (Math.PI - 1.25);
      angleDelta = angleDelta > 45 ? 45 : angleDelta;
      var angle = toRadian(
        chance.integer({ min: angleDelta, max: 90 - angleDelta })
      );

      var x = CENTER + radius * Math.cos(angle) * adjustX;
      var y = CENTER + radius * Math.sin(angle) * adjustY;

      return [x, y];
    }

    function thereIsCollision(blip, coordinates, allCoordinates) {
      return allCoordinates.some(function (currentCoordinates) {
        return (
          Math.abs(currentCoordinates[0] - coordinates[0]) < blip.width + 10 &&
          Math.abs(currentCoordinates[1] - coordinates[1]) < blip.width + 10
        );
      });
    }

    function plotBlips(quadrantGroup, rings, quadrantWrapper) {
      var blips, quadrant, startAngle, order;

      quadrant = quadrantWrapper.quadrant;
      startAngle = quadrantWrapper.startAngle;
      order = quadrantWrapper.order;

      d3.select(".quadrant-table." + order)
        .append("h2")
        .attr("class", "quadrant-table__name")
        .text(quadrant.name());

      blips = quadrant.blips();
      rings.forEach(function (ring, i) {
        var ringBlips = blips.filter(function (blip) {
          return blip.ring() === ring;
        });

        if (ringBlips.length === 0) {
          return;
        }

        var maxRadius, minRadius;

        minRadius = ringCalculator.getRadius(i);
        maxRadius = ringCalculator.getRadius(i + 1);

        var sumRing = ring
          .name()
          .split("")
          .reduce(function (p, c) {
            return p + c.charCodeAt(0);
          }, 0);
        var sumQuadrant = quadrant
          .name()
          .split("")
          .reduce(function (p, c) {
            return p + c.charCodeAt(0);
          }, 0);
        chance = new Chance(
          Math.PI *
            sumRing *
            ring.name().length *
            sumQuadrant *
            quadrant.name().length
        );

        var ringList = addRing(ring.name(), order);
        var allBlipCoordinatesInRing = [];

        ringBlips.forEach(function (blip) {
          const coordinates = findBlipCoordinates(
            blip,
            minRadius,
            maxRadius,
            startAngle,
            allBlipCoordinatesInRing,
            order
          );

          allBlipCoordinatesInRing.push(coordinates);
          drawBlipInCoordinates(
            blip,
            coordinates,
            order,
            quadrantGroup,
            ringList
          );
        });
      });
    }

    function findBlipCoordinates(
      blip,
      minRadius,
      maxRadius,
      startAngle,
      allBlipCoordinatesInRing,
      quadrantOrder
    ) {
      const maxIterations = 200;
      var coordinates = calculateBlipCoordinates(
        blip,
        chance,
        minRadius,
        maxRadius,
        startAngle,
        quadrantOrder
      );
      var iterationCounter = 0;
      var foundAPlace = false;

      while (iterationCounter < maxIterations) {
        if (thereIsCollision(blip, coordinates, allBlipCoordinatesInRing)) {
          coordinates = calculateBlipCoordinates(
            blip,
            chance,
            minRadius,
            maxRadius,
            startAngle,
            quadrantOrder
          );
        } else {
          foundAPlace = true;
          break;
        }
        iterationCounter++;
      }

      if (!foundAPlace && blip.width > MIN_BLIP_WIDTH) {
        blip.width = blip.width - 1;
        return findBlipCoordinates(
          blip,
          minRadius,
          maxRadius,
          startAngle,
          allBlipCoordinatesInRing,
          quadrantOrder
        );
      } else {
        return coordinates;
      }
    }

    function drawBlipInCoordinates(
      blip,
      coordinates,
      order,
      quadrantGroup,
      ringList
    ) {
      var x = coordinates[0];
      var y = coordinates[1];

      var group = quadrantGroup
        .append("g")
        .attr("class", "blip-link")
        .attr("id", "blip-link-" + blip.id());

      if (blip.isNew()) {
        triangle(blip, x, y, order, group);
      } else {
        circle(blip, x, y, order, group);
      }
      group
        .append("text")
        .attr("x", x)
        .attr("y", y + 4)
        .attr("class", "blip-text")
        // derive font-size from current blip width
        .style("font-size", (blip.width * 10) / 22 + "px")
        .attr("text-anchor", "middle")
        .text(blip.blipText());

      var blipListItem = ringList.append("li");
      var blipText =
        blip.blipText() +
        ". " +
        blip.name() +
        (blip.topic() ? ". - " + blip.topic() : "");
      blipListItem
        .append("div")
        .attr("class", "blip-list-item")
        .attr("id", "blip-list-item-" + blip.id())
        .text(blipText);

      var blipItemDescription = blipListItem
        .append("div")
        .attr("id", "blip-description-" + blip.id())
        .attr("class", "blip-item-description");
      if (blip.description()) {
        blipItemDescription.append("p").html(blip.description());
      }

      var mouseOver = function () {
        d3.selectAll("g.blip-link").attr("opacity", 0.3);
        group.attr("opacity", 1.0);
        blipListItem.selectAll(".blip-list-item").classed("highlight", true);
        tip.show(blip.name(), group.node());
      };

      var mouseOut = function () {
        d3.selectAll("g.blip-link").attr("opacity", 1.0);
        blipListItem.selectAll(".blip-list-item").classed("highlight", false);
        tip.hide().style("left", 0).style("top", 0);
      };

      blipListItem.on("mouseover", mouseOver).on("mouseout", mouseOut);
      group.on("mouseover", mouseOver).on("mouseout", mouseOut);

      var clickBlip = function () {
        d3.select(".blip-item-description.expanded").node() !==
          blipItemDescription.node() &&
          d3
            .select(".blip-item-description.expanded")
            .classed("expanded", false);
        blipItemDescription.classed(
          "expanded",
          !blipItemDescription.classed("expanded")
        );

        blipItemDescription.on("click", function () {
          d3.event.stopPropagation();
        });
      };

      blipListItem.on("click", clickBlip);
    }

    function removeHomeLink() {
      d3.select(".home-link").remove();
    }

    function createHomeLink(pageElement) {
      if (pageElement.select(".home-link").empty()) {
        pageElement
          .insert("div", "div#alternative-buttons")
          .html("&#171; Back to Radar home")
          .classed("home-link", true)
          .classed("selected", true)
          .on("click", redrawFullRadar)
          .append("g")
          .attr("fill", "#626F87")
          .append("path")
          .attr(
            "d",
            "M27.6904224,13.939279 C27.6904224,13.7179572 27.6039633,13.5456925 27.4314224,13.4230122 L18.9285959,6.85547454 C18.6819796,6.65886965 18.410898,6.65886965 18.115049,6.85547454 L9.90776939,13.4230122 C9.75999592,13.5456925 9.68592041,13.7179572 9.68592041,13.939279 L9.68592041,25.7825947 C9.68592041,25.979501 9.74761224,26.1391059 9.87092041,26.2620876 C9.99415306,26.3851446 10.1419265,26.4467108 10.3145429,26.4467108 L15.1946918,26.4467108 C15.391698,26.4467108 15.5518551,26.3851446 15.6751633,26.2620876 C15.7984714,26.1391059 15.8600878,25.979501 15.8600878,25.7825947 L15.8600878,18.5142424 L21.4794061,18.5142424 L21.4794061,25.7822933 C21.4794061,25.9792749 21.5410224,26.1391059 21.6643306,26.2620876 C21.7876388,26.3851446 21.9477959,26.4467108 22.1448776,26.4467108 L27.024951,26.4467108 C27.2220327,26.4467108 27.3821898,26.3851446 27.505498,26.2620876 C27.6288061,26.1391059 27.6904224,25.9792749 27.6904224,25.7822933 L27.6904224,13.939279 Z M18.4849735,0.0301425662 C21.0234,0.0301425662 23.4202449,0.515814664 25.6755082,1.48753564 C27.9308469,2.45887984 29.8899592,3.77497963 31.5538265,5.43523218 C33.2173918,7.09540937 34.5358755,9.05083299 35.5095796,11.3015031 C36.4829061,13.5518717 36.9699469,15.9439104 36.9699469,18.4774684 C36.9699469,20.1744196 36.748098,21.8101813 36.3044755,23.3844521 C35.860551,24.9584216 35.238498,26.4281731 34.4373347,27.7934053 C33.6362469,29.158336 32.6753041,30.4005112 31.5538265,31.5197047 C30.432349,32.6388982 29.1876388,33.5981853 27.8199224,34.3973401 C26.4519041,35.1968717 24.9791531,35.8176578 23.4016694,36.2606782 C21.8244878,36.7033971 20.1853878,36.9247943 18.4849735,36.9247943 C16.7841816,36.9247943 15.1453837,36.7033971 13.5679755,36.2606782 C11.9904918,35.8176578 10.5180429,35.1968717 9.15002449,34.3973401 C7.78223265,33.5978839 6.53752245,32.6388982 5.41612041,31.5197047 C4.29464286,30.4005112 3.33339796,29.158336 2.53253673,27.7934053 C1.73144898,26.4281731 1.10909388,24.9584216 0.665395918,23.3844521 C0.22184898,21.8101813 0,20.1744196 0,18.4774684 C0,16.7801405 0.22184898,15.1446802 0.665395918,13.5704847 C1.10909388,11.9962138 1.73144898,10.5267637 2.53253673,9.16153157 C3.33339796,7.79652546 4.29464286,6.55435031 5.41612041,5.43523218 C6.53752245,4.3160387 7.78223265,3.35675153 9.15002449,2.55752138 C10.5180429,1.75806517 11.9904918,1.13690224 13.5679755,0.694183299 C15.1453837,0.251464358 16.7841816,0.0301425662 18.4849735,0.0301425662 L18.4849735,0.0301425662 Z"
          );
      }
    }

    function removeRadarLegend() {
      d3.select(".legend").remove();
    }

    function drawLegend(order) {
      removeRadarLegend();

      var triangleKey = "New or moved";
      var circleKey = "No change";

      var container = d3
        .select("svg")
        .append("g")
        .attr("class", "legend legend" + "-" + order);

      var x = 10;
      var y = 10;

      if (order === "first") {
        x = (4 * size) / 5;
        y = (1 * size) / 5;
      }

      if (order === "second") {
        x = (1 * size) / 5 - 15;
        y = (1 * size) / 5 - 20;
      }

      if (order === "third") {
        x = (1 * size) / 5 - 15;
        y = (4 * size) / 5 + 15;
      }

      if (order === "fourth") {
        x = (4 * size) / 5;
        y = (4 * size) / 5;
      }

      d3.select(".legend")
        .attr("class", "legend legend-" + order)
        .transition()
        .style("visibility", "visible");

      triangleLegend(x, y, container);

      container
        .append("text")
        .attr("x", x + 15)
        .attr("y", y + 5)
        .attr("font-size", "0.8em")
        .text(triangleKey);

      circleLegend(x, y + 20, container);

      container
        .append("text")
        .attr("x", x + 15)
        .attr("y", y + 25)
        .attr("font-size", "0.8em")
        .text(circleKey);
    }

    function redrawFullRadar() {
      removeHomeLink();
      removeRadarLegend();
      tip.hide();
      d3.selectAll("g.blip-link").attr("opacity", 1.0);

      svg.style("left", 0).style("right", 0);

      d3.selectAll(".button")
        .classed("selected", false)
        .classed("full-view", true);

      d3.selectAll(".quadrant-table").classed("selected", false);
      d3.selectAll(".home-link").classed("selected", false);

      d3.selectAll(".quadrant-group")
        .transition()
        .duration(ANIMATION_DURATION)
        .attr("transform", "scale(1)");

      if (featureToggles.UIRefresh2022) {
        d3.select("#radar-plot").attr("width", size).attr("height", size);
        d3.selectAll(`.quadrant-bg-images`).each(function () {
          this.classList.remove("hidden");
        });
        d3.selectAll(".quadrant-group").style("display", "block");
      } else {
        d3.selectAll(".quadrant-group .blip-link")
          .transition()
          .duration(ANIMATION_DURATION)
          .attr("transform", "scale(1)");
      }
      d3.selectAll(".quadrant-group").style("pointer-events", "auto");
    }

    function renderFullRadar() {
      removeScrollListener();

      d3.select("#auto-complete").property("value", "");

      window.scrollTo({
        top: 0,
        left: 0,
        behavior: "smooth",
      });

      d3.select("#radar-plot").classed("sticky", false);
      d3.select("#radar-plot").classed("quadrant-view", false);
      d3.select("#radar-plot").classed("enable-transition", true);

      d3.select("#radar-plot").attr("data-quadrant-selected", null);

      const size = getGraphSize();
      d3.select(".home-link").remove();
      d3.select(".legend").remove();
      d3.select("#radar").classed("mobile", false);
      d3.select(".all-quadrants-mobile").classed(
        "show-all-quadrants-mobile",
        true
      );

      d3.select("li.quadrant-subnav__list-item.active-item").classed(
        "active-item",
        false
      );
      d3.select("li.quadrant-subnav__list-item").classed("active-item", true);

      d3.select(".quadrant-subnav__dropdown-selector").text("All quadrants");

      d3tip()
        .attr("class", "d3-tip")
        .html(function (text) {
          return text;
        })
        .hide();

      d3.selectAll("g.blip-link").attr("opacity", 1.0);

      svg
        .style("left", 0)
        .style("right", 0)
        .style("top", 0)
        .attr("transform", "scale(1)")
        .style("transform", "scale(1)");

      d3.selectAll(".button")
        .classed("selected", false)
        .classed("full-view", true);

      d3.selectAll(".quadrant-table").classed("selected", false);
      d3.selectAll(".home-link").classed("selected", false);

      d3.selectAll(".quadrant-group")
        .style("display", "block")
        .transition()
        .duration(ANIMATION_DURATION)
        .style("transform", "scale(1)")
        .style("opacity", "1")
        .attr("transform", "translate(0,0)");

      d3.select("#radar-plot").attr("width", size).attr("height", size);
      d3.select(`svg#radar-plot`).style("padding", "0");

      const radarLegendsContainer = d3.select(".radar-legends");
      radarLegendsContainer.attr("class", "radar-legends");
      radarLegendsContainer.attr("style", null);

      d3.selectAll("svg#radar-plot a")
        .attr("aria-hidden", null)
        .attr("tabindex", null);
      d3.selectAll(".quadrant-table button")
        .attr("aria-hidden", "true")
        .attr("tabindex", -1);
      d3.selectAll(".blip-list__item-container__name").attr(
        "aria-expanded",
        "false"
      );

      d3.selectAll(`.quadrant-group rect:nth-child(2n)`).attr("tabindex", 0);
    }

    function searchBlip(_e, ui) {
      const { blip, quadrant } = ui.item;
      const isQuadrantSelected = d3
        .select("div.button." + quadrant.order)
        .classed("selected");
      selectQuadrant.bind({}, quadrant.order, quadrant.startAngle)();
      const selectedDesc = d3.select("#blip-description-" + blip.id());
      d3.select(".blip-item-description.expanded").node() !==
        selectedDesc.node() &&
        d3.select(".blip-item-description.expanded").classed("expanded", false);
      selectedDesc.classed("expanded", true);

      d3.selectAll("g.blip-link").attr("opacity", 0.3);
      const group = d3.select("#blip-link-" + blip.id());
      group.attr("opacity", 1.0);
      d3.selectAll(".blip-list-item").classed("highlight", false);
      d3.select("#blip-list-item-" + blip.id()).classed("highlight", true);
      if (isQuadrantSelected) {
        tip.show(blip.name(), group.node());
      } else {
        // need to account for the animation time associated with selecting a quadrant
        tip.hide();

        setTimeout(function () {
          tip.show(blip.name(), group.node());
        }, ANIMATION_DURATION);
      }
    }

    function plotRadarHeader() {
      header = d3.select("header");

      buttonsGroup = header.append("div").classed("buttons-group", true);
      alternativeDiv = header.append("div").attr("id", "alternative-buttons");

      quadrantButtons = buttonsGroup
        .append("div")
        .classed("quadrant-btn--group", true);
    }

    function plotQuadrantButtons(quadrants) {
      function addButton(quadrant) {
        radarElement
          .append("div")
          .attr("class", "quadrant-table " + quadrant.order);

        quadrantButtons
          .append("div")
          .attr("class", "button " + quadrant.order + " full-view")
          .text(quadrant.quadrant.name())
          .on("mouseover", mouseoverQuadrant.bind({}, quadrant.order))
          .on("mouseout", mouseoutQuadrant.bind({}, quadrant.order))
          .on(
            "click",
            selectQuadrant.bind({}, quadrant.order, quadrant.startAngle)
          );
      }

      _.each([0, 1, 2, 3], function (i) {
        addButton(quadrants[i]);
      });

      buttonsGroup
        .append("div")
        .classed("print-radar-btn", true)
        .append("div")
        .classed("print-radar button no-capitalize", true)
        .text("Print this radar")
        .on("click", window.print.bind(window));

      alternativeDiv
        .append("div")
        .classed("search-box", true)
        .append("input")
        .attr("id", "auto-complete")
        .attr("placeholder", "Search")
        .classed("search-radar", true);

      AutoComplete("#auto-complete", quadrants, searchBlip);
    }

    function plotRadarFooter() {
      d3.select("body")
        .insert("div", "#radar-plot + *")
        .attr("id", "footer")
        .append("div")
        .attr("class", "footer-content")
        .append("p")
        .html("Something something: in radar.js");
    }

    function mouseoverQuadrant(order) {
      d3.select(".quadrant-group-" + order).style("opacity", 1);
      d3.selectAll(".quadrant-group:not(.quadrant-group-" + order + ")").style(
        "opacity",
        0.3
      );
    }

    function mouseoutQuadrant(order) {
      d3.selectAll(".quadrant-group:not(.quadrant-group-" + order + ")").style(
        "opacity",
        1
      );
    }

    function hideTooltipOnScroll(tip) {
      window.addEventListener("scroll", () => {
        tip.hide().style("left", 0).style("top", 0);
      });
    }

    function selectQuadrant(order, startAngle) {
      d3.selectAll(".home-link").classed("selected", false);
      createHomeLink(d3.select("header"));

      d3.selectAll(".button")
        .classed("selected", false)
        .classed("full-view", false);
      d3.selectAll(".button." + order).classed("selected", true);
      d3.selectAll(".quadrant-table").classed("selected", false);
      d3.selectAll(".quadrant-table." + order).classed("selected", true);
      d3.selectAll(".blip-item-description").classed("expanded", false);

      var scale = 2;

      var adjustX =
        Math.sin(toRadian(startAngle)) - Math.cos(toRadian(startAngle));
      var adjustY =
        Math.cos(toRadian(startAngle)) + Math.sin(toRadian(startAngle));

      var translateX =
        ((-1 * (1 + adjustX) * size) / 2) * (scale - 1) +
        -adjustX * (1 - scale / 2) * size;
      var translateY =
        -1 * (1 - adjustY) * (size / 2 - 7) * (scale - 1) -
        ((1 - adjustY) / 2) * (1 - scale / 2) * size;
      if (featureToggles.UIRefresh2022) {
        translateY = 0;
      }

      var translateXAll =
        (((1 - adjustX) / 2) * size * scale) / 2 +
        ((1 - adjustX) / 2) * (1 - scale / 2) * size;
      var translateYAll = (((1 + adjustY) / 2) * size * scale) / 2;

      var moveRight = ((1 + adjustX) * (0.8 * window.innerWidth - size)) / 2;
      var moveLeft = ((1 - adjustX) * (0.8 * window.innerWidth - size)) / 2;

      var blipScale = 3 / 4;
      var blipTranslate = (1 - blipScale) / blipScale;

      svg.style("left", moveLeft + "px").style("right", moveRight + "px");

      d3.select(".quadrant-group-" + order)
        .transition()
        .duration(ANIMATION_DURATION)
        .attr(
          "transform",
          "translate(" + translateX + "," + translateY + ")scale(" + scale + ")"
        );
      d3.selectAll(".quadrant-group-" + order + " .blip-link text").each(
        function () {
          var x = d3.select(this).attr("x");
          var y = d3.select(this).attr("y");
          d3.select(this.parentNode)
            .transition()
            .duration(ANIMATION_DURATION)
            .attr(
              "transform",
              "scale(" +
                blipScale +
                ")translate(" +
                blipTranslate * x +
                "," +
                blipTranslate * y +
                ")"
            );
        }
      );

      d3.selectAll(".quadrant-group").style("pointer-events", "auto");

      d3.selectAll(".quadrant-group:not(.quadrant-group-" + order + ")")
        .transition()
        .duration(ANIMATION_DURATION)
        .style("pointer-events", "none")
        .attr(
          "transform",
          "translate(" + translateXAll + "," + translateYAll + ")scale(0)"
        );

      if (d3.select(".legend.legend-" + order).empty()) {
        drawLegend(order);
      }
    }

    self.init = function () {
      radarElement = d3.select("#radar");
      if (!featureToggles.UIRefresh2022) {
        const selector = "body";
        radarElement = d3.select(selector).append("div").attr("id", "radar");
      }
      return self;
    };

    function plotAlternativeRadars(alternatives, currentSheet) {
      var alternativeSheetButton = alternativeDiv
        .append("div")
        .classed("multiple-sheet-button-group", true);

      alternativeSheetButton
        .append("p")
        .text("Choose a sheet to populate radar");
      alternatives.forEach(function (alternative) {
        alternativeSheetButton
          .append("div:a")
          .attr("class", "first full-view alternative multiple-sheet-button")
          .attr("href", constructSheetUrl(alternative))
          .text(alternative);

        if (alternative === currentSheet) {
          d3.selectAll(".alternative")
            .filter(function () {
              return d3.select(this).text() === alternative;
            })
            .attr("class", "highlight multiple-sheet-button");
        }
      });
    }

    self.plot = function () {
      var rings, quadrants, alternatives, currentSheet;

      rings = gRadar.rings();
      quadrants = gRadar.quadrants();
      alternatives = gRadar.getAlternatives();
      currentSheet = gRadar.getCurrentSheet();

      const radarHeader = d3.select("main .graph-header");
      const radarFooter = d3.select("main .graph-footer");

      renderBanner(renderFullRadar);

      if (featureToggles.UIRefresh2022) {
        renderQuadrantSubnav(radarHeader, quadrants, renderFullRadar);
        renderSearch(radarHeader, quadrants);
        renderQuadrantTables(quadrants, rings);
        renderButtons(radarFooter);

        const landingPageElements =
          document.querySelectorAll("main .home-page");
        landingPageElements.forEach((elem) => {
          elem.style.display = "none";
        });
      } else {
        plotRadarHeader();
        plotRadarFooter();
        if (alternatives.length) {
          plotAlternativeRadars(alternatives, currentSheet);
        }
        plotQuadrantButtons(quadrants);
      }

      svg = radarElement.append("svg").call(tip);

      if (featureToggles.UIRefresh2022) {
        const legendHeight = 40;
        radarElement.style("height", size + legendHeight + "px");
        svg.attr("id", "radar-plot").attr("width", size).attr("height", size);
      } else {
        radarElement.style("height", size + 14 + "px");
        svg
          .attr("id", "radar-plot")
          .attr("width", size)
          .attr("height", size + 14);
      }

      _.each(quadrants, function (quadrant) {
        let quadrantGroup;
        if (featureToggles.UIRefresh2022) {
          quadrantGroup = renderRadarQuadrants(
            size,
            svg,
            quadrant,
            rings,
            ringCalculator,
            tip
          );
          plotLines(quadrantGroup, quadrant);
          const ringTextGroup = quadrantGroup.append("g");
          plotRingNames(ringTextGroup, rings, quadrant);
          plotRadarBlips(quadrantGroup, rings, quadrant, tip);
          renderMobileView(quadrant);
          addQuadrantNameInPdfView(quadrant.order, quadrant.quadrant.name());
        } else {
          quadrantGroup = plotQuadrant(rings, quadrant);
          plotLines(quadrantGroup, quadrant);
          plotTexts(quadrantGroup, rings, quadrant);
          plotBlips(quadrantGroup, rings, quadrant);
        }
      });

      if (featureToggles.UIRefresh2022) {
        renderRadarLegends(radarElement);
        hideTooltipOnScroll(tip);
        addRadarLinkInPdfView();
      }
    };

    return self;
  };

  /**
   * UTIL: js/util/factory.js
   */
  const plotRadar = function (blips, currentRadarName) {
    document.title = "plotRaderTitle";

    var rings = _.map(_.uniqBy(blips, "ring"), "ring");
    var ringMap = {};

    _.each(rings, function (ringName, i) {
      ringMap[ringName] = new Ring(ringName, i);
    });

    var quadrants = {};
    _.each(blips, function (blip) {
      if (!quadrants[blip.quadrant]) {
        quadrants[blip.quadrant] = new Quadrant(
          blip.quadrant[0].toUpperCase() + blip.quadrant.slice(1)
        );
      }
      quadrants[blip.quadrant].add(
        new Blip(
          blip.name,
          ringMap[blip.ring],
          blip.isNew.toLowerCase() === "true",
          blip.topic,
          blip.description
        )
      );
    });

    var radar = new Radar();
    _.each(quadrants, function (quadrant) {
      radar.addQuadrant(quadrant);
    });

    if (currentRadarName !== undefined || true) {
      radar.setCurrentSheet(currentRadarName);
    }

    const size = featureToggles.UIRefresh2022
      ? getGraphSize()
      : window.innerHeight - 133 < 620
      ? 620
      : window.innerHeight - 133;
    new GraphingRadar(size, radar).init().plot();
  };

  function validateInputQuadrantOrRingName(
    allQuadrantsOrRings,
    quadrantOrRing
  ) {
    const quadrantOrRingNames = Object.keys(allQuadrantsOrRings);
    const regexToFixLanguagesAndFrameworks = /(-|\s+)(and)(-|\s+)|\s*(&)\s*/g;
    const formattedInputQuadrant = quadrantOrRing
      .toLowerCase()
      .replace(regexToFixLanguagesAndFrameworks, " & ");
    return quadrantOrRingNames.find(
      (quadrantOrRing) =>
        quadrantOrRing.toLowerCase() === formattedInputQuadrant
    );
  }

  const plotRadarGraph = function (blips, currentRadarName) {
    document.title = "Novo Nordisk Tech Radar";

    const ringMap = graphConfig.rings.reduce((allRings, ring, index) => {
      allRings[ring] = new Ring(ring, index);
      return allRings;
    }, {});

    const quadrants = graphConfig.quadrants.reduce((allQuadrants, quadrant) => {
      allQuadrants[quadrant] = new Quadrant(quadrant);
      return allQuadrants;
    }, {});

    blips.forEach((blip) => {
      const currentQuadrant = validateInputQuadrantOrRingName(
        quadrants,
        blip.quadrant
      );
      const ring = validateInputQuadrantOrRingName(ringMap, blip.ring);
      if (currentQuadrant && ring) {
        const blipObj = new Blip(
          blip.name,
          ringMap[ring],
          blip.isNew.toLowerCase() === "true",
          blip.topic,
          blip.description
        );
        quadrants[currentQuadrant].add(blipObj);
      }
    });

    const radar = new Radar();
    radar.addRings(Object.values(ringMap));

    _.each(quadrants, function (quadrant) {
      radar.addQuadrant(quadrant);
    });

    radar.setCurrentSheet(currentRadarName);

    const graphSize =
      window.innerHeight - 133 < 620 ? 620 : window.innerHeight - 133;
    const size = featureToggles.UIRefresh2022 ? getGraphSize() : graphSize;
    new GraphingRadar(size, radar).init().plot();
  };

  const CSVDocument = function (csvData) {
    var self = {};

    self.build = function () {
      csvfile = d3.csvParse(csvData);
      createBlips(csvfile);
    };

    var createBlips = function (data) {
      delete data.columns;
      var blips = _.map(data);
      featureToggles.UIRefresh2022
        ? plotRadarGraph(blips, "CSV File", [])
        : plotRadar(blips, "CSV File", []);
    };

    return self;
  };

  const Factory = function (sheetData) {
    var sheet;
    sheet = CSVDocument(sheetData);
    return sheet;
  };

  return Factory;
});
