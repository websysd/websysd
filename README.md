websysd
=======

Like systemd or Upstart - not as good, but in a browser!

And written in Go :)

### Getting started

- Optionally set a global environment - see [websysd.json](websysd.json) for an example
- Create a JSON workspace file - see [workspace.json](workspace.json) for an example
- Start `websysd`
- Open [http://localhost:7050](http://localhost:7050)

If you didn't name your config files `websysd.json` and `workspace.json`:

    websysd -global websysd.local.json -workspace myworkspace.json

And if you want to load multiple workspaces:

    websysd -workspace myworkspace.json -workspace someproject.json

### Why

Too many console windows.

### Useful info

- None of your environment is kept - explicitly set anything you need
  - This includes `$PATH` - you might want to set it to `/bin:/usr/local/bin`
- Logs are stored in memory by default (set task Stdout/Stderr to a filename)
- Active tasks will be killed if `websysd` dies or is stopped
- Use the `/bin/sh -c` executor on Linux

### Screenshots

![Screenshot of websysd task list](/images/websysd_tasks.png "websysd task list")
![Screenshot of websysd task view](/images/websysd_task.png "websysd task view")
![Screenshot of websysd log view](/images/websysd_stdout.png "websysd log view")

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.
