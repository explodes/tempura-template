PROGRAM_NAME=template

DESKTOP_MAIN_GOPKG=github.com/explodes/tempura-template/go/cmd/app
DESKTOP_MAIN=./go/cmd/app/main.go

MOBILE_MAIN_GOPKG=github.com/explodes/tempura-template/go/cmd/mobile

ANDROID_PKG=com.example.android.game
ANDROID_LIB=./android/gamelib/gamelib.aar
ANDROID_LIB_SRC=./android/gamelib/gamelib-sources.jar
ANDROID_ACTIVITY=com.example.android.game/com.example.android.game.MainActivity

GENERATED_SETTINGS=./go/core/gen.go
BUILD_OUTPUT_DIR=./build

## +-+-+ STATIC SETTINGS ##

RESOURCE_DIR=./resources
BINDATA_FILE=./go/res/bindata.go
BINDATA_FILE_RELATIVE_TO_RESOURCES=../go/res/bindata.go

## +-+-+ STANDARD +-+-+ ##

clean-res:
	rm ./go/game/res/bindata.go || true

clean-build:
	find . -type d -name build | xargs -n1 rm -rf
	rm "$(GENERATED_SETTINGS).bak" || true
	rm "$(ANDROID_LIB)" || true
	rm "$(ANDROID_LIB_SRC)" || true

clean: enable-debug clean-res clean-build

res: clean-res
	mkdir -p "$(RESOURCE_DIR)" || true
	mkdir -p "$(RESOURCE_DIR)/images" || true
	mkdir -p "$(RESOURCE_DIR)/fonts" || true
	mkdir -p "$(RESOURCE_DIR)/sound" || true
	mkdir -p "$(RESOURCE_DIR)/music" || true
	cd "$(RESOURCE_DIR)"; "$(GOPATH)/bin/go-bindata" -nocompress -o "$(BINDATA_FILE_RELATIVE_TO_RESOURCES)" -pkg res ./...

enable-debug:
	sed -i.bak 's/Debug = false/Debug = true/g' "$(GENERATED_SETTINGS)"

disable-debug:
	sed -i.bak 's/Debug = true/Debug = false/g' "$(GENERATED_SETTINGS)"

create-generated-files: res

## +-+-+ GO DEPS +-+-+ ##

go-deps:
	go get -u -v github.com/jteeuwen/go-bindata/...
	go get -u -v golang.org/x/mobile/cmd/gomobile
	go get -u -v golang.org/x/mobile/cmd/gobind
	go get -u github.com/gopherjs/gopherjs

go-prep:
	gomobile init


## +-+-+ ANDROID +-+-+ ##

android-lib: create-generated-files
	CGO_ENABLED=1 gomobile bind -target android -javapkg "$(ANDROID_PKG)" -o "$(ANDROID_LIB)" "$(MOBILE_MAIN_GOPKG)"

# debug: build & run
android: enable-debug android-lib
	cd ./android; ./gradlew ':app:installDebug'
	mkdir -p "$(BUILD_OUTPUT_DIR)" || true
	cp ./android/app/build/outputs/apk/app-debug.apk "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-android-debug.apk"
	adb shell am start -n "$(ANDROID_ACTIVITY)"

# debug: build
android-debug: enable-debug android-lib
	cd ./android; ./gradlew ':app:assembleDebug'
	mkdir -p "$(BUILD_OUTPUT_DIR)" || true
	cp ./android/app/build/outputs/apk/app-debug.apk "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-android-debug-unsigned.apk"

# release: build
android-release: disable-debug android-lib
	cd ./android; ./gradlew ':app:assembleRelease'
	mkdir -p "$(BUILD_OUTPUT_DIR)" || true
	cp ./android/app/build/outputs/apk/app-release-unsigned.apk "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-android-release-unsigned.apk"

## +-+-+ WEB +-+-+ ##

# debug: build & run
web: enable-debug create-generated-files
	google-chrome http://localhost:8080
	CGO_ENABLED=1 gopherjs serve -m "$(DESKTOP_MAIN_GOPKG)"

# debug: build
web-debug: enable-debug create-generated-files
	CGO_ENABLED=1 gopherjs build -m -q -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-web-debug.js" "$(DESKTOP_MAIN_GOPKG)"

# release: build
web-release: disable-debug create-generated-files
	CGO_ENABLED=1 gopherjs build -m -q -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-web-release.js" "$(DESKTOP_MAIN_GOPKG)"

## +-+-+ LINUX +-+-+ ##

# pre-reqs:
# ubunutu: sudo apt install libglu1-mesa-dev libgles2-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev
# fedora: sudo dnf install mesa-libGLU-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel alsa-lib-devel
# solus: sudo eopkg install libglu-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel

# debug: build & run
linux: enable-debug create-generated-files
	GOOS=linux GOARCH=amd64 go run "$(DESKTOP_MAIN)"

# debug: build
linux-debug: enable-debug create-generated-files
	GOOS=linux GOARCH=amd64 go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-linux-amd64-debug" "$(DESKTOP_MAIN_GOPKG)"

# release: build
linux-release: disable-debug create-generated-files
	GOOS=linux GOARCH=amd64 go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-linux-amd64-release" "$(DESKTOP_MAIN_GOPKG)"

## +-+-+ WINDOWS +-+-+ ##

# additional pre-reqs:
# ubunutu: sudo apt install gcc-multilib gcc-mingw-w64

# debug: build & run
windows: enable-debug create-generated-files
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go run "$(DESKTOP_MAIN)"

# debug: build
windows-debug: enable-debug create-generated-files
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-windows-i386-debug.exe" "$(DESKTOP_MAIN_GOPKG)"

# release: build
windows-release: disable-debug create-generated-files
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-windows-i386-release.exe" "$(DESKTOP_MAIN_GOPKG)"


## +-+-+ DARWIN +-+-+ ##

# additional pre-reqs:
# ubunutu: sudo apt install gobjc++ libopenal1 libopenal-dev libgnustep-gui-dev

# debug: build & run
darwin: enable-debug create-generated-files
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go run "$(DESKTOP_MAIN)"

# debug: build
darwin-debug: enable-debug create-generated-files
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-darwin-amd64-debug $(DESKTOP_MAIN_GOPKG)"

# release: build
darwin-release: disable-debug create-generated-files
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o "$(BUILD_OUTPUT_DIR)/$(PROGRAM_NAME)-amd64-release $(DESKTOP_MAIN_GOPKG)"


## +-+-+ ALL +-+-+ ##

releases: android-release web-release linux-release windows-release darwin-release
debugs: android-debug web-debug linux-debug windows-debug darwin-debug
all: releases debugs