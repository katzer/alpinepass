RED=`tput setaf 1`
RESET=`tput sgr0`

echo "${RED}Running unit tests.${RESET}"
cd src
go test ./...
echo "${RED}Building the application.${RESET}"
go build -o alpinepass
echo "${RED}Copying the application.${RESET}"
mv alpinepass ../alpinepass
