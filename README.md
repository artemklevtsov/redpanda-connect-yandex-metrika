# Redpanda Connect

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/artemklevtsov/redpanda-connect/tests.yml?branch=main)](https://github.com/artemklevtsov/redpanda-connect/actions?query=branch%3Amain)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/artemklevtsov/redpanda-connect)
[![GitHub Release](https://img.shields.io/github/v/release/artemklevtsov/redpanda-connect)](https://github.com/artemklevtsov/redpanda-connect/releases/latest)
[![GitHub License](https://img.shields.io/github/license/artemklevtsov/redpanda-connect)](https://github.com/artemklevtsov/redpanda-connect?tab=Apache-2.0-1-ov-file#readme)

Redpanda Connect with some extensions.

## Install

Build and Install binary:

```sh
go install github.com/artemklevtsov/redpanda-connect@latest
```

Or download binary with:

```sh
curl -s https://i.jpillora.com/artemklevtsov/redpanda-connect@latest! | bash
```

Remove `!` if you want to install to the current directory.

Alternatively pull a Docker image with:

```sh
docker pull ghcr.io/artemklevtsov/redpanda-connect@latest
```

## Run

```yaml
input:
  yandex_metrika_stat_table:
    ids:
      - 44147844
    metrics:
      - ym:s:users
      - ym:s:visits
    dimensions:
      - ym:s:date
      - ym:s:lastTrafficSource
    sort:
      - ym:s:date
      - ym:s:lastTrafficSource
    date1: 2025-02-01
    date2: 2025-02-28
    # filters: ym:s:lastTrafficSource=='direct'
    # filters: ym:s:date=='2025-02-03'

output:
  stdout: {}
```

And you can run it like this:

```sh
redpanda-connect run ./connect.yaml
```

See also `configs/` for the more examples.

## See also

### Documentation

#### Redpadna Connect

- [plugin docs](docs/modules/components/pages/)
- [redpanda connect docs](https://docs.redpanda.com/redpanda-connect/)

#### Yandex.Metrika API

##### Authorization

1. [Create app](https://oauth.yandex.ru/client/new/id)
2. Get token: <https://oauth.yandex.ru/authorize?response_type=token&client_id={APP_ID}>

- [debug token](https://yandex.ru/dev/id/doc/en/tokens/debug-token)

##### Reporting API

- [dimensions and metrics](https://yandex.ru/dev/metrika/en/stat/attrandmetr/dim_all)

##### Logs API

- [hits](https://yandex.ru/dev/metrika/en/logs/fields/hits)
- [visits](https://yandex.ru/dev/metrika/en/logs/fields/visits)
- [parametrization](https://yandex.ru/dev/metrika/en/logs/param)
