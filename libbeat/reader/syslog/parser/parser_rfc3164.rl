// Code generated by ragel. DO NOT EDIT.
package syslog

import (
    "time"
)

%%{
    machine rfc3164_parser;

    write data;
}%%

// parseRFC3164 parses an RFC 3164-formatted syslog message. loc is used to enrich
// timestamps that lack a time zone.
func parseRFC3164(data string, loc *time.Location) (message, error) {
    var m message
    var err error
    var p, cs, tok int

    pe := len(data)
    eof := len(data)

    %%{
        include common "common.rl";
        include rfc3164 "rfc3164.rl";

        main := pri timestamp sp hostname sp msg;

        write init;
        write exec;
    }%%

    if err != nil {
        return message{}, err
    }

    return m, nil
}