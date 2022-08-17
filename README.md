# age

This is a very simple app implementing a single-page webapp with server
rendering.

It works in a simple way:

1. Request, for example: `/user/qeaml`

2. Serve content with full layour (navbar, footer, styles, scripts etc.)

3. Hijack all otherwise everyday hyperlinks:
   Instead of the default browser behavior of loading the entire page, just the
   main content of the page is fetched and the main content of the current page
   is replaced with the newly fetched content. This way only the main content
   of the page is refreshed, but everything else remains unchanged (for example,
   the navbar is not changed, and the loaded scripts are not reloaded).

4. At this point the webapp turns into a single-page webapp with server-side
   rendering, using server-side routing.

## Technical

This little app is written in [Go](go) using [Fiber](fiber).

Building is as simple as:

```shell
$ cd target
$ go build -o app ..
# replace `app` with `app.exe` if you're on Windows!!
```

This will place the app in the target folder, where the required static content
is located.

[go]: https://go.dev/
[fiber]: https://gofiber.io/
