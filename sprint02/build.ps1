Write-Host "Building for all platforms..."

New-Item -ItemType Directory -Force -Path builds | Out-Null

$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o builds/cbp-gen-darwin-arm64 .
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o builds/cbp-gen-darwin-amd64 .
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o builds/cbp-gen-linux-amd64 .
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o builds/cbp-gen-windows-amd64.exe .

Write-Host "Done! Binaries are in /builds"