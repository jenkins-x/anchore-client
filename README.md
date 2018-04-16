# Anchore Client

This is a golang client for [anchore-engine](https://github.com/anchore/anchore-engine).  This project was generated from the anchore engine swagger.

First install go-swagger see https://github.com/go-swagger/go-swagger#installing

Now generate the client:

```
swagger generate client -f  https://raw.githubusercontent.com/anchore/anchore-engine/master/anchore_engine/services/apiext/swagger/swagger.yaml -A anchore
```