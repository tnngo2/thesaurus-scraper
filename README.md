# A therausus scraper
I used this project to scrap Cambridge therausus content and copy those scrapped content to Memrise for reviewing purpose.

The project is configured as a micro service which is ready for Heroku deployment.

Written in Go lang, and Gin web framework.

# Heroku deployment
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