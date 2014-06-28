package main

import (
	"os/user"
	"strings"
)

func AddDefaultVars(vars map[string]string) map[string]string {
	u, err := user.Current()
	if err == nil {
		if _, ok := vars["USER"]; !ok {
			vars["USER"] = u.Username
		}
		if _, ok := vars["UID"]; !ok {
			vars["UID"] = u.Uid
		}
		if _, ok := vars["GID"]; !ok {
			vars["GID"] = u.Gid
		}
		if _, ok := vars["HOME"]; !ok {
			vars["HOME"] = u.HomeDir
		}
	}
	return vars
}

func ReplaceVars(text string, vars map[string]string) string {
	vars = AddDefaultVars(vars)

	// FIXME use LTM so $LONGVAR is matched before $LONG
	// Should probably parse not global replace

	for k, v := range vars {
		text = strings.Replace(text, "$"+k, v, -1)
	}

	return text
}
