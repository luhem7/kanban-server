var myApp = angular.module('BoardApp', ['ngAnimate', 'ngResource']);

myApp.factory('taskColumnsRes', function($resource){
    return $resource('./data/columns', {});
});

myApp.factory('tasksRes', function($resource){
    return $resource('./data/tasks', {});
});

myApp.controller('BoardCtrl', ['$scope', 'taskColumnsRes', 'tasksRes', function($scope, taskColumnsRes, tasksRes){
    //Initializing data
    $scope.tasks = {};
    $scope.taskColumns = {};
    taskColumnsRes.get(function(data){
        $scope.taskColumns = data;
    });

    tasksRes.get(function(data){
        $scope.tasks = data;
    });

    $scope.getNumTasksByColumn = function(column){
        if (angular.isUndefined($scope.tasks[column.columnName])) {
            return 0;
        }
        return $scope.tasks[column.columnName].length;
    }

    $scope.getTasksByColumn = function(column){
        return $scope.tasks[column.columnName];
    };
}]);
