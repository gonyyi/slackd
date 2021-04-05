#!/usr/bin/python

import json
import signal
import sys


def new(timeout=10):
    return slackd(timeout)

class slackd:
    received = {}
    final = {}
    timeout = 0

    def __init__(self, timeout):
        # timeout for reading files.
        signal.signal(signal.SIGALRM, self.__timeout)
        signal.alarm(timeout)
        self.timeout = timeout
        # Set output
        self.final["custom"] = {}
        self.final["custom"]["outgoing"] = {}
        self.final["custom"]["outgoing"]["replyInThread"] = False
        self.final["custom"]["outgoing"]["messages"] = []
        self.final["custom"]["outgoing"]["files"] = []

        self.__getStdin()

    def __timeout(self, signum, frame):
        if self.received == {}:
            self.addError("timed out (%d sec)" % self.timeout)
            self.dump()
            exit(1)

    def __getStdin(self):
        raw = sys.stdin.readlines()
        try:
            self.received = json.loads("".join(raw))
        except:
            self.addError("The module can't be started. It received invalid JSON data. Contact the admin.")
            self.dump()
            exit(1)

    def getUser(self):
        return self.received.get("custom", {}).get("incoming", {}).get("user", "")

    def getText(self):
        return self.received.get("custom", {}).get("incoming", {}).get("text", "")

    def getType(self):
        return self.received.get("custom", {}).get("incoming", {}).get("type", "")

    def __add(self, type, markdown):
        self.final.get("custom", {}).get("outgoing", {}).get("messages", []).append({
            "type": type,
            "text": markdown
        })

    def dump(self):
        # sys.stdout.write(json.dumps(self.final, sort_keys=True, indent=3))
        sys.stdout.write(json.dumps(self.final, separators=(',', ':')))

    def addInfo(self, markdown):
        self.__add("info", markdown)

    def addWarn(self, markdown):
        self.__add("warning", markdown)

    def addError(self, markdown):
        self.__add("error", markdown)

    def addMarkdown(self, markdown):
        self.__add("markdown", markdown)

    def addFile(self, filename, title, comment, useGzip):
        self.final.get("custom", {}).get("outgoing", {}).get("files", []).append({
            "filename": filename,
            "title": title,
            "comment": comment,
            "useGzip": useGzip,
        })
