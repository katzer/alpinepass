cd src
call "go" test ./...
call "go" build -o alpinepass.exe
move alpinepass.exe ../alpinepass.exe
cd ..
