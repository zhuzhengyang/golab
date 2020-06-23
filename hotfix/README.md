[![GoDoc](https://godoc.org/github.com/zhuzhengyang/golab/hotfix?status.svg)](https://godoc.org/github.com/zhuzhengyang/golab/hotfix) [![Go Report Card](https://goreportcard.com/badge/github.com/zhuzhengyang/golab/hotfix)](https://goreportcard.com/report/github.com/zhuzhengyang/golab)
# Usage
- Add `hotfix.Watch()` in your main process, and run it.
- Use `hotfix.Build()` and `hotfix.RegisterPatch()` to generate a plugin in anther script.
- The main process will autoload the plugin and call the monkeypatch function inside it.
- Look up the ./example folder to explore the detail. Try `test.sh`.

