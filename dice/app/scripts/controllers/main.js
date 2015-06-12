'use strict';

/**
 * @ngdoc function
 * @name diceApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the diceApp
 */
angular.module('diceApp')
  .controller('MainCtrl', function ($http) {
    this.rollUrl = '/dice/roll';
    this.dice = {
        'd4': 0,
        'd6': 0,
        'd8': 0,
        'd10': 0,
        'd12': 0,
        'd20': 0
    };
    
    this.roll = {
        'd4': 0,
        'd6': 0,
        'd8': 0,
        'd10': 0,
        'd12': 0,
        'd20': 0
    };
    
    this.submit = function(dice) {
        var roll = $http.post(this.rollUrl, dice)
            .success(function(resp) {
                console.log("resp.data:");
                console.log(resp.data);
                roll = resp.data;
                return roll
        });
        this.roll = roll;
        console.log("this.roll:");
        console.log(this.roll);
    };
  });
