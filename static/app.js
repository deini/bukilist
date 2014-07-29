angular.module('app', [])

    .controller('MainCtrl', function($scope, $http) {
        $http.get('/api/bukiwishes')
            .success(function(data) {
                console.log(data);
                $scope.bukiwishes = data;
            });
    });