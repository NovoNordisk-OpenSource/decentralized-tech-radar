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
    
  }
  return config
});
