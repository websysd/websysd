websysd
=======

Like systemd or Upstart - not as good, but in a browser!

And written in Go :)

### Getting started

- Create a JSON task file - see [websysd.json](websysd.json) for an example
- Start `websysd -config taskfile.json`
- Open [http://localhost:7050](http://localhost:7050)

### Why

Too many console windows.

### Useful info

- None of your environment is kept - explicitly set anything you need
- Logs are (currently) stored in-memory - you might get out of memory errors!
- Active tasks will be killed if `websysd` dies or is stopped
- Use the `/bin/sh -c` executor on Linux
- You can load multiple JSON task files, use multiple `-config` flags
- Environment variables are local to each task or JSON task file

### Screenshots

![Screenshot of websysd task list](/images/websysd_tasks.png "websysd task list")
![Screenshot of websysd task view](/images/websysd_task.png "websysd task view")
![Screenshot of websysd log view](/images/websysd_stdout.png "websysd log view")

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.
