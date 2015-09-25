Param(
    [Parameter(Mandatory=$False)][boolean]$onlyCopyRes=$False
)

#Script variables
$project_name = "github.com\luhem7\kanban-board"
$src_dir = $Env:GOPATH + "\src\" + $project_name
$bin_dir = $Env:GOPATH + "\bin\" + $project_name
$res_path = "\res"

Write-Host "Building "$project_name" project"
Write-Host "Copying Res folder to bin directory"
Copy-Item ($src_dir+$res_path) $bin_dir -recurse -force

If(-Not($onlyCopyRes)) {
    Write-Host "Building the go executable"
    go build $project_name
}

Write-Host "Done"
