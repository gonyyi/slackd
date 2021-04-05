#!/usr/bin/python

import slackd

if __name__ == "__main__":
    s = slackd.new(3) # takes timeout value
    s.addInfo("Process finished and a file will be downloaded")
    s.addFile("test.py", "abctitle", "my xx" + s.getUser(), True)
    s.dump()

