# code-agent
agentic coding tool

## How to Use

1. git clone this repo
2. build the binary
```bash
   go build -o code-agent
   sudo mv code-agent /usr/local/bin/
```
3. Set ENV
```bash
export ANTHROPIC_API_KEY="your_api_key_here"
```
4. Run
```bash
code-agent "please write a hello world in python"
```