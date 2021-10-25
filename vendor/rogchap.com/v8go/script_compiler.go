package v8go

type ScriptCompilerCachedData []byte

type ScriptCompilerCompileOption int

const (
	ScriptCompilerCompileOptionNoCompileOptions = iota
	ScriptCompilerCompileOptionConsumeCodeCache
	ScriptCompilerCompileOptionEagerCompile
)
