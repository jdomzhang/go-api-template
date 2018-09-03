@set GOOS=linux
@rm -rf ./dist
@mkdir dist
@cp -r -f ./static ./dist/
@cp -r -f ./config ./dist/
@go build -o ./dist/{{name}} ./src
@set GOOS=
@rm ./dist/*.gz
@rm ./dist/*.tar
@tar -cvf ./dist/{{name}}.tar -C ./dist .
@gzip -v ./dist/{{name}}.tar