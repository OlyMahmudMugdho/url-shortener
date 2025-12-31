.PHONY: all build-frontend build-backend run clean

APP_NAME=url-shortener
FRONTEND_DIR=frontend/url-shortener-frontend
BACKEND_DIR=backend

all: clean build-frontend build-backend

build-frontend:
	@echo "Building Frontend..."
	cd $(FRONTEND_DIR) && npm install && npm run build -- --configuration production
	@echo "Copying frontend build to backend..."
	mkdir -p $(BACKEND_DIR)/dist
	cp -r $(FRONTEND_DIR)/dist/* $(BACKEND_DIR)/dist/

build-backend:
	@echo "Building Backend..."
	cd $(BACKEND_DIR) && go build -o $(APP_NAME) main.go

run: all
	@echo "Running Application..."
	cd $(BACKEND_DIR) && ./$(APP_NAME)

clean:
	@echo "Cleaning..."
	rm -rf $(BACKEND_DIR)/dist
	rm -f $(BACKEND_DIR)/$(APP_NAME)
