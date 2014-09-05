EXECUTABLE=karta-gui
APP_BUNDLE=Karta.app

all: build quit open

build:
	go build -o $(APP_BUNDLE)/Contents/MacOS/${EXECUTABLE}

quit:
	osascript -e 'quit app "$(APP_BUNDLE)"'

open:
	open -a $(APP_BUNDLE)

clean:
	rm -f $(APP_BUNDLE)/Contents/MacOS/$(EXECUTABLE)
