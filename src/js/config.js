define(function() {
  const config = () => {
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
    
    //return process.env.ENVIRONMENT ? env[process.env.ENVIRONMENT] : env
    
  }

  // require-config.js
// require.config({
//   baseUrl: 'path/to/your/js/folder', // This should be the base path for your JavaScript files.
//   paths: {
//     'underscore': 'path/to/underscore', // Update these paths to point to where you have these libraries.
//     'd3': 'path/to/d3'
//   },
//   shim: {
//     // If these libraries do not support AMD, you might need to use shim configuration.
//   }
// });

// // Load the main app module to start the app
// requirejs(['path/to/your/main']);

  return config
});
