import subprocess
import json
import argparse
import os

def grpc_call(port, rpc_name, input=None, auth=''):
    
    echo = subprocess.Popen(('echo', f'{input}'), stdout=subprocess.PIPE)
    try:
        output = subprocess.check_output(
            [
                'evans', 'cli', '-r',
                '-p', str(port),
                'call', rpc_name,
                '--header', f'authorization={auth}',
            ],
            # input gets echoed + piped into the evans command
            stdin=echo.stdout
        )
    except subprocess.CalledProcessError as e:
        print(e)
        print('\nCould not login with given credentials.')
        exit(1)
    return json.loads(output.decode('utf-8'))

if os.name == 'nt':
    # Windows system, cannot Popen
    print('Cannot run on Windows.')
    exit(1)

parser = argparse.ArgumentParser()
parser.add_argument('-u', '--user', type=str, required=True)
parser.add_argument('-p', '--password', type=str, required=True)

args = parser.parse_args()

login_json = f'{{"username": "{args.user}", "password": "{args.password}"}}'
login = grpc_call(50054, 'LoginUser', input=login_json)
access_token = login['result']['AccessToken']


print(f'Access Token for {args.user}:')
print(f'"{access_token}"')
