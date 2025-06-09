### dont_at_me

The TUI for checking if a username is available on various social media platforms. Built with only the Go standard library and ANSI escape codes.

![dont_at_me](./assets/Screenshot%202023-03-29%20at%206.00.07%20PM.png)

### Running the project

```shell
cd dont_at_me
go run cmd/main.go
```

### This Fork

I forked the original project to tinker with it, but realized it's too much of a pain to get reliable answers.

1. At least some of the existing string matches don't work. I changed the
2. I did update the `socials` package to be much simpler, using a struct to store all the info about the sites instead
of separate variables.
3. I made the string to search for part of the struct so it would be easy to add a new site - maybe even via a config
file. I added a boolean that would determine whether the string match indicates that the username exists or is taken.

In the end, I think most sites are too complex to do simple scraping for a string to determine whether an account
exists. Or maybe I just coincidentally started with the more opaque ones. In any case, the [Sherlock](https://github.com/sherlock-project/sherlock)'s
long [list of sites that used to work but no longer do](https://github.com/sherlock-project/sherlock/blob/4423230c117a5c931a1c854d722609160bf5fcb3/docs/removed-sites.md)
seems to show that it's generally quite an annoying task.

