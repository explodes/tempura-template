

## +-+-+ STANDARD +-+-+ ##

clean-res:
	rm ./go/game/res/bindata.go || true

clean-build:
	find . -type d -name build | xargs -n1 rm -rf
	rm ./go/game/config.go.bak || true
	rm ./android/gamelib/gamelib.aar || true
	rm ./android/gamelib/gamelib-sources.jar || true

clean: enable-debug clean-res clean-build

res: clean-res
	cd ./go/internal/resources; $(GOPATH)/bin/go-bindata -nocompress -o ../res/bindata.go -pkg res -ignore '\.*' ./...

enable-debug:
	sed -i.bak 's/debug = false/debug = true/g' ./go/internal/game/config.go

disable-debug:
	sed -i.bak 's/debug = true/debug = false/g' ./go/internal/game/config.go

## +-+-+ ANDROID +-+-+ ##

android-lib: res
	gomobile bind -target android -javapkg com.example.android.game -o ./android/gamelib/gamelib.aar github.com/explodes/tempura-template/go/cmd/mobile

# debug: build & run
android: enable-debug android-lib
	cd ./android; ./gradlew ':app:installDebug'
	mkdir -p ./build || true
	cp ./android/app/build/outputs/apk/app-debug.apk ./build/game-debug.apk
	adb shell am start -n com.example.android.game/com.example.android.game.MainActivity

# release: build
android-release: disable-debug android-lib
	cd ./android; ./gradlew ':app:assembleRelease'
	mkdir -p ./build || true
	cp ./android/app/build/outputs/apk/app-release-unsigned.apk ./build/game-release-unsigned.apk

## +-+-+ WEB +-+-+ ##

# debug: build & run
web: enable-debug res
	google-chrome http://localhost:8080
	gopherjs serve -m github.com/explodes/tempura-template/go/cmd/app

# release: build
web-release: disable-debug res
	gopherjs build -m -q -o ./build/game.js github.com/explodes/tempura-template/go/cmd/app

## +-+-+ LINUX +-+-+ ##

# pre-reqs:
# ubunutu: sudo apt install libglu1-mesa-dev libgles2-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev
# fedora: sudo dnf install mesa-libGLU-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel alsa-lib-devel
# solus: sudo eopkg install libglu-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel

# debug: build & run
linux: enable-debug res
	GOOS=linux GOARCH=amd64 go run ./go/cmd/app/main.go

# release: build
linux-release: disable-debug res
	GOOS=linux GOARCH=amd64 go build -o ./build/game github.com/explodes/tempura-template/go/cmd/app


## +-+-+ ALL +-+-+ ##

releases: android-release web-release linux-release