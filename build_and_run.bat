@echo off

call wails3 generate bindings -d frontend

cd "frontend"

call npm install

call npm run build

if %errorlevel% neq 0 (
    echo Error: npm run build failed.
    exit /b %errorlevel%
)

cd ..

call go mod tidy
call go run .
