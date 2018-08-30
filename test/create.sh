#!/bin/bash
request_body=$(cat <<EOF
{
        "resource":"task"
    }
    EOF
    )

    curl -v -X POST \
             -d "$request_body" \
             'http://127.0.0.1:7878/reboot/api/v1/namespaces/chenchao/tasks' | python -m json.tool
