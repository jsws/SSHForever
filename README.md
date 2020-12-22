# SSHForever
[![Go Report Card](https://goreportcard.com/badge/github.com/jsws/SSHForever)](https://goreportcard.com/report/github.com/jsws/SSHForever)

An SSH tarpit written in Go with Prometheus metrics, inspired by [endlessh](https://github.com/skeeto/endlessh).

The [RFC](https://tools.ietf.org/html/rfc4253#page-4) for SSH allows the server to send "other lines of data" before sending the version string. This means the server can keep sending data to connected clients until they give up, slowing down automated scripts. 

Of course in the grand scheme of internet SSH scanners and dictionary attackers it doesn't make a massive difference, but it's just for fun! Over a 5 day period 8414 clients connected, cumulatively wasting over 25 weeks of time, the mean time wasted was 31 minutes, the longest a client was stuck was 81 hours.


## Usage
```
$ ./sshforever
Usage of ./sshforever:
  -l string
        Socket to listen for SSH connections on. (default "0.0.0.0:22")
  -m string
        Socket to run the Prometheus metrics server on. (default "0.0.0.0:2112")
```

## Metrics
Prometheus metrics are exposed via HTTP at `/metrics` on the socket given by the `-m` argument (default `0.0.0.0:2112`).  
The default Go metrics are exposed in addition to the following:

| Metric                   | Description                                                                     | Type      |
|--------------------------|---------------------------------------------------------------------------------|-----------|
| sshforever_time_wasted_milliseconds | The time clients stayed connected in milliseconds. 10 buckets, 60 seconds wide. | Histogram |
| sshforever_connections_current      | The current number of SSH connections being held up.                            | Gauge     |