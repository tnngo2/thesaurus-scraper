#A therausus scraper for vocabulary studying purpose
This micro service will generate file that is used as imported content to Memrise.

Written in Go lang, based on Gin web framework.

##Ready for heroku deployment

```
$ heroku login
$ git init
$ git add -A .
$ git commit -m code
$ echo 'web: demoapp' > Procfile
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
$ git push heroku master
$ heroku open
$ heroku ps
$ heroku logs --tail
```
