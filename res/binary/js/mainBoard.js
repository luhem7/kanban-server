var testColumns =
{
    "Columns": [
        "TO DO",
        "In Progress",
        "Waiting",
        "Done"
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
    $scope.taskColumns = testColumns;
    $scope.testTasks = testTasks;

    $scope.getTasksByColumnName = function(columnName){
        return testTasks[columnName+"-Tasks"];
    };
}]);
