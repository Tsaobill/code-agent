package function

type FunctionDef struct {
    Name        string                 `json:"name"`
    Description string                 `json:"description"`
    Parameters  map[string]interface{} `json:"parameters"`
}

type FunctionCall struct {
    Name      string          `json:"name"`
    Arguments json.RawMessage `json:"arguments"`
}


func registerFunctions() []FunctionDef {
    return []FunctionDef{
        {
            Name:        "generate_file",
            Description: "生成源代码文件",
            Parameters: map[string]interface{}{
                "type": "object",
                "properties": map[string]interface{}{
                    "filename": map[string]interface{}{
                        "type":        "string",
                        "description": "文件名",
                    },
                    "content": map[string]interface{}{
                        "type":        "string",
                        "description": "文件内容",
                    },
                },
                "required": []string{"filename", "content"},
            },
        },
        // add more func here ...
    }
}

func handleFunctionCall(call FunctionCall) (string, error) {
    switch call.Name {
    case "generate_file":
        var args struct {
            Filename string `json:"filename"`
            Content  string `json:"content"`
        }
        if err := json.Unmarshal(call.Arguments, &args); err != nil {
            return "", err
        }
        return generateFile(args.Filename, args.Content)
    // handle other function calls...
    default:
        return "", fmt.Errorf("unknown function: %s", call.Name)
    }
}
