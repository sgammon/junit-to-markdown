# junit-to-markdown

## Usage

```
$ make dockerbuild
$ cd *application*
$ docker run -v target/test-reports:/test-reports junit-to-markdown /test-reports
# Test results
### com.pagero.services.example.handlers.HelloRequestHandlerTest
|Success|Test|
|-------|----|
|:white_check_mark:|getResponse should say hello to someone when asked|
|:x:|getResponse should be extra polite when speaking to a queen|
```

# -->

# Test results
### com.pagero.services.example.handlers.HelloRequestHandlerTest
|Success|Test|
|-------|----|
|:white_check_mark:|getResponse should say hello to someone when asked|
|:x:|getResponse should be extra polite when speaking to a queen|
