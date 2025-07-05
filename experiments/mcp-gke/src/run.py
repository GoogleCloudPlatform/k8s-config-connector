import argparse
import asyncio
import os
from server import MCPForGKEServer


def main():
    KUBE_CONFIG_DEFAULT_LOCATION = os.environ.get('KUBECONFIG', '~/.kube/config')

    server = MCPForGKEServer(kubeconfig=KUBE_CONFIG_DEFAULT_LOCATION)
    server.run(transport='stdio')


if __name__ == "__main__":
    main()