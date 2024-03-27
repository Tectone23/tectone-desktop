PACKAGE="tectone/tectone-desktop"
COMMIT_HASH="%(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

ifndef VERSION
override VERSION="$(git describe --tags --always --abrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
endif

# FOR LOCAL DEVELOPMENT: If the URLs are not set in the environment the default algorand urls are used.

ifndef SDK_URL
override SDK_URL= "https://github.com/algorand/go-algorand-sdk.git"
endif

ifndef SANDBOX_URL 
override SANDBOX_URL= "https://github.com/algorand/sandbox.git"
endif


LDFLAGS=(
  "-X '${PACKAGE}/main.Version=${VERSION}'"
  "-X '${PACKAGE}/main.CommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}/main.BuildTime=${BUILD_TIMESTAMP}'"
  "-X '${PACKAGE}/main.SDKGitURL=${SDK_URL}'"
  "-X '${PACKAGE}/main.SandboxGitURL=${SANDBOX_URL}'"

)

# create-dmg specific values
APP_NAME="Tectone Desktop"
DMG_FILE_NAME="${APP_NAME}-Installer.dmg"
VOLUME_NAME="${APP_NAME} Installer"
SOURCE_FOLDER_PATH="build/bin/"
BACKGROUND_IMAGE="build/installer_background.png"

all:
ifeq ($(OS),$(filter $(OS), Windows_NT, Darwin))
	all: build-macos build-windows
else
	all: build-macos, build-windows, build-linux
endif

build-macos:
	wails build  --platform darwin/amd64 --ldflags=${LDFLAGS[*]}

	# create dmg 
	$CREATE_DMG \
  --volname "${VOLUME_NAME}" \
  --background "${BACKGROUND_IMAGE}" \
  --window-pos 200 120 \
  --window-size 800 400 \
  --icon-size 100 \
  --icon "${APP_NAME}.app" 200 190 \
  --hide-extension "${APP_NAME}.app" \
  --app-drop-link 600 185 \
  "${DMG_FILE_NAME}" \
  "${SOURCE_FOLDER_PATH}"

build-windows:
	wails build --nsis  --platform windows/amd64 --ldflags=${LDFLAGS[*]}

build-linux:
	wails build  --platform linux/amd64 --ldflags=${LDFLAGS[*]}
