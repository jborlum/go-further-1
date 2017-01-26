Earlier today I stumbled upon this little gem showcasing a really cool way to use channels to create a pause / continuation construct for async queues (or whatever you need it for).

The pattern uses the fact that channels are first-class values which allows you to pass channels through channels. This enables you to create a very simple, thread-safe way to communicate with a running goroutine.

I've create a distilled example to better understand it and wanted to share it with whomever would be interested :)