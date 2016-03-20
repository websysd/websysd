websysd
=======

Like [Marathon](https://github.com/mesosphere/marathon) or [Upstart](https://code.launchpad.net/upstart), for your desktop!

![Screenshot of websysd workspace list](/images/websysd_workspaces.png "websysd workspaces")

### Getting started

Download a [binary release](https://github.com/ian-kent/websysd/releases), or run it with Docker:

```
docker run -v `pwd`/workspace.json:/workspace.json -v `pwd`/websysd.json:/websysd.json iankent/websysd -workspace=/workspace.json
```

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

- Active tasks will be killed if `websysd` dies or is stopped
- Use the `/bin/sh -c` executor on Linux
- Use custom columns and functions to add UI metadata, e.g. display git branch name
  - see examples in [websysd.json](websysd.json) and [workspace.json](workspace.json)

#### Environment

Default websysd behaviour is to ignore all preset environment variables.

This means you will need to set any variables you want explicitly (including `$PATH`,
which you might want to set to `/bin:/usr/bin:/usr/local/bin`).

You can change this behaviour by setting `InheritEnvironment` in either the global
configuration or a workspace configuration file:

    {
    	"InheritEnvironment": true
    }

If `InheritEnvironment` is true in the global workspace, the setting is ignored by
individual workspaces and the full environment will be inherited anyway.

#### STDOUT/STDERR logs

Output from tasks is stored in-memory by default.

You can set `Stdout` and `Stderr` to a filename per-task to override this behaviour.

Filenames can include environment variables, and `$TASK` and `$RUN` are set by websysd.

    {
    	"Stdout": "/tmp/$TASK-$RUN.out"
    }

### Screenshots

![Screenshot of websysd task list](/images/websysd_tasks.png "websysd task list")
![Screenshot of websysd task view](/images/websysd_task.png "websysd task view")
![Screenshot of websysd log view](/images/websysd_stdout.png "websysd log view")

### Licence

Copyright ©‎ 2014 - 2016, Ian Kent (http://iankent.uk).

Released under MIT license, see [LICENSE](LICENSE.md) for details.
