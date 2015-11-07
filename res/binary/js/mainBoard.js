var testColumns =
{
    "columns": [
        {
            "columnName" : "TO DO",
            "numTasks" : 0
        },
        {
            "columnName" : "In Progress",
            "numTasks" : 0
        },
        {
            "columnName" : "Waiting",
            "numTasks" : 0
        },
        {
            "columnName" : "Done",
            "numTasks" : 0
        }
    ]
};

var testTasks = {
    "TO DO-Tasks": [
        {
            "task-title" : "Do the dishes",
            "task-body" : ""
        },
        {
            "task-title" : "Take car for oil change",
            "task-body" : "30 percent oil change life left"
        }
    ],
    "In Progress-Tasks": [
        {
            "task-title" : "Buy cat food",
            "task-body" : "He feels the hunger. Quis auctor feugiat egestas ut rutrum ante turpis etiam auctor ante tellus, taciti tellus suspendisse leo fermentum proin consectetur dapibus vivamus tincidunt velit, facilisis aenean consectetur mi nisl luctus erat feugiat odio quis"
        }
    ],
    "Waiting-Tasks": [
        {
            "task-title" : "Submit form",
            "task-body" : "Submit form for tortillas"
        },
        {
            "task-title" : "Order case",
            "task-body" : "Order case from france"
        }
    ]
};

var myApp = angular.module('BoardApp', []);

myApp.controller('BoardCtrl', ['$scope', function($scope){
    //Initializing
    $scope.taskColumns = testColumns;
    $scope.testTasks = testTasks;

    $scope.getNumTasksByColumnName = function(columnName){
        if (angular.isUndefined(testTasks[columnName+"-Tasks"])) {
            return 0;
        }
            return testTasks[columnName+"-Tasks"].length;
    }

    $scope.getTasksByColumnName = function(columnName){
        return testTasks[columnName+"-Tasks"];
    };
}]);
