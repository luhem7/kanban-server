var testColumns =
{
    "columns": [
        {
            "columnName" : "TO DO",
            "numTasks" : 0,
            "isVisible" : true
        },
        {
            "columnName" : "In Progress",
            "numTasks" : 0,
            "isVisible" : true
        },
        {
            "columnName" : "Waiting",
            "numTasks" : 0,
            "isVisible" : true
        },
        {
            "columnName" : "Done",
            "numTasks" : 0,
            "isVisible" : true
        }
    ]
};

var testTasks = {
    "TO DO": [
        {
            "task-title" : "Do the dishes",
            "task-body" : ""
        },
        {
            "task-title" : "Win the lottery!",
            "task-body" : "Need the money man"
        },
        {
            "task-title" : "Book the ticket",
            "task-body" : "To travel the world!"
        },
        {
            "task-title" : "Book the movie ticket",
            "task-body" : "TO watch the movie!"
        },
        {
            "task-title" : "Find a better way to animate column show/hide behavior",
            "task-body" : "I dont like the current fade in and fade out method. I would rather have a cool rollup drawer type effect."
        },
        {
            "task-title" : "Take car for oil change",
            "task-body" : "30 percent oil change life left"
        }
    ],
    "In Progress": [
        {
            "task-title" : "Buy cat food",
            "task-body" : "He feels the hunger. Quis auctor feugiat egestas ut rutrum ante turpis etiam auctor ante tellus, taciti tellus suspendisse leo fermentum proin consectetur dapibus vivamus tincidunt velit, facilisis aenean consectetur mi nisl luctus erat feugiat odio quis"
        }
    ],
    "Waiting": [
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

var myApp = angular.module('BoardApp', ['ngAnimate']);

myApp.controller('BoardCtrl', ['$scope', function($scope){
    //Initializing
    $scope.taskColumns = testColumns;
    $scope.testTasks = testTasks;

    $scope.getNumTasksByColumn = function(column){
        if (angular.isUndefined(testTasks[column.columnName])) {
            return 0;
        }
        return testTasks[column.columnName].length;
    }

    $scope.getTasksByColumn = function(column){
        return testTasks[column.columnName];
    };
}]);
