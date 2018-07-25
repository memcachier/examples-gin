# MemCachier and Gin on Elastic Beanstalk tutorial

This is an example Gin Gonic app that uses the
[MemCachier add-on](https://addons.heroku.com/memcachier) on
[Elastic Beanstalk](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/Welcome.html).

Detailed instructions for developing this app are available
[here](https://blog.memcachier.com/2018/07/30/gin-elastic-beanstalk-and-memcache).

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

Then visit `http://localhost:5000` to play with the app.

Note: instead of running a local `memcached` server you can also create a
[MemCachier](https://www.memcachier.com/) cache and add the `MEMCACHIER_*`
variables to the environment).

## Deploy to Elastic Beanstalk

You can deploy this app yourself to Elastic Beanstalk to play with.

```bash
$ eb init

# We'll stick with the default for now.
Select a default region
[...]
(default is 3): 3

Select an application to use
?) [ Create new Application ]
(default is 1): # Select whichever option lets you create a new application

# You can make the name whatever you like.
# By default it will match the file directory.
Enter Application Name
(default is "gin-memcache"): gin-memcache
Application gin-memcache has been created.

# This is a go tutorial, so we'll pick go.
Select a platform.
[...]
10) Go
[...]
(default is 1): 10


Select a platform version.
1) Go 1.10
2) Go 1.9
3) Go 1.8
4) Go 1.6
5) Go 1.5
6) Go 1.4
(default is 1): # Select your version of go.

# don't worry about this bit for now. You can always turn it on later using `eb init`
Note: Elastic Beanstalk now supports AWS CodeCommit; a fully-managed source control service. To learn more, see Docs: https://aws.amazon.com/codecommit/
Do you wish to continue with CodeCommit? (y/N) (default is n): n

# You can if you want, but we'll skip that that now.
Do you want to set up SSH for your instances?
(Y/n): n
```

Now that you've set up your application repository, we'll need to create the EB
instance.

```bash
$ eb create

# Can be whatever you want. We'll stick with the default for now.
Enter Environment Name
(default is gin-memcache-dev): gin-memcache-dev

# This will be the beginning of all of your URLs. We won't be doing anything
# with this, so we'll use the default option again.
Enter DNS CNAME prefix
(default is gin-memcache-dev): gin-memcache-dev

# In case you're sensing a theme here.. We'll stick with default for now.
Select a load balancer type
1) classic
2) application
3) network
(default is 1): 1
```

```bash
$ eb setenv MEMCACHIER_USERNAME=<username> MEMCACHIER_PASSWORD=<password> MEMCACHIER_SERVERS=<servers>
# ...
$ eb deploy
```
