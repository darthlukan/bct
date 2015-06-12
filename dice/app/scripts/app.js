'use strict';

/**
 * @ngdoc overview
 * @name diceApp
 * @description
 * # diceApp
 *
 * Main module of the application.
 */
angular
  .module('diceApp', [
    'ngMessages',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/dice', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .otherwise({
        redirectTo: '/dice'
      });
  });
