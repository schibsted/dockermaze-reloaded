import requests
import json
import base64
import time
import os

ENDPOINT = os.getenv('HEART_ENDPOINT', 'head:80')


def main():
    start_time = int(time.time())
    try:
        print '[*] Sending heartbeat...'
        c = requests.get("http://{0}/heart/beat".format(ENDPOINT))
        if not c.ok:
            raise
    except:
        print '[!] Unable to retrieve heartbeat response'
        import sys
        sys.exit(1)

    payload = c.json()
    random_check = base64.b64decode(payload['challenge'])

    print '[*] Heartbeat response received...'

    result = {'challenge': payload['challenge'], 'start_time': start_time,
              'response': base64.b64encode(random_check),
              'end_time': int(time.time())}

    headers = {'Content-Type': 'application/json'}
    print '[*] Connecting heart...'

    try:
        res = requests.post("http://{0}/heart".format(ENDPOINT),
                            headers=headers, data=json.dumps(result))

        res_payload = res.json()

        print base64.b64decode(res_payload['message'])
        print "HEART MODULE ASSEMBLED: {0}".format(res_payload['success'])
        print "SCORE RATIO: {0}\n".format(res_payload['score_rate'])
    except Exception as e:
        print "<p>Error: %s</p>" % str(e)
        print '[!] Unable to process heartbeat'
        import sys
        sys.exit(1)
