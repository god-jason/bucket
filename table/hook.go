package table

import "github.com/dop251/goja"

type Hook struct {
	string
	program *goja.Program
}

func (h *Hook) Compile() (err error) {
	if h.string != "" {
		h.program, err = CompileJavaScript(h.string)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Hook) Run(context map[string]any) error {
	if h.program != nil {
		runtime := CreateJavaScriptRuntime()
		err := runtime.Set("context", context)
		if err != nil {
			return err
		}

		_, err = runtime.RunProgram(h.program)
		//打印返回值？？？
		if err != nil {
			return err
		}
	}
	return nil
}
