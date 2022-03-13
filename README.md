##Example:

run server command:
    `go run main.go`

send request like this:
    `curl --location --request POST 'http://localhost:3000/test/results?baseUrl={YOUR_ALLURE_RESULT_HOSTING_URL}' \
    --form "file=@\"{YOUR_PATH_OF_ALLURE_REPORT_FOLDER}/suites.json\""`

suites.json located inside allure-report folder
