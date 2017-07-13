# go-url-preview

This is a small Go server that will fetch the title of a webpage given to it.

I.E. a get request that looks like:

http://<wherever-this-is-running>/getTitle?url=https://google.com

will return

{"title" : "google"}

Enjoy!

###Installing

You'll need [Go 1.7+](https://golang.org/dl/), and [glide](https://github.com/Masterminds/glide). Install Go, install glide, and navigate to this directory.

When you've found your way here, run:
```glide up && go run main.go```

That should get the server running on `127.0.0.1:8000`

Have fun!
