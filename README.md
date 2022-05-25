# Envs API

## Configuring Development environment with Docker:

1. Clone repo
```
git clone git@github.com:amomama/envs_demo.git
cd envs_demo
```

2. Add or change .env file

3. Build containers
```./make debug``` or ```./make rebuild```

### When server starts you`ll see
```
┌───────────────────────────────────────────────────┐
│                   Fiber v2.25.0                   │
│               http://127.0.0.1:8080               │
│       (bound on host 0.0.0.0 and port 8080)       │
│                                                   │
│ Handlers ............ 62  Processes ........... 1 │
│ Prefork ....... Disabled  PID ............. 26389 │
└───────────────────────────────────────────────────┘
```

## Swager Api docs

```http://127.0.0.1:8080/docs/```

Login:password for docs: ```admin:sysadmin```

## Run tests

```./make tests```

