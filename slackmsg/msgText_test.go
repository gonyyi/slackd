package slackmsg

import (
	"fmt"
	"testing"
)

const (
	// Original
	testMsg1 = `{
    "token": "XLVmORnn6ha4XWXDqbcxfW0h",
    "team_id": "T3W9PKG9G",
    "api_app_id": "A01HB0YV6QM",
    "event": {
        "client_msg_id": "f16483e7-d87a-44f6-946c-cefcd53f4c7e",
        "type": "message", 
        "text": "hello1",
        "user": "U01CTJH5ZJ9",
        "ts": "1617207014.001100",
        "team": "T3W9PKG9G",
        "blocks": [
            {
                "type": "rich_text",
                "block_id": "nx0",
                "elements": [
                    {
                        "type": "rich_text_section",
                        "elements": [
                            {
                                "type": "text",
                                "text": "hello2"
                            }
                        ]
                    }
                ]
            }
        ],
        "channel": "D01HSN5P7JM",
        "event_ts": "1617207014.001100",
        "channel_type": "im"
    },
    "type": "event_callback",
    "event_id": "Ev01SS3CSUCW",
    "event_time": 1617207014,
    "authorizations": [
        {
            "enterprise_id": null,
            "team_id": "T3W9PKG9G",
            "user_id": "U01HE9PDSP5",
            "is_bot": true,
            "is_enterprise_install": false
        }
    ],
    "is_ext_shared_channel": false,
    "event_context": "1-message-T3W9PKG9G-D01HSN5P7JM"
}`
	// Update
	testMsg2 = `{
    "token": "XLVmORnn6ha4XWXDqbcxfW0h",
    "team_id": "T3W9PKG9G",
    "api_app_id": "A01HB0YV6QM",
    "event": {
        "type": "message",
        "subtype": "message_changed", 
        "hidden": true,
        "message": {
            "client_msg_id": "f16483e7-d87a-44f6-946c-cefcd53f4c7e",
            "type": "message",
            "text": "hello there3",
            "user": "U01CTJH5ZJ9",
            "team": "T3W9PKG9G",
            "edited": {
                "user": "U01CTJH5ZJ9",
                "ts": "1617207095.000000"
            },
            "blocks": [
                {
                    "type": "rich_text",
                    "block_id": "qDd=G",
                    "elements": [
                        {
                            "type": "rich_text_section",
                            "elements": [
                                {
                                    "type": "text",
                                    "text": "hello there4"
                                }
                            ]
                        }
                    ]
                }
            ],
            "ts": "1617207014.001100",
            "source_team": "T3W9PKG9G",
            "user_team": "T3W9PKG9G"
        },
        "channel": "D01HSN5P7JM",
        "previous_message": {
            "client_msg_id": "f16483e7-d87a-44f6-946c-cefcd53f4c7e",
            "type": "message",
            "text": "hello",
            "user": "U01CTJH5ZJ9",
            "ts": "1617207014.001100",
            "team": "T3W9PKG9G",
            "blocks": [
                {
                    "type": "rich_text",
                    "block_id": "nx0",
                    "elements": [
                        {
                            "type": "rich_text_section",
                            "elements": [
                                {
                                    "type": "text",
                                    "text": "hello"
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        "event_ts": "1617207095.001300",
        "ts": "1617207095.001300",
        "channel_type": "im"
    },
    "type": "event_callback",
    "event_id": "Ev01T51VLL8L",
    "event_time": 1617207095,
    "authorizations": [
        {
            "enterprise_id": null,
            "team_id": "T3W9PKG9G",
            "user_id": "U01HE9PDSP5",
            "is_bot": true,
            "is_enterprise_install": false
        }
    ],
    "is_ext_shared_channel": false,
    "event_context": "1-message-T3W9PKG9G-D01HSN5P7JM"
}`
	// Delete
	testMsg3 = `{
    "token": "XLVmORnn6ha4XWXDqbcxfW0h",
    "team_id": "T3W9PKG9G",
    "api_app_id": "A01HB0YV6QM",
    "event": {
        "type": "message",
        "subtype": "message_deleted",
        "hidden": true,
        "deleted_ts": "1617207014.001100",
        "channel": "D01HSN5P7JM",
        "previous_message": {
            "client_msg_id": "f16483e7-d87a-44f6-946c-cefcd53f4c7e",
            "type": "message",
            "text": "hello there5",
            "user": "U01CTJH5ZJ9",
            "ts": "1617207014.001100",
            "team": "T3W9PKG9G",
            "edited": {
                "user": "U01CTJH5ZJ9",
                "ts": "1617207095.000000"
            },
            "blocks": [
                {
                    "type": "rich_text",
                    "block_id": "qDd=G",
                    "elements": [
                        {
                            "type": "rich_text_section",
                            "elements": [
                                {
                                    "type": "text",
                                    "text": "hello there6"
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        "event_ts": "1617207252.001500",
        "ts": "1617207252.001500",
        "channel_type": "im"
    },
    "type": "event_callback",
    "event_id": "Ev01SYSUCRTL",
    "event_time": 1617207252,
    "authorizations": [
        {
            "enterprise_id": null,
            "team_id": "T3W9PKG9G",
            "user_id": "U01HE9PDSP5",
            "is_bot": true,
            "is_enterprise_install": false
        }
    ],
    "is_ext_shared_channel": false,
    "event_context": "1-message-T3W9PKG9G-D01HSN5P7JM"
}`
)

func TestPostFile(t *testing.T) {
	if err := postFile("xoxb-437964455474-1171621666130-db5LmmMXAl3lDARLL39rWiIt", "D015QGYPV0F", "1616776764.002600", "./", "field.go", "Test field", "awessome comment", true); err != nil {
		println(err.Error())
	}
}

func TestMsgTxt(t *testing.T) {
	mi := incoming{}
	mi.Load([]byte(testMsg2))
	mi.process()
	fmt.Printf("%s\n", mi.Custom.User)
}

func TestOutgoing(t *testing.T) {
	o := outgoing{}
	o.Load([]byte(`{"custom":{"files":[{"comment":"my xxgon","title":"abctitle","useGzip":true,"filename":"test.py"}],"messages":[{"text":"process finished and a file will be downloaded","type":"info"}],"replyInThread":false}}`))
	o.Token = "xoxb-437964455474-1171621666130-db5LmmMXAl3lDARLL39rWiIt"
	o.Channel = "D015QGYPV0F"
	o.ThreadTs = "1616776764.002600"
	o.dir = "/Users/gonyi/go/src/github.com/gonyyi/slackd/modules/testModBasic"

	if false {
		o.Text = "my random text"
		o.Custom.ReplyInThread = true
		o.Custom.Messages = []customOutgoingMsg{
			{Text: "error happens", Type: MSG_ERROR},
			{Text: "information about ...", Type: MSG_INFO},
			{Text: "information about some warning", Type: MSG_WARNING},
			{Text: "just typical...", Type: MSG_MARKDOWN},
		}
		o.Custom.Files = []customOutgoingFile{
			{
				Filename: "element.go",
				Title: "Element file",
				Gzip: true,
			},
		}
	}

	o.Process()
	if err := o.Post(); err != nil {
		println(err.Error())
	}
}