# Remember ./ in Linux is .\ in Windows.
#
# Development:
# ./start.ps1 "./pfs.toml" -D
#
# Production:
# ./start.ps1 "./pfs.toml"

Param(
  [Parameter(Mandatory = $true, Position = 0)]
  [ValidateNotNullOrEmpty()]
  [string] $PathToConfig,
  
  [Parameter(Mandatory = $false, Position = 1)]
  [switch] $Dev
)

If (-not (Test-Path -Path $PathToConfig -PathType Leaf)) {
  Throw "$PathToConfig is not a valid path"
}

Copy-Item $PathToConfig backend
Copy-Item $PathToConfig nginx/conf
Copy-Item $PathToConfig frontend
If ($Dev) {
  Write-Host "Using development configuration"
  docker compose -f dev-docker-compose.yml up
} Else {
  Write-Host "Using production configuration"
  docker compose up
}
