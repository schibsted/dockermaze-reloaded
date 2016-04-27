import requests
import json
import base64
from maze_lib.maze_solver import MazeSolver
import time
import os


# LEGS VERIFICATION MODULE
#
# This module verifies that robot legs are able
# to find exit path from any maze in coordination
# with the head and other tools present in the robot.
#
# Mazes are generated randomly, obtained from robot
# head module. Some are hard, some are easy, but legs
# module must resolve all mazes with optimal path to
# obtain maximum score.
#
# Success conditions:
#  - At least, 70% of the mazes must be resolved.
# Score conditions:
#  - Correct but not optimal solutions reduces score.
#  - Shortest path solution to escape the mazes gives
#    max score.
#  - Robot legs module is unable to check solutions
#    with x4 number of commands than optimal exit path.
#  - Score rate is limited to 0.5 if not all solutions
#    can be checked.
#  - Score rate is limited to 0.7 if not all solutions
#    even if are good, are shortest path.
# Summary:
#  - To obtain maximum score, solve all mazes providing
#    optimal exit path (shortest exit path).

ENDPOINT = os.getenv('LEGS_ENDPOINT', 'head:80')


def main():
    start_time = int(time.time())
    # To verify robot legs function is required get
    # new maze challenges from the head module.
    try:
        print '[*] Obtaining challenge...'
        c = requests.get("http://{0}/challenge/new".format(ENDPOINT))
        if not c.ok:
            raise
    # If no challenge can be retrieved, ensure that
    # legs can contact robot head module.
    except:
        print '[!] Unable to retrieve the challenge'
        import sys
        sys.exit(1)

    payload = c.json()
    mazes = json.loads(base64.b64decode(payload['challenge']))
    solutions = {}
    print '[*] Calculating solutions...'
    for maze_id, m in mazes.iteritems():
        # MazeSolver provides main functionality to legs module.
        # In order to assemble legs module and provide best
        # performance ensure that mazes can be resolved and exit
        # paths are found in less possible steps.
        ms = MazeSolver(m)
        solution = ms.solve_maze()
        # solutions is a list of commands (UP,DOWN,LEFT,RIGHT)
        # which defines a path to exit the maze.
        solutions[maze_id] = solution
        # DEBUG?
        # print "MAZE:\n{0}\nSOLUTION:\n{1}\n".format(m,solution)

    result = {'challenge': payload['challenge'], 'start_time': start_time,
              'response': base64.b64encode(json.dumps(solutions)),
              'end_time': int(time.time())}

    headers = {'Content-Type': 'application/json'}
    print '[*] Sending solutions...'
    try:
        res = requests.post("http://{0}/challenge".format(ENDPOINT),
                            headers=headers, data=json.dumps(result))
        res_payload = res.json()
        print base64.b64decode(res_payload['message'])
        print "LEGS MODULE ASSEMBLED: {0}".format(res_payload['success'])
        print "SCORE RATIO: {0}\n".format(res_payload['score_rate'])
    except:
        print '[!] Unable to post solution'
        import sys
        sys.exit(1)
