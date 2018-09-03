@set GOOS=linux
@rm -rf ./dist
@mkdir dist
@cp -r -f ./static ./dist/
@cp -r -f ./config ./dist/
@go build -o ./dist/{{name}} ./src
@set GOOS=
@rm ./tmp/*.gz
@rm ./tmp/*.tar
@tar -cvf ./tmp/{{name}}.tar -C ./dist .
@gzip -v ./tmp/{{name}}.tar