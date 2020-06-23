# Usage
- Add `hotfix.Watch()` in your main process, and run it.
- Use `hotfix.Build()` to generate a plugin in anther script.
- The main process will autoload the plugin and call the monkeypatch function inside it.
- Look up the ./example folder to explore the detail. Try `test.sh`.

