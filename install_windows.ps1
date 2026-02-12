$installDir = "$env:LOCALAPPDATA\Programs\reporeport"
New-Item -ItemType Directory -Force -Path $installDir | Out-Null
Invoke-WebRequest -Uri "https://github.com/mathealgou/reporeport/releases/latest/download/reporeport.exe" -OutFile "$installDir\reporeport.exe"
[Environment]::SetEnvironmentVariable("Path", "$env:Path;$installDir", "User")
Write-Host "Installed to $installDir. Restart your terminal to use 'reporeport'."