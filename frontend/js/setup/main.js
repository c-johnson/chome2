var C = C || {};

require.config({
  shim: {
    'foundation': {
      deps: ['jquery', 'modernizr'],
      exports: 'Foundation'
    },
    'fo-topbar': {
      deps: ['foundation'],
      exports: 'FO-Topbar'
    },
    'angular': {
      exports: 'angular'
    },
    'ui-router': {
      deps: ['angular']
    }
  },
  baseUrl: "public/js",
  paths: {
    "angular": "../bowerc/angular/angular",
    "ui-router": "../bowerc/angular-ui-router/release/angular-ui-router",
    "domReady": "../bowerc/requirejs-domready/domReady",
    "jquery": "../bowerc/jquery/jquery",
    "modernizr": "../bowerc/modernizr/modernizr",
    "foundation": "../bowerc/foundation/js/foundation",
    "fo-topbar": "../bowerc/foundation/js/foundation/foundation.topbar",
    "underscore": "../bowerc/underscore/underscore"
  },
  deps: ['setup/bootstrap'],
  waitSeconds: 15
});