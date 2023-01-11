name: generate golang docs
on: [push, pull_request]
  jobs:
    generate_docs:
      runs-on: ubuntu-latest
      steps:
        - name: Checkout repo
          uses: actions/checkout@v2
        - name: setup go
          uses: actions/setup-go@v2
        - name: generate docs
          uses: gofmt ./...
