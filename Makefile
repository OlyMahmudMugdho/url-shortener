.PHONY: all build-frontend build-backend run clean test test-backend test-frontend

APP_NAME=url-shortener
FRONTEND_DIR=frontend/url-shortener-frontend
BACKEND_DIR=backend


DOCKER_IMAGE=olymahmudmugdho/url-shortener

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

test-backend:
	@echo "Testing Backend..."
	cd $(BACKEND_DIR) && go test ./...

test-frontend:
	@echo "Testing Frontend..."
	cd $(FRONTEND_DIR) && npm install && npm test -- --watch=false --browsers=ChromeHeadless

test: test-backend test-frontend

run: all
	@echo "Running Application..."
	cd $(BACKEND_DIR) && ./$(APP_NAME)

docker-build: all
	@echo "Building Docker Image..."
	cd $(BACKEND_DIR) && docker build -t $(DOCKER_IMAGE) .

docker-push: docker-build
	@echo "Pushing Docker Image..."
	docker push $(DOCKER_IMAGE)

clean:
	@echo "Cleaning..."
	rm -rf $(BACKEND_DIR)/dist
	rm -f $(BACKEND_DIR)/$(APP_NAME)
