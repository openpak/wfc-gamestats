# Changelog

## v0.1.2 - 2026-11-03

- Internal `/internal/` polls are counted and summarised once a minute instead of logged per call (wfc-common v0.3.1). Parse-error and peer-message log lines now record byte counts instead of the raw message; gamestats no longer logs set-data values.

## v0.1.1 - 2026-11-03

- Every gamestats message now logs one line to stderr, `gamestats cmd <name> -> <result> <ms>ms` (verb and result only, never payload values) (results ok, error, denied, unknown). Internal HTTP logs as above through wfc-common v0.3.0 `httplog`.

