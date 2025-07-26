$remoteFile = "https://github.com/raysan5/raylib/releases/download/5.5/raylib-5.5_win64_msvc16.zip"
$ignoreDir = "./ignore"

$fontRemote = "https://github.com/googlefonts/roboto-2/releases/download/v2.138/roboto-unhinted.zip"
$fontArchiveLocal = "$ignoreDir/roboto-unhinted.zip"

# Setup Ignore
if (!(Test-Path -Path $ignoreDir)) {
  New-Item -ItemType Directory -Path $ignoreDir
}
$localFile = "$ignoreDir/raylib.zip"
$localDir = "$ignoreDir/raylib-extracted"

# Fetch and copy DLL
Invoke-WebRequest $remoteFile -OutFile $localFile

Expand-Archive -Path $localFile -DestinationPath $localDir

$dll = "$localDir\raylib-5.5_win64_msvc16\lib\raylib.dll"

Copy-Item -Path $dll -Destination ".\"

# Fetch Font
Invoke-WebRequest $fontRemote -OutFile $fontArchiveLocal

Expand-Archive -Path $fontArchiveLocal -DestinationPath "$ignoreDir/roboto"
