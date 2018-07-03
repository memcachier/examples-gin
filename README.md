# MemCachier and Gin on Heroku tutorial

This is an example Gin Gonic app that uses the
[MemCachier add-on](https://addons.heroku.com/memcachier) on
[Heroku](http://www.heroku.com/). A running version of this app can be
found [here](http://memcachier-examples-gin.herokuapp.com).

Detailed instructions for developing this app are available
[here](https://devcenter.heroku.com/articles/gin-memcache).

## Deploy to Heroku

You can deploy this app yourself to Heroku to play with.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Running Locally

Run the following commands to get started running this app locally:

```sh
$ cd $GOPATH/src
$ git clone https://github.com/memcachier/examples-gin.git
$ cd examples-gin
$ govendor sync
$ memcached &  # run a local memcached server instance
$ MEMCACHIER_SERVERS=localhost:11211 go run main.go
```

Then visit `http://localhost:3000` to play with the app.

Note: instead of running a local `memcached` server you can also create a
[MemCachier](https://www.memcachier.com/) cache and add the `MEMCACHIER_*`
variables to the environment).

## Deploying to Heroku

Run the following commands to deploy the app to Heroku:

```sh
$ git clone https://github.com/memcachier/examples-gin.git
$ cd examples-gin
$ heroku create
Creating app... done, ⬢ rocky-chamber-17110
https://rocky-chamber-17110.herokuapp.com/ | https://git.heroku.com/rocky-chamber-17110.git
$ heroku addons:create memcachier:dev
$ git push heroku master
$ heroku open
```

## Configuring MemCachier

Install [`memcachier/mc`](https://github.com/memcachier/mc) and configure it in
Gin as follows to use it with MemCachier:

```go
username := os.Getenv("MEMCACHIER_USERNAME")
password := os.Getenv("MEMCACHIER_PASSWORD")
servers := os.Getenv("MEMCACHIER_SERVERS")

config := mc.DefaultConfig()
config.Hasher = mc.NewModuloHasher()         // default
config.Retries = 2                           // default
config.RetryDelay = 200 * time.Millisecond   // default
config.Failover = true                       // default
config.ConnectionTimeout = 2 * time.Second   // default
config.DownRetryDelay = 60 * time.Second     // default
config.PoolSize = 1                          // default
config.TcpKeepAlive = true                   // default
config.TcpKeepAlivePeriod = 60 * time.Second // default
config.TcpNoDelay = true                     // default

mcClient := mc.NewMCwithConfig(servers, username, password, config)
```

## Get involved!

We are happy to receive bug reports, fixes, documentation enhancements,
and other improvements.

Please report bugs via the
[github issue tracker](http://github.com/memcachier/examples-gin/issues).

Master [git repository](http://github.com/memcachier/examples-gin):

* `git clone git://github.com/memcachier/examples-gin.git`

## Licensing

This example is open-sourced software licensed under the
[MIT license](https://opensource.org/licenses/MIT).
