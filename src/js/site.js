// require(['./common.js'])

// const Factory = require(['./util/factory.js'])

//Factory('name,ring,quadrant,isNew,moved,description\nRJDcGqUtECqUYD,Hold,data management,true,0,uOceYNVwMWrBLNFVjExmhkSxSeodnnggLHyGSqZpX').build()

require(['./js/common.js','./js/util/factory.js'], 
  function(common, Factory) {
  
  // Use the imported modules
  const data = `name,ring,quadrant,isNew,moved,description
                 RJDcGqUtECqUYD,Hold,data management,true,0,uOceYNVwMWrBLNFVjExmhkSxSeodnnggLHyGSqZpX`;

  // Call the factory function from 'util/factory.js'
  Factory(data).build();
});