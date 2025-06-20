NAME = goNest
APP_NAME = gonest

MAIN_FILE = cmd/gonest/main.go
BUILD_DIR = build

build:
	@echo "[${NAME}] >>> Building..."
	go build -o ${BUILD_DIR}/${APP_NAME} ${MAIN_FILE}

run:
	@echo "[${NAME}] >>> Running..."
	./${BUILD_DIR}/${APP_NAME}

clean:
	@echo "[${NAME}] >>> Cleaning..."
	rm -rf ${BUILD_DIR}/

copy:
	cp -r ${BUILD_DIR}/${APP_NAME} ./

all: clean build copy run