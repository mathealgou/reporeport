$ErrorActionPreference = "Stop"

$installDir = "$env:LOCALAPPDATA\Programs\reporeport"
$exePath = Join-Path $installDir "reporeport.exe"
$downloadUri = "https://github.com/mathealgou/reporeport/releases/latest/download/reporeport.exe"

New-Item -ItemType Directory -Force -Path $installDir | Out-Null
Invoke-WebRequest -Uri $downloadUri -OutFile $exePath

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $userPath) { $userPath = "" }

if ($userPath -notmatch [Regex]::Escape($installDir)) {
    $newPath = if ($userPath) { "$userPath;$installDir" } else { $installDir }
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
}

Write-Host "Installed to $installDir. Restart your terminal to use 'reporeport'."
