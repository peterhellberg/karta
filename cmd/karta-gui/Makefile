EXECUTABLE=karta-gui
APP_BUNDLE=Karta.app

all: build open

build:
	go build -o $(APP_BUNDLE)/Contents/MacOS/${EXECUTABLE}

open:
	open -a Karta.app

clean:
	rm -f $(APP_BUNDLE)/Contents/MacOS/$(EXECUTABLE)
