document.addEventListener("touchstart", function() {},false); //Let mobile Safari use :active css pseudo class

var myApp = angular.module('BoardApp', ['ngAnimate', 'ngResource']);

myApp.factory('taskColumnsRes', function($resource){
    return $resource('./data/columns', {});
});

myApp.factory('tasksRes', function($resource){
    return $resource('./data/tasks', {});
});

myApp.controller('BoardCtrl', ['$scope', 'taskColumnsRes', 'tasksRes', function($scope, taskColumnsRes, tasksRes){
    //Initializing data definitions
    $scope.tasks = {}; //The tasks on the board
    $scope.taskColumns = {}; //The task columns on the board
    $scope.isEditingColumns = false; //User state where board is being edited.

    //Fetching initial data values
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
