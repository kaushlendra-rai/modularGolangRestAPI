This project exposes REST API for a imaginary Employee entity and provides basic CRUD opertaion around it.

I have used:
1) github.com/gorilla/mux : For REST API handling and routing
2) github.com/jinzhu/gorm : For ORM functionalities
3) Connects to Postgres DB: At this point, I do not have automatic DB  creation flow. One needs to manually create 'sinkar' named DB in the target Postgres instance.
4) github.com/spf13/viper : To read configuration from YAML file
5) Modules for Dependency management and keeping project out of GO_PATH location for better management.
6) main_test.go exists for basic test without actually deploying the service. The test however runs against actual Postgres instance instead of mocking postgres.
7) Code coverage aspect is also shared in the main_test.go file.
8) Postman collection 'GoLangSampleWebApp.postman_collection.json' is also shared at top level of project fo reasy testing of the code.


Ensure that GOPATH & GOROOT are available as environment variable. Optionally, you can also add ${GOPATH}/bin to PATH variable so that you could execute
the installed (go install) executables of go application.