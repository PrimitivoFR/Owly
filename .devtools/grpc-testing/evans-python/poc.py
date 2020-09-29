import subprocess
import json
from pprint import pprint


def grpc_call(port, rpc_name, input_file, auth=''):
    output = subprocess.check_output(
        [
            'evans', 'cli', '-r', '-p',
            str(port),
            'call', rpc_name,
            '--file', input_file,
            '--header', f'authorization={auth}'
        ]
    )
    return json.loads(output.decode('utf-8'))


login = grpc_call(50051, 'LoginUser', 'json/login.json')
access_token = login['result']['AccessToken']

user = grpc_call(
    50051, 'SearchUserByUsername', 'json/search_user.json'
)['results'][0]

chatrooms = grpc_call(
    50052, 'GetChatroomsByUser', 'json/empty.json', auth=access_token
)

print('login response:')
pprint(login)
print(f'user:\n{user}\n')
print('chatrooms:')
pprint(chatrooms)
