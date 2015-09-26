Param(
    [Parameter(Mandatory=$False)][boolean]$skipCopyRes=$False,
    [Parameter(Mandatory=$False)][boolean]$skipCompileGo=$False,
    [Parameter(Mandatory=$False)][boolean]$skipForceClean=$False
)

#Script variables
$project_name = "github.com\luhem7\kanban-board"
$src_dir = $Env:GOPATH + "\src\" + $project_name
$bin_dir = $Env:GOPATH + "\bin\" + $project_name
$res_path = "\res"

If($skipCompileGo) {
    Write-Host "Skipping building the go executable"
} Else {
    Write-Host "Building the go executable"
    Set-Location $bin_dir
    go build $project_name
}

If($skipForceClean){
    Write-Host "Skipping force clean of res directory"
} Else {
    Remove-Item ($bin_dir+$res_path) -recurse -force
}

If($skipCopyRes){
    Write-Host "Skipping copying Res folder to bin directory"
} Else {
    Write-Host "Building "$project_name" project"
    Write-Host "Copying Res folder to bin directory"
    Copy-Item ($src_dir+$res_path) $bin_dir -recurse -force
}

Write-Host "Done"
